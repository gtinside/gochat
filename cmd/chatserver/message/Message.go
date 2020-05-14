package message

import (
	"encoding/json"
	"fmt"
	"github.com/gtinside/gochat/internal/pkg/modal"
	"log"
	"net/http"
)

//short variable declarations are only allowed in functions
var awsSession = modal.GetDynamoDBSession()
var sqsSession = modal.GetSQSSession()

// SaveMessage to save the chat messages
func SaveMessage(w http.ResponseWriter, req *http.Request)  {
	var message modal.Message
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		handleError(err, w, " Failed to save messages ")
		return
	}

	if message.ConversationId = message.To +"-"+ message.From; message.From > message.To {
		message.ConversationId = message.From +"-"+ message.To
	}
	err = awsSession.Table("MessageDetails").Put(message).Run()
	if err != nil {
		handleError(err, w, " Failed to save messages ")
		return
	}
	log.Printf("Message from %v, to %v saved successfully with conv id %v",message.To,
		message.From, message.ConversationId)
	response := &modal.Response{Status:modal.COMPLETED, Msg: ""}
	result, _ := json.Marshal(response)
	fmt.Fprintf(w, string(result))
	return
}

// GetMessages to return all the messages received and sent
func GetMessages(w http.ResponseWriter, req *http.Request) {
	var message modal.Message
	err := json.NewDecoder(req.Body).Decode(&message)
	if err != nil {
		handleError(err, w, " Failed to save messages ")
		return
	}
	if message.ConversationId = message.To +"-"+ message.From; message.From > message.To {
		message.ConversationId = message.From +"-"+ message.To
	}
	var messages []modal.Message
	err = awsSession.Table("MessageDetails").Get("ConversationId", message.ConversationId).
		Index("ConversationIdIndex").All(&messages)
	if err != nil {
		handleError(err, w, " Error retrieving messages ")
		return
	}
	//Resolve the messages with User Name for proper handling
	m := make(map[string]string)
	var from, to modal.User
	tbl := awsSession.Table("UserDetails")
	tbl.Get("UserId", message.To).Index("UserIdIndex").One(&to)
	tbl.Get("UserId", message.From).Index("UserIdIndex").One(&from)

	m[message.From] = from.Name
	m[message.To] = to.Name

	for i := range messages {
		messages[i].To = m[messages[i].To]
		messages[i].From = m[messages[i].From]
	}
	u, _ := json.Marshal(&messages)
	response := &modal.Response{Status:modal.COMPLETED, Msg: string(u)}
	res, _ := json.Marshal(response)
	fmt.Fprintf(w, string(res))
}

// HandleError : Generic method to handle errors
func handleError(err error, w http.ResponseWriter, msg string) {
	log.Print(err)
	response := &modal.Response{Status:modal.ERROR, Msg: msg}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
