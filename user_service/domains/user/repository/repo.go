package activityrepo

import (
	usercore "managerservice/domains/user/core"
	usermodel "managerservice/domains/user/model"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Insert(userCore usercore.Core) (usercore.Core, error) {
	model := usermodel.ToModel(userCore)
	tx := r.db.Create(&model)

	if tx.Error != nil {
		return usercore.Core{}, tx.Error
	}

	return usermodel.ToCore(model), nil
}

func (r *userRepo) Update(userCore usercore.Core) (usercore.Core, error) {
	model := usermodel.ToModel(userCore)
	tx := r.db.Model(usermodel.User{}).Where("id", userCore.Id).Updates(&model)
	if tx.Error != nil {
		return usercore.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		return usercore.Core{}, gorm.ErrRecordNotFound
	}

	tx = r.db.Model(usermodel.User{}).First(&model)
	if tx.Error != nil {
		return usercore.Core{}, tx.Error
	}

	return usermodel.ToCore(model), nil
}

func (r *userRepo) GetByEmail(userCore usercore.Core) (bool, error) {
	model := usermodel.ToModel(userCore)
	tx := r.db.Model(usermodel.User{}).Where("email", userCore.Email).First(&model)
	if tx.Error != nil {
		return false, tx.Error
	}

	if tx.RowsAffected < 1 {
		return false, gorm.ErrRecordNotFound
	}

	tx = r.db.Model(usermodel.User{}).First(&model)
	if tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
