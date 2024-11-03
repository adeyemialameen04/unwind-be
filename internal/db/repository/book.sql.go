// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: book.sql

package repository

import (
	"context"
)

const findAllBooks = `-- name: FindAllBooks :many
SELECT
  id, title, author, plublished_date
FROM
  book
`

func (q *Queries) FindAllBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.Query(ctx, findAllBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.PlublishedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
