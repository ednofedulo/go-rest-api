package main

type Person struct {
	_id string `json:"id,omitempty"`
	FirstName	string `json:"firstname,omitempty"`
	LastName	string `json:"lastname,omitempty"`
	Email	string `json:"email,omitempty"`
	Age	int `json:"age,omitempty"`
}