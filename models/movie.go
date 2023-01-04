package models

type Movie struct {
	Id          string `json: "id"`
	Title       string `json: "title"`
	Duraction   string `json: "duraction"`
	Description string `json: "description"`
}
