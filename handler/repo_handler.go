package handler

import (
	"backend-github-trending/model"
	"backend-github-trending/repository"
	"github.com/labstack/echo"
	"net/http"
)

type RepoHandler struct {
	GithubRepo repository.GithubRepo
}

func (r RepoHandler) RepoTrending(c echo.Context) error {
	repos, _ := r.GithubRepo.SelectRepos(c.Request().Context(), 25)
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       repos,
	})
}
