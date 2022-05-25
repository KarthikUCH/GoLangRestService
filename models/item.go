package models

type Item struct {
	ID          int    `json:"id"`
	Name        string `ijson:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ItemList struct {
	Items []Item `json:"items"`
}
