package models

import (
	"database/sql"
	"errors"
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
	stmt := `INSERT INTO texts (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TextModel) Get(id int) (Text, error) {
	stmt := `SELECT id, title, content, created FROM texts
		WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	var t Text
	err := row.Scan(&t.ID, &t.Title, &t.Content, &t.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Text{}, ErrNoRecord
		} else {
			return Text{}, err
		}
	}

	return t, nil
}

func (m *TextModel) Latest() ([]Text, error) {
	stmt := `SELECT id, title, content, created FROM texts
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var texts []Text
	for rows.Next() {
		var t Text
		err = rows.Scan(&t.ID, &t.Title, &t.Content, &t.Created)
		if err != nil {
			return nil, err
		}
		texts = append(texts, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return texts, nil
}
