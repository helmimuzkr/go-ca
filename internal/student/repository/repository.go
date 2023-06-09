package repository

import (
	"fmt"
	"wide_technologies/internal/student"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) student.Service {
	return &studentRepository{db: db}
}

func (ss *studentRepository) GetStudentByID(studentID int) (student.Core, error) {
	model := Student{}
	tx := ss.db.First(&model, studentID)
	if tx.Error != nil {
		return student.Core{}, fmt.Errorf("get student in student repository error: %w", tx.Error)
	}
	core := student.Core{
		ID:   int(model.ID),
		NPM:  model.NPM,
		Name: model.Name,
	}
	return core, nil
}

func (ss *studentRepository) Create(studentCore student.Core) error {
	model := Student{
		NPM:  studentCore.NPM,
		Name: studentCore.Name,
	}
	tx := ss.db.Create(&model)
	if tx.Error != nil {
		return fmt.Errorf("create student in student repository error: %w", tx.Error)
	}
	return nil
}

func (ss *studentRepository) GetList() ([]student.Core, error) {
	models := []Student{}
	tx := ss.db.Find(&models)
	if tx.Error != nil {
		return nil, fmt.Errorf("Get list student in student repository error: %w", tx.Error)
	}
	cores := []student.Core{}
	for _, v := range models {
		c := student.Core{
			ID:   int(v.ID),
			NPM:  v.NPM,
			Name: v.Name,
		}
		cores = append(cores, c)
	}
	return cores, nil
}

func (ss *studentRepository) Update(studentID int, studentCore student.Core) error {
	model := Student{
		NPM:  studentCore.NPM,
		Name: studentCore.Name,
	}
	tx := ss.db.Where("id = ?", studentID).Updates(&model)
	if tx.Error != nil {
		return fmt.Errorf("update student in student repository error: %w", tx.Error)
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("update student in student repository error: no rows affected")
	}
	return nil
}

func (ss *studentRepository) Delete(studentID int) error {
	tx := ss.db.Delete(&Student{}, studentID)
	if tx.Error != nil {
		return fmt.Errorf("delete student in student repository error: %w", tx.Error)
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("delete student in student repository error: no rows affected")
	}
	return nil
}
