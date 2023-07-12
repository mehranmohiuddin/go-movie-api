package model

type Movie struct {
	Name     string `json:"name"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}
