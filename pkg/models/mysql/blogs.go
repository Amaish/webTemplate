package mysql

import (
	"database/sql"

	"github.com/Amaish/webTemplate/pkg/models"
)

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(title, author, content, expires string) (int, error) {
	return 0, nil
}

func (m *BlogModel) Get(id int) (*models.Blog, error) {

}
