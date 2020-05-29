package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/domain/entity"
	"github.com/k-ueki/AGPlus/server/domain/repository"
)

type (
	ReviewGetRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewReviewGetRepository(db *gorm.DB) repository.ReviewGetRepository {
	return &ReviewGetRepositoryImpl{DB: db}
}

func (r *ReviewGetRepositoryImpl) FindAll() ([]*entity.Review, error) {
	var rows []*entity.Review
	if err := r.DB.Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *ReviewGetRepositoryImpl) FindByClassID(classID int) ([]*entity.Review, error) {
	var rows []*entity.Review
	if err := r.DB.Where("class_id = ?", classID).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *ReviewGetRepositoryImpl) IsExist(userID, classID int) (bool, error) {
	var row entity.Review
	err := r.DB.Where("user_id = ? AND class_id = ?", userID, classID).First(&row).Error
	return ErrorToIsExist(err, "review(userId=%d,classId=%d)", userID, classID)
}
