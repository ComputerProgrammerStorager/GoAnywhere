package main

import "time"

// like many function programming languanges, Go support struct to group different types of
// data as an entity
// The following is classic employ struct type
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var james Employee

// not like member_of_pointer operator (->), Go uses dot(.) to get member from a pointer
