package main

import "fmt"

type Addres struct {
	City   string
	Street string
}

type Person struct {
	Name   string
	Email  string
	Addres Addres
}

func (p *Person) printer() {
	fmt.Println(p.Name, p.Email)
}

func (p Person) Printer() {
	fmt.Println(p.Name, p.Email, p.Addres.City, p.Addres.Street)
}

func main() {
	p := Person{Name: "John", Email: "blabla@mail.ru", Addres: Addres{City: "Moscow", Street: "Lenina"}}
	p.printer()

}
