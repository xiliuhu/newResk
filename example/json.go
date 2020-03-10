package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Id      int    `json:"id,string"`
	Name    string `json:"username"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"-"`
}

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	u := User{
		Id:      18,
		Name:    "etcd",
		Age:     0,
		Address: "沧州",
	}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	u2 := &User{}
	err = json.Unmarshal(data, u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", u)
	fmt.Printf("%+v \n", u2)
}
