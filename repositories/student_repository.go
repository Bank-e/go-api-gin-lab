package repositories

import (
	"database/sql"

	"example.com/student-api/models"
)

type StudentRepository struct {
	DB *sql.DB
}

// GET ALL
func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query("SELECT id, name, major, gpa FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
		students = append(students, s)
	}
	return students, nil
}

// GET with id
func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	row := r.DB.QueryRow(
		"SELECT id, name, major, gpa FROM students WHERE id = ?",
		id,
	)

	var s models.Student
	err := row.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// POST (create)
func (r *StudentRepository) Create(s models.Student) error {
	_, err := r.DB.Exec(
		"INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)",
		s.Id, s.Name, s.Major, s.GPA,
	)
	return err
}

// PUT (update)
func (r *StudentRepository) Update(id string, s models.Student) error {
	result, err := r.DB.Exec(
		"UPDATE students SET name = ?, major = ?, gpa = ? WHERE id = ?",
		s.Name, s.Major, s.GPA, id,
	)
	if err != nil {
		return err
	}

	// เช็คว่ามีแถวไหนถูกแก้ไขไหม ถ้าไม่มีแปลว่าไม่พบ ID นี้
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// Delete (delete)
func (r *StudentRepository) Delete(id string) error {
	result, err := r.DB.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
