package service

import (
	"testing"
)

func TestAuth_Verify(t *testing.T) {
	a := Auth{}
	token, err := a.Verify("admin", "abcd.1234")
	if err != nil {
		t.Errorf("Verify() error = %v", err)
		return
	}
	//fmt.Println(token)
	_ = token
}

func TestAuth_ParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzQyNDkzMTQ3LCJuYmYiOjE3NDI0MDY3NDcsImlhdCI6MTc0MjQwNjc0N30.msdTbyCo_vuQR2JxFMXGq96YlhTERZVOKNvQosM6oBs"
	a := Auth{}
	got, err := a.ParseToken(tokenString)
	if err != nil {
		t.Errorf("ParseToken() error = %v", err)
	}
	//types.PrintJson(got)
	_ = got
}
