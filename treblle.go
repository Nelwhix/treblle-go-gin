package treblle

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	timeoutDuration = 2 * time.Second
)

func sendToTreblle(treblleInfo MetaData) {
	log.Println("inside treblle goroutine")
	bytesRepresentation, err := json.Marshal(treblleInfo)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, Config.ServerURL, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Printf("error creating treblle request: %v", err.Error())
	}
	// Set the content type from the writer, it includes necessary boundary as well
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", Config.APIKey)

	// Do the request
	client := &http.Client{Timeout: timeoutDuration}
	_, err = client.Do(req)
	if err != nil {
		log.Printf("error making treblle request: %v", err.Error())
	}
}
