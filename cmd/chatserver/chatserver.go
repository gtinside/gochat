package main

import (
	"github.com/gtinside/gochat/cmd/chatserver/message"
	"github.com/gtinside/gochat/cmd/chatserver/user"
	"net/http"
)


func main() {
	http.HandleFunc("/register", user.RegisterAction)
	http.HandleFunc("/saveMessage", message.SaveMessage)
	http.HandleFunc("/getMessages", message.GetMessages)
	http.HandleFunc("/getAllUsers", user.GetAllUsers)
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/getUser", user.GetUser)
	http.ListenAndServe(":8090", nil)
}