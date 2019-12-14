package repo_impl

import (
	"backend-github-trending/banana"
	"backend-github-trending/db"
	"backend-github-trending/log"
	"backend-github-trending/model"
	"backend-github-trending/repository"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type GithubRepoImpl struct {
	sql *db.Sql
}

func NewGithubRepo(sql *db.Sql) repository.GithubRepo {
	return &GithubRepoImpl{
		sql: sql,
	}
}

func (g GithubRepoImpl) SelectRepoByName(context context.Context, name string) (model.GithubRepo, error) {
	var repo = model.GithubRepo{}
	err := g.sql.Db.GetContext(context, &repo,
		`SELECT * FROM repos WHERE name=$1`, name)

	if err != nil {
		if err == sql.ErrNoRows {
			return repo, banana.RepoNotFound
		}
		log.Error(err.Error())
		return repo, err
	}
	return repo, nil
}

func (g GithubRepoImpl) SaveRepo(context context.Context, repo model.GithubRepo) (model.GithubRepo, error) {
	// name, description, url, color, lang, fork, stars, stars_today, build_by, created_at, updated_at
	statement := `INSERT INTO repos(
					name, description, url, color, lang, fork, stars, 
 			        stars_today, build_by, created_at, updated_at) 
          		  VALUES(
					:name,:description, :url, :color, :lang, :fork, :stars, 
					:stars_today, :build_by, :created_at, :updated_at
				  )`

	repo.CreatedAt = time.Now()
	repo.UpdatedAt = time.Now()

	_, err := g.sql.Db.NamedExecContext(context, statement, repo)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return repo, banana.RepoConflict
			}
		}
		log.Error(err.Error())
		return repo, banana.RepoInsertFail
	}

	return repo, nil
}

func (g GithubRepoImpl) SelectRepos(context context.Context, limit int) ([]model.GithubRepo, error) {
	var repos []model.GithubRepo
	err := g.sql.Db.SelectContext(context, &repos, "SELECT * FROM repos ORDER BY updated_at ASC LIMIT 25")
	if err != nil {
		log.Error(err.Error())
		return repos, err
	}

	return repos, nil
}

func (g GithubRepoImpl) UpdateRepo(context context.Context, repo model.GithubRepo) (model.GithubRepo, error) {
	// name, description, url, color, lang, fork, stars, stars_today, build_by, created_at, updated_at
	sqlStatement := `
		UPDATE repos
		SET 
			stars  = :stars,
			fork = :fork,
			stars_today = :stars_today,
			build_by = :build_by,
			updated_at = :updated_at
		WHERE name = :name
	`

	result, err := g.sql.Db.NamedExecContext(context, sqlStatement, repo)
	if err != nil {
		log.Error(err.Error())
		return repo, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return repo, banana.RepoNotUpdated
	}
	if count == 0 {
		return repo, banana.RepoNotUpdated
	}

	return repo, nil
}
