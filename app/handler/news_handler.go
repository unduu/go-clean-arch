package handler

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/unduu/ibadf/app/persistence/mysql"
	"github.com/unduu/ibadf/app/usecase"
	"github.com/unduu/ibadf/app/domain/model"
)

type NewsResponse struct {
	Id  int `json:"id"`
	Title string `json:"title"`
	PublishedDate  string `json:"published_date"`
	Image  string `json:"image"`
	Preview  string `json:"preview"`
	TotalLike  int `json:"total_like"`
	Comments  []*model.Comment `json:"comments"`
}

func GetNews(w http.ResponseWriter, r *http.Request, id string) {

	var response []NewsResponse

	newsid, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
	}

	newsRepo := mysql.NewNewsRepository()
	commentRepo := mysql.NewCommentRepository()
	newsUC, err := usecase.NewNewsUsecase(newsRepo, commentRepo), nil
	newsArr, err := newsUC.ListNews()

	for _, n := range newsArr {
		news := NewsResponse{
			Id:  newsid,
			Title: n.Title,
			PublishedDate:  "2019-08-08",
			Image:  n.Image,
			Preview:  n.Preview,
			TotalLike:  20,
			Comments:  n.Comments,
		}

		response = append(response, news)
	}

	respondSuccessJSON(w, response)
}