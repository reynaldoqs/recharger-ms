package main

import (
	"context"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	msgbroker "audiman/gu-project/recharger-ser/messageBroker"
	"audiman/gu-project/recharger-ser/services"
)

func main() {

	mb := setupRabbitMQ()
	rs := services.NewRechargerService(mb)
	go rs.InitRechargesListener()

	defer rs.CloseListener()

	// try
	sa := option.WithCredentialsFile("./config/audiman.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Println(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Println(err)
	}

	quote := getQoute()
	log.Println(quote)

	result, err := client.Collection("sampleData").Doc("inspiration").Set(context.Background(), quote)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
	defer client.Close()
	//end try

}

func getQoute() *Quote {
	myQuote := Quote{
		"Audiman 22 ",
		"Karem is tha best singer",
	}
	return &myQuote
}

type Quote struct {
	autor   string
	message string
}

func setupRabbitMQ() msgbroker.MessageBroker {
	//rabbitURL := os.Getenv("RABBIT_URL")
	rabbitURL := "amqp://guest:guest@localhost:5672/"
	brk, err := msgbroker.NewRabbitMqBroker(rabbitURL)
	if err != nil {
		time.Sleep(time.Second * 2)
		log.Printf("Trying to connect: %s\n", rabbitURL)
		return setupRabbitMQ()
	}
	return brk
}
