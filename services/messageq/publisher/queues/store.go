package queues

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
)
type Store struct{
	*Queue
}
func NewStore(l *log.Logger,ch *amqp.Channel) *Store{
	qq:=NewQueue(l,ch,"store")
	return &Store{qq}
}
func (store *Store) Submit( message string){
	store.Queue.qlog.Printf("from store %s",message)
}
func (store *Store) ServeHTTP(w http.ResponseWriter,r *http.Request){
	store.Queue.qlog.Printf("store queue")
	store.Queue.ServeHTTP(w,r)
}