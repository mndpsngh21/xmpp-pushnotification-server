package httpserve

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func FCMSendNotificationToDevice(token string) {
	// Initialize the FCM app
	ctx := context.Background()
	opt := option.WithAPIKey("AAAAuxp2HgI:APA91bFvKBzaG83AWvKN33VTVwPGtcGACLMa9afYf71UYuonq0x7XTnKbzT_jgJSxHQIZlh8VQIrDWXqbDABW7YpNm-kDeCsPTzOtEL4Zd3bwCdCDHgfroj7Q1KWYvuHWt8sz8SsmlfN-ltSaMGSA_PcjGwoLMcJcQ")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Get the FCM client
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Failed to get FCM client: %v", err)
	}

	// Send a test notification
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Test Notification",
			Body:  "This is a test notification from Go!",
		},
		Token: token,
	}

	// Send the message
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Print the response
	fmt.Println("Successfully sent message:", response)

}
