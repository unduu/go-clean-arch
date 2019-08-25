package mysql

import (
	"database/sql"
	// "fmt"

	"github.com/unduu/ibadf/app/domain/model"
)

type commentRepository struct {
	conn    *sql.DB
}

func NewCommentRepository() *commentRepository {
	return &commentRepository{
		conn: db,
	}
}

func (r *commentRepository) FindByNewsId(id int) ([]*model.Comment, error) {
  	rows, err := r.conn.Query("SELECT u.username, u.avatar, c.content FROM comments c JOIN users u ON c.user_id=u.id WHERE news_id = ?", id)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    commentList := make([]*model.Comment, 0)
    for rows.Next() {
		c := new(model.Comment)

		err = rows.Scan(
			&c.Username,
			&c.Avatar,
			&c.Content,
		)
		if err != nil {
            return nil, err
        }

		commentList = append(commentList, c)
    }

    return commentList, nil
}