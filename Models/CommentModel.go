package Models

//comment struct
type Comment struct{
	ID	string	`json:"id"`
	Name	string	`json:"name"`
	Content	string	`json:"content"`
	PostID string `json:"postid"`
}