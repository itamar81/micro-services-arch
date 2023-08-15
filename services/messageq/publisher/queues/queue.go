package queues

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
	"io"
	"context"
	"time"
)
type Queue struct{
	qlog *log.Logger
	ch *amqp.Channel
	rabbit_q *amqp.Queue
}

func NewQueue(qlog *log.Logger,ch *amqp.Channel,queueName string) *Queue {
	rabbit_q  := createQueue(queueName,ch,qlog)
	return &Queue{qlog,ch,rabbit_q}
}
func (q *Queue) Submit(message string){
	data := []byte(message)
	q.SubmitBytes(data)
}
func createQueue(queueName string,ch *amqp.Channel,qlog *log.Logger) *amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	failOnError(qlog,err,"cannot create queue")
	return &q
}
func (u *Queue) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet {
		u.qlog.Printf("handle get")
		// u.handleGet(w ,r)
		return
	}
	resBody, err := io.ReadAll(r.Body)
	// u.qlog.Printf("Path: %s",r.URL.Query().Encode())
	failOnError(u.qlog,err,"client: could not read response body:")
	if len(resBody) == 0 {
		resBody = []byte(r.URL.Query().Encode())
	}
	u.qlog.Printf("send %s",resBody)
	u.SubmitBytes(resBody)
}

func (q *Queue) SubmitBytes(resBody []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := q.ch.PublishWithContext(ctx,
		"",
		q.rabbit_q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(resBody),
		})
	failOnError(q.qlog, err, "Failed to publish a message")
	q.qlog.Printf("Message sent: %s\n", resBody)
}

func failOnError(l *log.Logger, err error, msg string) {
	if err != nil {
		l.Panicf("%s: %s", msg, err)
	}
  }