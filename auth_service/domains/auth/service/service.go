package activityservice

import (
	authcore "authentication/domains/auth/core"
	"authentication/exceptions"
	"authentication/utils/helpers"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	repo authcore.IRepoAuth
}

func New(repo authcore.IRepoAuth) *authService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Verify(authCore authcore.Core) authcore.Core {

	result, err := s.repo.GetByEmail(authCore)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewBadRequestError("email or password is wrong"))
		}
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(authCore.Password))
	if err != nil {
		panic(exceptions.NewBadRequestError("email or password is wrong"))
	}

	token, err := helpers.CreateToken(int(result.Id), result.Role)
	if err != nil {
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	result.Token = token
	return result
}
