package models

import (
	"encoding/json"
	"testing"
)

func TestAPIResponseMarshal(t *testing.T) {
	expectedResponse := APIResponse{
		Success: true,
		Data:    map[string]string{"message": "hi"},
	}

	data, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{"success":true,"data":{"message":"hi"}}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestAPIResponseUnmarshal(t *testing.T) {
	data := []byte(`{"success":true,"data":{"message":"hi"}}`)

	var response APIResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !response.Success || response.Data.(map[string]interface{})["message"] != "hi" {
		t.Errorf("Unexpected response: %+v", response)
	}
}
