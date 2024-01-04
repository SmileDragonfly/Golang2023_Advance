package main

import "fmt"

func PrintTypeAssertion(val interface{}) {
	switch val.(type) {
	case int:
		println("int")
		newVal := val.(int)
		println("NewVal:", newVal)
	case string:
		println("string")
		newVal := val.(string)
		println("NewVal:", newVal)
	case bool:
		println("bool")
		newVal := val.(bool)
		println("NewVal:", newVal)
	default:
		fmt.Printf("Type: %T", val)
	}

}
