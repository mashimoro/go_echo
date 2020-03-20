package main

// import "testing"
import "testing"

func TestSayHi(t *testing.T) {
	actualResult := "Say hi from demo"
	if actualResult != "Say hi from demo" {
		t.Error("Error from dev")
	}
}
