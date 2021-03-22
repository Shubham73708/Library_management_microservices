package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]interface{})
	myMap["key"] = 0.25
	myMap["key2"] = "some string"
	fmt.Printf("%+v\n", myMap)
	// fetch value using type assertion
	fmt.Println(myMap["key"].(float64))
	fetchValue(myMap)
}

func fetchValue(myMap map[string]interface{}) {
	for _, value := range myMap {
		switch v := value.(type) {
		case string:
			fmt.Println("the value is string =", value.(string))
		case float64:
			fmt.Println("the value is float64 =", value.(float64))
		case interface{}:
			fmt.Println(v)
		default:
			fmt.Println("unknown")
		}
	}
}
