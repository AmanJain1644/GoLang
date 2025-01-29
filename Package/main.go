package main

import (
	"log"

	"github.com/amanjain1644/GoLang/Package/helpers"
)

func GetTheNum(intChan chan int){
	intChan <- helpers.RandomNumber(10);

}

func main(){
	var myVar helpers.SomeType
	myVar.TypeName = "SomeName"
	myVar.TypeNumber = 99

	log.Println(myVar.TypeName,myVar.TypeNumber)

	//channels 
	intChan := make(chan int);
	defer close(intChan)

	go GetTheNum(intChan)

	num:= <-intChan
	log.Println(num)

}