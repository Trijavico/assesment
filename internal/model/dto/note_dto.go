package dto

type Note struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Archived bool   `json:"archived"`
	UserID   uint   `json:"userId"`
}

type CreateNote struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Archived bool   `json:"archived"`
	UserID   uint   `json:"user_id"`
}
