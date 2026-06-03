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
	err := json.Unmarshal([]byte(tab), &tabString)
	if err != nil {
		log.Fatal(err)
	}
	return tabString
}
