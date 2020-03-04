package main

type ContactList struct {
	id    int
	name  string
	email string
	number   int
	age int
}

type ContactListint interface {
	Add(a ContactList) error
	Update(name string, email string, number int,age int, id int) error
	Delete(id int) error
	GetContact(id int)  error
}
