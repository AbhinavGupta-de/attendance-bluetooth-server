package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/service"

	"github.com/labstack/echo/v4"
)

var charUUID string
var serUUID string
var authToken string

func Index(c echo.Context) error {
	postURL := "https://5a43-106-51-13-146.ngrok-free.app/attendance/recievedatafromextserver/"
	jsonData := []byte(`{"key": "values"}`)

	response, err := service.PostData(postURL, jsonData)
	if err != nil {
		// Print an error message to the console
		fmt.Println("Error during POST request:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Parse the JSON response
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal([]byte(response), &jsonResponse); err != nil {
		fmt.Println("Error parsing response:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Check if the response has the "status" field set to "success"
	status, ok := jsonResponse["status"].(string)
	if !ok || status != "success" {
		fmt.Println("Error: Response status is not success")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Extract values from the "data" field
	data, ok := jsonResponse["data"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Data field is missing or not a map")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	charUUID, _ = data["characterUUID"].(string)
	serUUID, _ = data["serverUUID"].(string)
	authToken, _ = data["authToken"].(string)

	fmt.Println(charUUID, " ", serUUID, " ", authToken)

	fmt.Println(response)

	return c.String(http.StatusOK, "Hello, this is the test endpoint of the server.")
}
