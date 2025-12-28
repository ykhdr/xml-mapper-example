package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	registry := NewMapperRegistry()

	registry.Register(&UserMapper{})
	registry.Register(&ProductMapper{})
	registry.Register(&OrderMapper{})

	examples := []struct {
		file string
		mapperType string
	}{
		{"examples/user.xml", "user"},
		{"examples/product.xml", "product"},
		{"examples/order.xml", "order"},
	}

	for _, example := range examples {
		fmt.Printf("Mapping %s file\n", example.file)

		data, err := os.ReadFile(example.file)
		if err != nil {
			log.Printf("Error reading file %s: %v\n", example.file, err)
			continue
		}

		result, err := registry.MapXML(example.mapperType, data)
		if err != nil {
			log.Printf("Error mapping XML: %v\n", err)
			continue
		}

		fmt.Printf("Mapped result: %+v\n\n", result)
	}

	fmt.Println("Test custom mapping")
	customXML := `<user id="999">
		<name>Jane Smith</name>
		<email>jane@test.com</email>
		<age>25</age>
	</user>`

	result, err := registry.MapXML("user", []byte(customXML))
	if err != nil {
		log.Printf("Error mapping custom XML: %v\n", err)
	} else {
		fmt.Printf("Custom mapping result: %+v\n", result)
	}
}