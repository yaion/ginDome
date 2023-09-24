package dto

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUserJson(t *testing.T) {
	str := "{\n    \"user_name\": \"john_doe\",\n    \"nick_name\": \"John\",\n    \"phone\": \"12345678901\",\n    \"page_size\": 10,\n    \"sort\": {\n        \"user_name\": \"asc\",\n        \"nick_name\": \"desc\",\n        \"phone\": \"desc\"\n    },\n    \"page\": 1\n}"
    params:=UserListParams{}
	err:=json.Unmarshal([]byte(str),&params)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(params)
}
