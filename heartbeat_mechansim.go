package main

import (
	"fmt"
	"log"
	"net/http"
)

func sendHeartBeats(appName string, instanceId string) error {
	log.Println("Sending heart beat signal to instance id :", instanceId)
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8761/eureka/apps/%s/%s", appName, instanceId), nil)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to create the put request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to make the request to the eureka server %w", err)
	}

	defer resp.Body.Close()

	log.Println("Heartbeat signal response-status: ", resp.Status)

	if resp.StatusCode != 200 {
		return fmt.Errorf("ERROR: Getting status code: %s", resp.Status)
	}
	return nil
}
