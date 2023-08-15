package main

import (
	"log"
	"net/http"
	q "github.com/itamar81/micro-service-arch/messageq/publisher/queues"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"fmt"
	"encoding/json"
)
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
// "amqp://user:vkbHAhGx0KPiONXR@34.133.42.205:5672/"
var rabbit_host = getEnv("RABBIT_HOST","rabbitmq.rabbitmq.svc.cluster.local")
var rabbit_port = getEnv("RABBIT_PORT","5672") 
var rabbit_user = getEnv("RABBIT_USERNAME","user")
var rabbit_password = getEnv("RABBIT_PASSWORD","password")

func main(){
	l := log.New(os.Stdout, "api-gateway: ", log.LstdFlags)
	l.Printf("Application started")
	mux := http.NewServeMux()
	
	ch := openQueueConnection(l)
	userQueue := q.NewUsers(l,ch)
	storeQueue := q.NewStore(l,ch)
	mux.Handle("/users",userQueue)
	mux.Handle("/store",storeQueue)
	mux.HandleFunc("/",handleRoot)
	l.Printf("listen on port 8080")
	http.ListenAndServe(":8080",mux)
}
func handleRoot(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Resource Not Found"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
func openQueueConnection(l *log.Logger)  *amqp.Channel {
	rabbit_url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
	rabbit_user,
	rabbit_password,
	rabbit_host,
	rabbit_port)

	conn, err := amqp.Dial(rabbit_url)
	
	failOnError(l,err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(l,err, "Failed to open a channel")
	return ch
}
func failOnError(l *log.Logger,err error, msg string) {
	if err != nil {
	  l.Panicf("%s: %s", msg, err)
	}
  }

