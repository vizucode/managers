package activityrepo

import (
	authcore "authentication/domains/auth/core"
	authmodel "authentication/domains/auth/model"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *authRepo {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) GetByEmail(authCore authcore.Core) (authcore.Core, error) {
	model := authmodel.User{}
	tx := r.db.Model(authmodel.User{}).Where("email", authCore.Email).First(&model)
	if tx.Error != nil {
		return authcore.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		return authcore.Core{}, gorm.ErrRecordNotFound
	}

	core := authmodel.ToCore(model)
	core.Id = model.ID
	return core, nil
}
