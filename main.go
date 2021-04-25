package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// StringInt create a type alias for type string
type StringInt int

// UnmarshalJSON create a custom unmarshal for the StringInt
/// this helps us check the type of our value before unmarshalling it

func (st *StringInt) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into a int we convert it
	///otherwise we return an error
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		///here convert the string into
		///an integer
		i, err := strconv.Atoi(v)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err

		}
		*st = StringInt(i)

	}
	return nil
}

func main() {

	type Item struct {
		Name   string    `json:"name"`
		ItemId StringInt `json:"item_id"`
	}
	jsonData := []byte(`{"name":"item 1","item_id":"30"}`)
	var item Item
	err := json.Unmarshal(jsonData, &item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", item)

}



