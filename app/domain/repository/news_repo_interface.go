package repository

import "github.com/unduu/ibadf/app/domain/model"

type NewsRepository interface {
	FindAll() ([]*model.News, error)
}
