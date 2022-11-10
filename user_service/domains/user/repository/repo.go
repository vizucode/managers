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

func (r *userRepo) Insert(userCore usercore.Core) error {
	model := usermodel.ToModel(userCore)
	tx := r.db.Create(&model)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *userRepo) Update(userCore usercore.Core) error {
	model := usermodel.ToModel(userCore)
	tx := r.db.Model(usermodel.User{}).Where("id", userCore.Id).Updates(&model)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	tx = r.db.Model(usermodel.User{}).First(&model)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *userRepo) Delete(activityCore usercore.Core) error {
	model := usermodel.User{}
	tx := r.db.Model(usermodel.User{}).Where("id", activityCore.Id).Delete(&model)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *userRepo) FindAll() ([]usercore.Core, error) {
	modelList := []usermodel.User{}
	tx := r.db.Model(usermodel.User{}).Find(&modelList)
	if tx.Error != nil {
		return []usercore.Core{}, tx.Error
	}
	coreList := []usercore.Core{}
	for _, value := range modelList {
		coreList = append(coreList, usermodel.ToCore(value))
	}

	return coreList, nil
}

func (r *userRepo) First(userCore usercore.Core) (usercore.Core, error) {
	model := usermodel.ToModel(userCore)
	tx := r.db.Model(usermodel.User{}).Where("id", userCore.Id).First(&model)
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
