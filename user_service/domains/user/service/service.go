package activityservice

import (
	"fmt"
	usercore "managerservice/domains/user/core"
	"managerservice/exceptions"

	"gorm.io/gorm"
)

type userService struct {
	repo usercore.IRepoUser
}

func New(repo usercore.IRepoUser) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(activityCore usercore.Core) usercore.Core {
	result, err := s.repo.Insert(activityCore)

	if err != nil {
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	return result
}

func (s *userService) Verify(activityCore usercore.Core) usercore.Core {

	isExist, err := s.repo.GetByEmail(activityCore)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	if isExist {
		activityCore.IsActive = true
		result, err := s.repo.Update(activityCore)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				msg := fmt.Sprintf("Activity with Email %s Not Found", activityCore.Email)
				panic(exceptions.NewNotFoundError(msg))
			} else {
				panic(exceptions.NewInternalServerError(err.Error()))
			}
		}

		return result
	}

	return usercore.Core{}
}
