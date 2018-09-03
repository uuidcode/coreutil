package coreutil

import (
	"github.com/satori/go.uuid"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateUuid() string {
	value, err := uuid.NewV4()
	CheckErr(err)
	uuid := value.String()
	uuid = strings.Replace(uuid, "-", "", -1)
	return uuid
}

func DeleteItemByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
