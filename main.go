package main

import (
	"log"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cloudeventprovider "github.com/eclipse-xfsc/cloud-event-provider"
	"github.com/eclipse-xfsc/event-log-service/config"
)

func init() {
	config.LoadConfig()
}

func receive(event cloudevents.Event) {
	eventData := string(event.Data())

	for _, keyword := range config.CurrentLogServiceConfig.Keywords {
		if strings.Contains(strings.ToLower(eventData), strings.ToLower(keyword)) {
			log.Printf("Event with keyword %s detected: %s", keyword, event)
		}
	}
}

func main() {
	client, err := cloudeventprovider.NewClient(cloudeventprovider.Sub, "events")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	log.Fatal(client.Sub(receive))
}
