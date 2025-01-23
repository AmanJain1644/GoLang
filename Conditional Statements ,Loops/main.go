package main

import "log"

func main() {

	isTrue := true
	something:= true
	myVar:= "Cat"

	if isTrue {
		log.Println("it's true")
	}else if  isTrue && something{
		log.Println("it's else if here ")
	}else{
		log.Println("it's false")
	}

	switch myVar{ //* no need to use break go will do it its self
	case "cat":
		log.Println("var is set to cat")
	case "dog":
		log.Println("var is set to dog")
	default:
		log.Println("var is something else!!")
	}
}