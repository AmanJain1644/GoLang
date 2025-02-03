package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name  string
	Breed string
}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func main() {
	dog1 := Dog{
		Name:  "doggy",
		Breed: "German Shepherd",
	}
	PrintInfo(dog1);

}

func PrintInfo(a Animal) {
	log.Println("This Animal says",a.Says(),"it has",a.NumberOfLegs(),"legs");
}