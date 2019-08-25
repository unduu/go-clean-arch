package model

type News struct {
	Id    int
	Title string
	PublishedDate string
	Image string
	Preview string
	Comments []*Comment
}

func NewNews(id int, title string, publishedDate string, image string, preview string) *News {
	return &News{
		Id:    id,
		Title: title,
		PublishedDate: publishedDate,
		Image: image,
		Preview: preview,
	}
}

func (n *News) SetComments(comments []*Comment) {
	n.Comments = comments
}
