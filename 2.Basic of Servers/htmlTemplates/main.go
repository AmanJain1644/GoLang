package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter,r*http.Request){
	renderTemplate(w,"home.page.html")
}

func About(w http.ResponseWriter,r* http.Request){
	renderTemplate(w,"about.page.html")
}

func renderTemplate(w http.ResponseWriter,templ string){
	parsedTemplate,_ := template.ParseFiles("./Templates/" + templ)
	err:=parsedTemplate.Execute(w,nil)
	if err!= nil{
		fmt.Println("error while parsing",err)
		return
	}
}
func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)

	http.ListenAndServe(":8080",nil)
}