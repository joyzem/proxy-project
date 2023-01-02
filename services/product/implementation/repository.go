package implementation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joyzem/proxy-project/internal/services/product"
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) (product.Repository, error) {
	return &repository{
		db:     db,
		logger: log.With(logger, "rep"),
	}, nil
}

func (r *repository) CreateProduct(ctx context.Context, product product.Product) error {
	sql := `
			INSERT INTO products (money, unit_id, name)
			VALUES ($1, $2, $3)
			`
	_, err := r.db.ExecContext(ctx, sql, product.Price, product.Unit.Id, product.Name)
	if err != nil {
		level.Error(r.logger).Log("err", err.Error())
		return err
	}
	return nil
}

func (r *repository) GetProducts(ctx context.Context) ([]product.Product, error) {
	sql := `
			SELECT * FROM products INNER JOIN units ON units.id = products.unit_id
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		level.Error(r.logger).Log("err", err.Error())
		return []product.Product{}, err
	}
	defer rows.Close()

	products := []product.Product{}

	fmt.Println(rows)

	for rows.Next() {
		product := product.Product{} // Current Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Unit.Id, &product.Unit.Id, &product.Unit.Name)
		if err != nil {
			level.Error(r.logger).Log("err", err.Error())
			continue
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) GetProduct(ctx context.Context, id int) (product.Product, error) {
	return product.Product{}, nil
}

func (r *repository) UpdateProduct(ctx context.Context, id int, product product.Product) error {
	sql := `
			UPDATE products SET price = $1, unit_id = $2, name = $3 where id = $4
	`
	_, err := r.db.ExecContext(ctx, sql, product.Price, product.Unit.Id, product.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteProduct(ctx context.Context, id int) error {
	sql := `
			DELETE FROM products WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}
