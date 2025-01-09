package models

import (
	"database/sql"
	"time"
)

type Text struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type TextModel struct {
	DB *sql.DB
}

func (m *TextModel) Insert(title string, content string) (int, error) {
	return 0, nil
}

func (m *TextModel) Get(id int) (Text, error) {
	return Text{}, nil
}

func (m *TextModel) Latest() ([]Text, error) {
	return nil, nil
}
