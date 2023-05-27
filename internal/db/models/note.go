package models

type Notes struct {
	ID    uint
	Title string `validate:"required,lte=250"`
	Body  string `validate:"required,lte=2000"`
}
