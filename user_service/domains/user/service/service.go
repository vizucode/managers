package activityservice

import (
	usercore "managerservice/domains/user/core"
	"managerservice/exceptions"

	"golang.org/x/crypto/bcrypt"
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

func (s *userService) Create(activityCore usercore.Core) {
	password := activityCore.Password

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	activityCore.Password = string([]byte(hashPassword))

	err = s.repo.Insert(activityCore)

	if err != nil {
		panic(exceptions.NewInternalServerError(err.Error()))
	}
}

func (s *userService) Update(activityCore usercore.Core) {
	err := s.repo.Update(activityCore)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
		panic(exceptions.NewInternalServerError(err.Error()))
	}
}

func (s *userService) Delete(activityCore usercore.Core) {
	err := s.repo.Delete(activityCore)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
		panic(exceptions.NewInternalServerError(err.Error()))
	}
}

func (s *userService) GetAll() []usercore.Core {
	result, err := s.repo.FindAll()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	return result
}

func (s *userService) Get(activityCore usercore.Core) usercore.Core {
	result, err := s.repo.First(activityCore)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))
		}
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	return result
}
