package fonction_go

import (
	"encoding/json"
	"log"
)

func tabStringToJson(tab []string) []byte {
	stringJson, err := json.Marshal(tab)
	if err != nil {
		log.Fatal(err)
	}
	return stringJson
}

func jsonToTabString(tab []byte) []string {
	var tabString []string
	if len(tab) == 0 {
		return nil
	}
	err := json.Unmarshal(tab, &tabString)
	if err != nil {
		log.Println("Warning: invalid JSON tag value:", err, string(tab))
		return nil
	}
	return tabString
}
