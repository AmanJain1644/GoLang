package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"This is the Home Page dude!!")
}

func About(w http.ResponseWriter,r *http.Request){
	_,_ = fmt.Fprintf(w,"This this the About Page and 2+2 is %d",addValues(2,2))
}

func addValues(x,y int)int{
	return x+y
}

func Divide(w http.ResponseWriter,r *http.Request){
	n,err:=divide(100.0,0.0)
	if err!=nil{
		fmt.Fprintf(w,"Error %s",err.Error())
		return
	}
	fmt.Fprintf(w,"on dividing %f with %f we get %f",100.0,0.0,n)
}

func divide(x,y float32)(float32,error){
	if y==0{
		return 0,errors.New("can not divide with zero")
	}
	return x/y,nil
}

func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	http.HandleFunc("/divide",Divide)

	fmt.Println(fmt.Sprintf("Started listening on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}
