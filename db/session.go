package db

import (
	"encoding/json"
	"log"

	"github.com/tidwall/buntdb"
)

type Message struct {
	MemberId int
	Content string
}

type Session struct {
	Members    []User
	Messages []Message
	Identifier string
	OwnerId      int
}


func SessSave(key string, session Session) {
	db, err := buntdb.Open(SESSIONDB)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	dataJson, err := json.Marshal(session)
	if err != nil {
		log.Printf("Error converting data to json")
	}
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, string(dataJson), nil)
		return err
	})
	if err != nil {
		log.Printf("Save: %s\n", err.Error())
		return
	}
}

func SessRetrieve(key string) (*Session, bool) {
	db, err := buntdb.Open(SESSIONDB)
	if err != nil {
		return nil, false
	}
	defer db.Close()

	var session Session
	err = db.View(func(tx *buntdb.Tx) error {
		data, err := tx.Get(key)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(data), &session)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, false
	}
	return &session, true
}

func SessList() ([]Session, bool) {
	db, err := buntdb.Open(SESSIONDB)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	sessions := []Session{}
	err = db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			//fmt.Printf("key: %s, value: %s\n", key, value)
			session :=  Session{}
			json.Unmarshal([]byte(value), &session)
			sessions = append(sessions, session)
			return true // continue iteration
		})
		return err
	})
	if err != nil {
		return nil, false
	}
	return sessions, true
}

func SessDelete(keys []string) {
	db, err := buntdb.Open(SESSIONDB)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	
	err = db.View(func(tx *buntdb.Tx) error {
		for _, k := range keys {
			if _, err = tx.Delete(k); err != nil {
				return err
			}
		}
		return nil
	})
}
