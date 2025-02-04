package main

import (
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

func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/About",About)

	fmt.Println(fmt.Sprintf("Started listening on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}
