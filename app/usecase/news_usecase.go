package usecase

import (
	"github.com/unduu/ibadf/app/domain/model"
	"github.com/unduu/ibadf/app/domain/repository"
)

type News struct {
	Id    int
	Title string
}

type NewsUsecase interface {
	ListNews() ([]*News, error)
}

type newsUsecase struct {
	newsRepo    repository.NewsRepository
	commentRepo repository.CommentRepository
}

func NewNewsUsecase(newsRepo repository.NewsRepository, commentRepo repository.CommentRepository) *newsUsecase {
	return &newsUsecase{
		newsRepo:    newsRepo,
		commentRepo: commentRepo,
	}
}

func (n *newsUsecase) ListNews() ([]*model.News, error) {
	newsArr, err := n.newsRepo.FindAll()
	if err != nil {
		return nil, err
	}

	for _, news := range newsArr {
		commentArr, err := n.commentRepo.FindByNewsId(news.Id)
		if err != nil {
			return nil, err
		}
		news.SetComments(commentArr)
	}

	return newsArr, nil
	// return toNews(news), nil
}

func toNews(newsArr []*model.News) []*News {
	res := make([]*News, len(newsArr))
	for i, news := range newsArr {
		res[i] = &News{
			Id:    news.Id,
			Title: news.Title,
		}
	}
	return res
}