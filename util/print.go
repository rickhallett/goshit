package util

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(message ...interface{}) {
	for _, m := range message {
		jsonMessage, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(jsonMessage))
	}
}
