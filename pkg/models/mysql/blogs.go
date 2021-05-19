package mysql

import (
	"database/sql"

	"github.com/Amaish/webTemplate/pkg/models"
)

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(title, author, content, expires string) (int, error) {
	stmt := `INSERT INTO blogs (title,author,content,created,expires) 
	VALUES(?,?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, author, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *BlogModel) Get(id int) (*models.Blog, error) {
	stmt := `SELECT id,title,author,content,created,expires FROM blogs
	WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Blog{}
	err := row.Scan(&s.ID, &s.Title, &s.Author, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *BlogModel) Latest() ([]*models.Blog, error) {
	stmt := `SELECT id,title,author,content,created,expires FROM blogs
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	blogs := []*models.Blog{}
	for rows.Next() {
		s := &models.Blog{}
		err = rows.Scan(&s.ID, &s.Title, &s.Author, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}
