// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Contact struct {
	Firstname string `json:"firstname"`
}

type NewContact struct {
	Firstname string `json:"firstname"`
}

type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Schema  string `json:"schema"`
}
