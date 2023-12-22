package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue"`
	NullIntValue    *int      `json:"nullIntValue"`
}
type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

func main() {
	otherInt := 4321
	data := &myJSON{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherInt,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}
