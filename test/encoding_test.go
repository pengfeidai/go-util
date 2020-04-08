package main

import (
	"fmt"
	"json/json"
	"testing"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Addresss string `json:"address"`
}

func TestEncode(t *testing.T) {
	user := &User{"mochu", 18, "hangzhou"}
	_, err := json.Encode(user)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDecode(t *testing.T) {
	s := `{"name":"mochu","age":18,"address":"hangzhou"}`
	var user User
	if err := json.Decode([]byte(s), &user); err != nil {
		t.Error(err.Error())
	}
	fmt.Println("user name:", user.Name) // mochu
}
