package models

type Subscription struct {
	Id         int    `json:"id" db:"id"` // 0 - Ультра; 1 - Премиум; 2 - Базовый
	Name       string `json:"name" db:"name" binding:"required"`
	Price      int    `json:"price" db:"price" binding:"required"`
	Duration   int    `json:"duration" db:"duration" binding:"required"`
	ChatLimit  int    `json:"chat_limit" db:"chat_limit" binding:"required"`   // Лимит сообщений в чате в день
	EssayLimit int    `json:"essay_limit" db:"essay_limit" binding:"required"` // Лимит сочинений в день
}
