package treblle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	timeoutDuration = 2 * time.Second
)

func sendToTreblle(treblleInfo MetaData) {
	bytesRepresentation, err := json.Marshal(treblleInfo)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, Config.ServerURL, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		if Config.Logger != nil {
			Config.Logger.Printf("treblle: error creating request: %v", err.Error())
		}
	}

	// Set the content type from the writer, it includes necessary boundary as well
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", Config.APIKey)

	// Do the request
	client := &http.Client{Timeout: timeoutDuration}
	_, err = client.Do(req)
	if err != nil {
		if Config.Logger != nil {
			Config.Logger.Printf("treblle: error making request: %v", err.Error())
		}
	}
}
