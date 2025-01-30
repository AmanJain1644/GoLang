package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name":"Clark",
			"last_name":"Kent",
			"hair_color":"black",
			"has_dog":true		
		},
		{
			"first_name":"Bruce",
			"last_name":"wayne",
			"hair_color":"black",
			"has_dog":false	
		}
	
	]`

	var unMarshalled []Person
	err := json.Unmarshal([]byte(myJson),&unMarshalled)
	if err!= nil{
		log.Println("Error unmarshalling json", err)			
	}

	log.Printf("unmarshalled: %v",unMarshalled);

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Aman"
	m1.LastName = "Jain"
	m1.HairColor = "Black"
	m1.HasDog = false

	var m2 Person
	m2.FirstName = "Anmol"
	m2.LastName = "Pandey"
	m2.HairColor = "Black"
	m2.HasDog = true

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson,err:= json.MarshalIndent(mySlice,"","   ")
	if err!=nil {
		log.Println("marshalling err:",err)
	}

	fmt.Println(string(newJson))
}