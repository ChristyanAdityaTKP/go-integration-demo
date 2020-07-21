package integrationtest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dh258/go-integration-demo/controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateAddress(t *testing.T) {
	setupDB()
	gin.SetMode(gin.TestMode)

	err := clearTable()
	if err != nil {
		log.Fatalf("Failed to clear table: %v", err)
	}

	testCases := []struct {
		input           map[string]interface{}
		statusCode      int
		expectedName    string
		expectedCountry string
		errorMessage    string
	}{
		{
			input: map[string]interface{}{
				"name":    "Rumah Besar",
				"country": "Indonesia",
			},
			statusCode:      201,
			expectedName:    "Rumah Besar",
			expectedCountry: "Indonesia",
		},
		{
			input: map[string]interface{}{
				"name":    "",
				"country": "Indonesia",
			},
			statusCode:   422,
			errorMessage: "Please enter a valid name",
		},
		{
			input: map[string]interface{}{
				"name":    "Test name",
				"country": "",
			},
			statusCode:   422,
			errorMessage: "Please enter a valid country",
		},
		{
			input: map[string]interface{}{
				"name": 123,
			},
			statusCode:   422,
			errorMessage: "JSON cannot be processed",
		},
	}

	for _, tc := range testCases {
		// Prepare routes for test
		r := gin.Default()
		r.POST("/addresses", controllers.CreateAddress)

		// Create request and serve
		jsonData, err := json.Marshal(&tc.input)
		if err != nil {
			t.Errorf("Failed to create JSON input: %v\n", err)
		}
		req, _ := http.NewRequest(http.MethodPost, "/addresses", bytes.NewBuffer(jsonData))
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		// Fetch response and retrieve the body
		response := map[string]interface{}{}
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Failed to retrieve response: %v\n", err)
		}

		assert.Equal(t, recorder.Code, tc.statusCode)

		if tc.statusCode == 201 {
			assert.Equal(t, response["name"], tc.expectedName)
			assert.Equal(t, response["country"], tc.expectedCountry)
		}

		switch tc.statusCode {
		case 400, 422, 500:
			assert.Equal(t, response["message"], tc.errorMessage)
		}

	}
}
