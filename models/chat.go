package models

type Message struct {
	ID      int    `db:"id" json:"id"`
	ChatID  string `db:"chat_id" json:"chat_id"`
	Role    string `db:"role" json:"role"` // "user" или "assistant"
	Content string `db:"content" json:"content"`
}

type Chat struct {
	ID       string    `db:"id" json:"id"`
	UserId   string    `db:"user_id" json:"user_id"`
	TaskID   int       `db:"task_id" json:"task_id"`
	Messages []Message `json:"messages"`
}
