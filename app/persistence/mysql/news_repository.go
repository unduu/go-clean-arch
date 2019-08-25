package mysql

import (
	"database/sql"
	// "fmt"

	"github.com/unduu/ibadf/app/domain/model"
)

type newsRepository struct {
	conn    *sql.DB
}

func NewNewsRepository() *newsRepository {
	return &newsRepository{
		conn: db,
	}
}

func (r *newsRepository) FindAll() ([]*model.News, error) {
  	rows, err := r.conn.Query("SELECT id, news_title, news_published, news_image, news_preview FROM news")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    newsList := make([]*model.News, 0)
    for rows.Next() {
		n := new(model.News)

		err = rows.Scan(
			&n.Id,
			&n.Title,
			&n.PublishedDate,
			&n.Image,
			&n.Preview,
		)
		if err != nil {
            return nil, err
        }

		newsList = append(newsList, n)
    }

    return newsList, nil
}