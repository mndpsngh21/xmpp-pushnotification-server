package httpserve

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type DeviceReqisterRequest struct {
	DeviceToken string `json:"device_token"`
	JID         string `json:"jid"`
}

type PushNotificationRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

var tokenMap map[string]string = make(map[string]string)

func registerForNotification(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for registerForNotification")
	// Parse the request body
	var request DeviceReqisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process the request
	// Here, you can send the push notification using the device token and message

	// Print the received data
	fmt.Printf("Device Token: %s\n", request.DeviceToken)
	fmt.Printf("Message: %s\n", request.JID)
	tokenMap[request.JID] = request.DeviceToken
	// Send a response
	response := map[string]string{
		"status": "success",
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

}

func sendNotificationToDevice(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for Message")
	var request PushNotificationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process the request
	// Here, you can send the push notification using the device token and message

	// Print the received data
	fmt.Printf("From: %s\n", request.From)
	fmt.Printf("To: %s\n", request.To)
	fmt.Printf("Body: %s\n", request.Body)

	deviceToken := tokenMap[request.To]
	fCMSendNotificationToDevice(deviceToken, request.From, request.Body)
	//tokenMap[request.JID] = request.DeviceToken
	// Send a response
	response := map[string]string{
		"status": "success",
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func fCMSendNotificationToDevice(token string, from string, messageStr string) {
	// Initialize the FCM app
	ctx := context.Background()
	// opt := option.WithAPIKey("AAAAuxp2HgI:APA91bFvKBzaG83AWvKN33VTVwPGtcGACLMa9afYf71UYuonq0x7XTnKbzT_jgJSxHQIZlh8VQIrDWXqbDABW7YpNm-kDeCsPTzOtEL4Zd3bwCdCDHgfroj7Q1KWYvuHWt8sz8SsmlfN-ltSaMGSA_PcjGwoLMcJcQ")
	// config := &firebase.Config{ProjectID: "parentbustracking"}
	// app, err := firebase.NewApp(ctx, config, opt)
	// if err != nil {
	// 	log.Fatalf("Failed to initialize Firebase app: %v", err)
	// }
	config := &firebase.Config{ProjectID: "parentbustracking"}
	opt := option.WithCredentialsFile("trakom.json")
	//opt := option.WithAPIKey("AAAAuxp2HgI:APA91bFvKBzaG83AWvKN33VTVwPGtcGACLMa9afYf71UYuonq0x7XTnKbzT_jgJSxHQIZlh8VQIrDWXqbDABW7YpNm-kDeCsPTzOtEL4Zd3bwCdCDHgfroj7Q1KWYvuHWt8sz8SsmlfN-ltSaMGSA_PcjGwoLMcJcQ")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		//log.()
		fmt.Printf("error initializing app: %v", err)
	}
	// Get the FCM client
	client, err := app.Messaging(ctx)
	if err != nil {
		fmt.Printf("Failed to get FCM client: %v", err)
	}

	var payload map[string]string = make(map[string]string)
	payload["sender"] = from
	payload["sender"] = from
	message := &messaging.Message{
		Data:  payload,
		Token: token,
	}

	// Send the message
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}

	// Print the response
	fmt.Println("Successfully sent message:", response)

}

func StartHttpServer() {

	http.HandleFunc("/registerForNotification", registerForNotification)

	http.HandleFunc("/sendNotification", sendNotificationToDevice)

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
