package main

import (
	"testing"
)

var (
	a ContactListint
)

func TestNewAnimalList(t *testing.T) {
	_, err = NewContactList()
	if err != nil {
		t.Error("NewContactList cannot be created")
	}
}

func Test_Add(t *testing.T) {
	a, err = NewContactList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}
	a.Add(ContactList{
		name:  "Asadbek",
		email: "asadbekbakhodirov2@gmail.com",
		number:   5,
		age : 22,
	}) 
}

func Test_Update(t *testing.T) {
	a, err = NewContactList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}

	a.Update(

		"Rustam",
		"Rustagram@gmail.com",
		998485484, 2,2,
	)
	
}
func Test_Delete(t *testing.T) {
	a, err = NewContactList()
	if err != nil {
		t.Error("NewContactList cannot be created")
	}

	a.Delete(
2,
	)
	
}
func Test_GetTask(t *testing.T) {
	a, err = NewContactList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}

	a.GetContact(
		2,
	)
}
