package main

import (
	"log"
	"time"
)

type Person struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var str string = SaySomeThing("Aman")
	log.Println(str)
	var intro string
	intro, _ = SaySomeMore("Hello") //* we need to use two variables as it returns two values [if we need to ignore something use "_" ]
	log.Println(intro)

	//Struct
	person1:=Person{
		FirstName: "Aman",
		LastName: "Jain",
		PhoneNumber: "8954547xxx",
		Age: 22,
		BirthDate: time.Now(),
	}

	person2:=Person{"Robin","Subaski","7884454xxx",30,time.Now()}

	log.Println(person1)
	log.Println(person2)

	var myvar MyStruct 
	myvar.FirstName="hehe"

	log.Println(myvar.PrintFirstName())


	// Maps
	MyMap:= make(map[string]string)
	MyMap["Name"]="Aman"
	MyMap["Age"]="22"
	MyMap["Company"]="GreekSoft"
	log.Println(MyMap["Company"])

	var MyMap2 = map[string]int{ // creating map with map literals
		"Age":2,
		"Name":256,
	}

	log.Println(MyMap2)

	var Mymap3 map[string]int // nil map we need to intialise it again with map or like literal in order to use it 
 
	log.Println(Mymap3)

	//Slice
	var MySlice []int
	MySlice=append(MySlice,1)
	MySlice=append(MySlice,2)
	
	MySlice2:= []int{3,4,5,6,7,8}

	log.Println(MySlice2)
}

func SaySomeThing(s string) string {
	return s
}

func SaySomeMore(s string) (string, string) { //* function returning more than one value we need to use () for multiple values
	return s, "world"
}

// receive struct with function
type MyStruct struct {
	FirstName string
}

func (m *MyStruct) PrintFirstName() string{
	return m.FirstName
} 