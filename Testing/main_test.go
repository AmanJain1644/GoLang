package main

import (
	
	"testing"
)

// this is the manual manner to do test
var tests = []struct{
	name string
	divident float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data",100.0,10.0,10.0,false},
	{"Invalid-data",100.0,0.0,0.0,true},
	{"valid-data2",100.0,5.0,20.0,false},


}

func TestDivison(t *testing.T){
	for _,tt:= range tests{
		got,err:=divide(tt.divident,tt.divisor)
		if tt.isErr{
			if err==nil {
				t.Error("expected an error but did not get one")
			}
		}else{
			if err!=nil{
				t.Error("got unexpected error", err.Error())
			}
		}

		if got != tt.expected{
			t.Errorf("expected %f but got %f , but %f, %f",tt.expected,got,tt.divident,tt.divisor)
		}
	}
}

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