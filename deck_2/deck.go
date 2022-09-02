/* using type def */

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
By doing do deck is equivalent to slice of string
*/
type deck []string

func (d deck) print() {
	for i, elem := range d {
		fmt.Println(i, elem)
	}
}

func newCard() (cards deck) {
	cardFace := [4]string{"Spades", "Diamond", "Heart", "Clubs"}
	cardValue := [4]string{"Ace", "Two", "Three", "Four"}

	for _, face := range cardFace {
		for _, value := range cardValue {
			cards = append(cards, face+" of "+value)
		}
	}
	return
}

func deal(d deck, handsNumber int) (deck, deck) {
	return d[:handsNumber], d[handsNumber:]
}

func (d deck) toString() (singleString string) {
	//1.Convert deck to type []string
	temp := []string(d)

	//2.Convert []string to string. for thsi u haveto import string package
	singleString = strings.Join(temp, ",")
	return

}

func (d deck) saveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) (d deck) {
	data, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Kuch Gadbad hai", error)
		os.Exit(1)
	}
	//converting byte to String
	strdata := string(data)

	//becoz deck itself is an []string thus we converted it directly
	d = strings.Split(strdata, ",")
	return
}

func (d deck) shuffleCard() {
	length := len(d)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randomNum := r.Intn(length - 1)
		d[i], d[randomNum] = d[randomNum], d[i]
	}
	return
}
