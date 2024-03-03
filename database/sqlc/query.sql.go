// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package sqlc

import (
	"context"
)

const createGoods = `-- name: CreateGoods :one
INSERT INTO goods (
    seller_id, title, price, description, image, category, rating,
    discount, status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, seller_id, title, price, description, image, category, rating, discount, status, created_at
`

type CreateGoodsParams struct {
	SellerID    int64
	Title       string
	Price       int64
	Description string
	Image       string
	Category    string
	Rating      string
	Discount    int64
	Status      string
}

func (q *Queries) CreateGoods(ctx context.Context, arg CreateGoodsParams) (Good, error) {
	row := q.db.QueryRowContext(ctx, createGoods,
		arg.SellerID,
		arg.Title,
		arg.Price,
		arg.Description,
		arg.Image,
		arg.Category,
		arg.Rating,
		arg.Discount,
		arg.Status,
	)
	var i Good
	err := row.Scan(
		&i.ID,
		&i.SellerID,
		&i.Title,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.Category,
		&i.Rating,
		&i.Discount,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteGoodsById = `-- name: DeleteGoodsById :exec
DELETE FROM goods
WHERE id = $1
`

func (q *Queries) DeleteGoodsById(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteGoodsById, id)
	return err
}

const getGoodsByTitle = `-- name: GetGoodsByTitle :one
SELECT id, seller_id, title, price, description, image, category, rating, discount, status, created_at FROM goods
WHERE title = $1
`

func (q *Queries) GetGoodsByTitle(ctx context.Context, title string) (Good, error) {
	row := q.db.QueryRowContext(ctx, getGoodsByTitle, title)
	var i Good
	err := row.Scan(
		&i.ID,
		&i.SellerID,
		&i.Title,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.Category,
		&i.Rating,
		&i.Discount,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
