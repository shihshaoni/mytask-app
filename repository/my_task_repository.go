package repository

import (
	"mytask-app/model"

	"github.com/jinzhu/gorm"
)

type MyTaskRepository struct {
	DB *gorm.DB
}

func NewMyTaskRepository(db *gorm.DB) *MyTaskRepository {
	return &MyTaskRepository{DB: db}
}

func (r *MyTaskRepository) Create(MyTask *model.MyTask) (*model.MyTask, error) {
	err := r.DB.Create(MyTask).Error
	return MyTask, err
}

func (r *MyTaskRepository) FindByID(id uint) (*model.MyTask, error) {
	var MyTask model.MyTask
	err := r.DB.First(&MyTask, id).Error
	return &MyTask, err
}

func (r *MyTaskRepository) Update(MyTask *model.MyTask) (*model.MyTask, error) {
	err := r.DB.Save(MyTask).Error
	return MyTask, err
}

func (r *MyTaskRepository) Delete(MyTask *model.MyTask) error {
	err := r.DB.Delete(MyTask).Error
	return err
}
