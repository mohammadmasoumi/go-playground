package main

import (
	"encoding/json"
	"fmt"
)

type MyCareer struct {
	Job     string `json:"job"`
	Company string `json:"company"`
}

type MyType struct {
	Name   string   `json:"name"`
	Career MyCareer `json:"career"`
}

func main() {
	var myJson = `{"name": "asghar", "career": {"job": "sw", "company": "arvan"}}`
	var m MyType

	err := json.Unmarshal([]byte(myJson), &m)

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(m)
}
