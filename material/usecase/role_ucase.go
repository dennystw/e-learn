package usecase

import (
	"context"
	"time"

	"example/e-learn/domain"
)

type roleUsecase struct {
	roleRepo       domain.RoleRepository
	contextTimeout time.Duration
}

// NewRoleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewRoleUsecase(a domain.RoleRepository, timeout time.Duration) domain.RoleUsecase {
	return &roleUsecase{
		roleRepo:       a,
		contextTimeout: timeout,
	}
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */

func (a *roleUsecase) Store(c context.Context, m *domain.Role) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err = a.roleRepo.Store(ctx, m)
	return
}
