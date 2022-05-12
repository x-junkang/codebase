package main

import (
	"encoding/json"
	"fmt"
)

type TrainResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func main() {
	resp := &Response{
		Code:    200,
		Message: "hello",
		Data:    map[string]interface{}{"hello": "100", "test": map[string]string{"hel": "gfsda"}},
	}
	tmp, err := json.Marshal(resp)
	check(err)
	fmt.Printf("%s\n", tmp)
	ans := &TrainResponse{}
	err = json.Unmarshal(tmp, ans)
	check(err)
	fmt.Println(ans)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
