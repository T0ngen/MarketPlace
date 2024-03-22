-- name: GetGoodsByTitleOLD :many
SELECT * FROM goods
WHERE title ILIKE $1;




-- name: CreateGoods :one
INSERT INTO goods (
    seller_id, title, price, description, image, category, rating,
    discount, status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: DeleteGoodsById :exec
DELETE FROM goods
WHERE id = $1;



