package responses

import (
	"encoding/json"
	"log"
)

var NOT_FOUND []byte

func init() {
	var err error
	NOT_FOUND, err = json.Marshal(map[string]string{"error": "not found"})
	if err != nil {
		log.Fatalf("Failed to marshal NOT_FOUND: %v", err)
	}
}
