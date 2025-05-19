package models

type EssayTheme struct {
	Id    int    `json:"id"`
	Theme string `json:"theme"`
	Text  string `json:"text"`
}
