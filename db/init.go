package db

import (
	"log"

	"github.com/tidwall/buntdb"
)

const USERSDB = "./database/users.db"
const SESSIONDB = "./database/sessions.db"

func Init() {
	db1, err := buntdb.Open(USERSDB)
	if err != nil {
		log.Printf("Error rendering %s: %s\n",USERSDB, err)
		return
	}
	defer db1.Close()
	db1.CreateIndex("contract", "*", buntdb.IndexJSON("user.id"))

	db2, err := buntdb.Open(SESSIONDB)
	if err != nil {
		log.Printf("Error rendering %s: %s\n",SESSIONDB, err)
		return
	}
	defer db2.Close()
}
