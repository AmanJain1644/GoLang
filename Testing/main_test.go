package main

import (
	"log"
	"testing"
)

// this is the manual manner to do test

func TestDivide(t *testing.T){
	_,err := divide(100.0,10.0)
	if err!=nil {
		t.Error("got unexpected error")
	}
}


func TestBadDivide(t *testing.T){
	_,err := divide(100.0,0)
	if err==nil {
		t.Error("we should have an error ")
	}
}