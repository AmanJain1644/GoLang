package helpers

import "math/rand"

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(n int) int {
	return rand.Intn(n);
}