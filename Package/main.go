package main

import (
	"log"

	"github.com/amanjain1644/GoLang/Package/helpers"
)

func main(){
	var myVar helpers.SomeType
	myVar.TypeName = "SomeName"
	myVar.TypeNumber = 99

	log.Println(myVar.TypeName,myVar.TypeNumber)
}