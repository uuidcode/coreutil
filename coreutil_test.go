package coreutil

import (
	"fmt"
	"testing"
)

func TestCreateUuid(t *testing.T) {
	uuid := CreateUuid()
	fmt.Println(uuid)
}
