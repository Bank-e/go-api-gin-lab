package services

import (
	"errors"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) error {
	if student.Name == "" {
		return errors.New("name is required")
	}
	if student.GPA < 0 || student.GPA > 4 {
		return errors.New("gpa must be between 0.00 and 4.00")
	}
	return s.Repo.Update(id, student)
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.Delete(id)
}
