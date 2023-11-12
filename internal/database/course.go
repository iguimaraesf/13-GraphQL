package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryID string) (*Course, error) {
	query := "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	id := uuid.New().String()
	_, err := c.db.Exec(query, id, name, description, categoryID)
	if err == nil {
		return &Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  categoryID,
		}, nil
	}
	return nil, err
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: category_id})
	}
	return courses, nil
}
