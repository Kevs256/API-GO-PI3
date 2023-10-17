package models

type Claims struct {
	Id        int    `json:"id_user"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
