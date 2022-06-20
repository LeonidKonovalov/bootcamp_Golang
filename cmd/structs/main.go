package main

import "fmt"

type Addres struct {
	City   string
	Street string
}

type Person struct {
	Name   string
	Email  string
	Addres Addres // embeding встраивание типа
}

func (p *Person) printer() {
	// функция принадлежащая структуре Person
	// * указатель
	fmt.Println(p.Name, p.Email)
}

func (p Person) Printer() {
	// func (p person) Printer() {
	// инкапсульровано Регистром

	fmt.Println(p.Name, p.Email, p.Addres.City, p.Addres.Street)
}

func main() {
	p := Person{Name: "John", Email: "blabla@mail.ru", Addres: Addres{City: "Moscow", Street: "Lenina"}}
	p.printer()

}
