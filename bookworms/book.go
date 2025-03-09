package main

type Book struct {
	Author string `json:"author"`
	Name   string `json:"name"`
}

type BookCount map[Book]uint
