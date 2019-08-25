package repository

import "github.com/unduu/ibadf/app/domain/model"

type CommentRepository interface {
	FindByNewsId(newsId int) ([]*model.Comment, error)
}
