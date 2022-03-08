// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type List struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	UserID string    `json:"userId"`
	Items  []*string `json:"items"`
}

type NewList struct {
	Title  string    `json:"title"`
	UserID string    `json:"userId"`
	Items  []*string `json:"items"`
}

type NewUser struct {
	Name string  `json:"name"`
	City *string `json:"city"`
}

type User struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	City *string `json:"city"`
}
