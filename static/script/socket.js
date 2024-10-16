const socket = new WebSocket('localhost:9000/');

// Connection opened
socket.addEventListener('open', function (event) {
    console.log('Connected to the WebSocket server.');
});

// Listen for messages
socket.addEventListener('message', function (event) {
    const messageData = JSON.parse(event.data);
    
    // Create message element
    const chatBox = document.querySelector('.chat-box');
    const messageElement = document.createElement('div');
    messageElement.className = 'message other'; // Default to 'other' class

    messageElement.innerHTML = `
        <span class="sender">${messageData.MemberId}:</span>
        <span>${messageData.Content}</span>
        <span class="timestamp">(${messageData.Timestamp})</span>
    `;
    
    chatBox.appendChild(messageElement);
    chatBox.scrollTop = chatBox.scrollHeight; // Scroll to the bottom
});

// Send a message
document.querySelector('form[action="/send"]').addEventListener('submit', function (event) {
    event.preventDefault(); // Prevent form submission

    const messageInput = document.getElementById('message');
    const messageContent = messageInput.value.trim();

    if (messageContent) {
        // Create a message object with the current user's ID
        const message = {
            MemberId: currentUserId, // Use the current user's ID
            Content: messageContent,
            Timestamp: new Date().toLocaleTimeString()
        };

        // Send the message as a JSON string
        socket.send(JSON.stringify(message));
        messageInput.value = ''; // Clear the input
    }
});

// Handle errors
socket.addEventListener('error', function (event) {
    console.error('WebSocket error observed:', event);
});

// Handle connection close
socket.addEventListener('close', function (event) {
    console.log('WebSocket connection closed:', event);
});
