package modal

//Friend struct for user details
type Friend struct {
	Name   string
	Email  string
	UserId string
}

//User struct for user details
type User struct {
	Friend
	Password string
}

//Message struct for chat messages
type Message struct {
	To             string
	From           string
	Message        string
	Sent           string
	ConversationId string
}

//Response struct
type Response struct {
	Status StatusType
	Msg    string
}

type StatusType string

const (
	COMPLETED StatusType = "COMPLETED"
	ERROR     StatusType = "ERROR"
)
