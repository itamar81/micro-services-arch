package queues

import (
	"log"
	"net/http"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Users struct{
	*Queue
}
func NewUserQueue(l *log.Logger,ch *amqp.Channel) *Users{
	qq := NewQueue(l,ch,"users")
	return &Users{qq}
}
func (u *Users) Submit( message string){
	u.Queue.Submit(message)
}
func (u Users) ServeHTTP(w http.ResponseWriter,r *http.Request){
	u.qlog.Printf("users handler called")
	u.Queue.ServeHTTP(w,r)
}

// func (u *Users) handleGet(w http.ResponseWriter,r *http.Request){
// 	msgs, err := u.ch.Consume(
// 		u.rabbit_q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	  )
// 	  failOnError(u.l, err, "Failed to register a consumer")
// 		for d := range msgs {
// 		  u.l.Printf("Received a message: %s", d.Body)
// 		  fmt.Fprintf(w, "Received a message: %s", d.Body)
// 		}
// 	  u.ch.Close()
// 	}

