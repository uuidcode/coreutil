package coreutil

import (
	"github.com/satori/go.uuid"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateUuid() string {
	value, err := uuid.NewV4()
	CheckErr(err)
	return value.String()
}

func DeleteItemByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
