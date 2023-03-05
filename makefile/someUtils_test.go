package main

import "testing"

func TestIntMin(t *testing.T){
	answer := IntMin(1,-3)
	if answer != -3{
		t.Error("Test Failed in IntMin(1,-3), want 1, got", answer)
	}
	
}