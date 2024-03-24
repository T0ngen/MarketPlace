package sqlc

import (
	"context"
	"database/sql"
	"fmt"

	"marketplace/pkg/common/config"

	_ "github.com/lib/pq"
)




func OpenPostgresConnection(conf config.Config) (*sql.DB, error) {
    connStr := "user=root password=1234 dbname=market sslmode=disable port=5435"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    return db, nil
}



const getGoodsByTitle2 = `-- name: GetGoodsByTitle :many
SELECT id, seller_id, title, price, description, image, category, rating, discount, status, created_at
FROM goods
WHERE title ILIKE '%' || $1 || '%';
`

func (q *Queries) GetGoodsByTitle(ctx context.Context, title string) ([]Good, error) {
	rows, err := q.db.QueryContext(ctx, getGoodsByTitle2, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Good
	for rows.Next() {
		var i Good
		if err := rows.Scan(
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
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}


