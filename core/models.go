package core

// Work is the unit of request
type Work struct {
	Time int `json:"time"`
}

// Message return a string messge
type Message struct {
	Message string `json:"message"`
}
