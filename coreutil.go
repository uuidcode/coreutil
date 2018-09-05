package coreutil

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateUuid() string {
	value := uuid.NewV4()
	uuid := value.String()
	uuid = strings.Replace(uuid, "-", "", -1)
	return uuid
}

func DeleteItemByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func ToJson(object interface{}) string {
	bytes, err := json.MarshalIndent(object, "", "    ")
	CheckErr(err)

	return string(bytes)
}
