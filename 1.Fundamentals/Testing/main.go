package main

import (
	"errors"
	"log"
)

func main() {
	result,err := divide(10,0)

	if(err!=nil){
		log.Println(err)
		return
	}

	log.Println("Result of division is ", result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("division with 0 is prohibited")
	}
	result = x / y

	return result,nil

}