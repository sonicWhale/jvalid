package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)


type User struct {
	Email string `json:"email,omitempty"`
}

func main() {

	var user = User{
		Email: "asd.asd.asdf@gmail.com",
	}

	goLoader := gojsonschema.NewGoLoader(user)
	//log.Println(goLoader)

	schemaDoc := gojsonschema.NewReferenceLoader("file:///Users/a/go/src/awesomeProject/schema.json")

	schema, _ := gojsonschema.NewSchema(schemaDoc)
	//log.Println(schema)
	result, _ := schema.Validate(goLoader)


	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}
