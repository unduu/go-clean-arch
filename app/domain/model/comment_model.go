package model

type Comment struct {
	Username string
	Avatar string
	Content string
}

func NewComment(username string, avatar string, content string) *Comment {
	return &Comment{
		Username: username,
		Avatar: avatar,
		Content: content,
	}
}