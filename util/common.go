package util

import (
	"log"
	"encoding/json"
)

func JsonEncode(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	return string(jsonBytes)
}
