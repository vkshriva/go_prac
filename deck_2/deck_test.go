/*
1: File Name should be name_test.go
2: unit test cases should be start with keyword "Test" and after that firstLetter should be \
   Capital "NewDeck"
3: run command go test to run all the test
4: t *testing.T k baare m baad m padengay

*/

package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newCard()
	if len(d) != 16 {
		t.Errorf("Expected 16 but you got %v", len(d))
	}
}
