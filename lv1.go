package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}
type dove interface {
	gugugu()
}
type repeater interface {
	repeat(string)
}
type lemon  interface{
	sour(string)
}

type smellgood interface {
	regret()
}
func (p *person) repeat(word string) {
	fmt.Println(word)
}

func (p *person) gugugu() {
	fmt.Println(p.name, "又鸽了")
}
func (p *person) sour (word string) {
	fmt.Println(p.name,"又酸了")
}
func (p *person) regret  () {
	fmt.Println(p.name,"上演王境泽真香定律")
}

func main() {

	p := &person{
		name:   "李越华",
		age:    17,
		gender: "sweet girl",
	}
	p.gugugu()
	var r repeater
	r = p
	r.repeat("God bless me!")

	var l lemon
	l = p
	l.sour("many times!")

	p.regret()





}