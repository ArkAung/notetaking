package models

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
