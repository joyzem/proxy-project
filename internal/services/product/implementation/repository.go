package implementation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kit/log"
	Product "github.com/joyzem/proxy-project/internal/services/product"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB, logger log.Logger) (Product.Repository, error) {
	return &repository{
		db: db,
	}, nil
}

func (r *repository) CreateProduct(ctx context.Context, product Product.Product) error {
	sql := `
			INSERT INTO products (money, unit_id, name)
			VALUES ($1, $2, $3)
			`
	_, err := r.db.ExecContext(ctx, sql, product.Price, product.Unit.Id, product.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetProducts(ctx context.Context) ([]Product.Product, error) {
	sql := `
			SELECT * FROM products INNER JOIN units ON units.id = products.unit_id
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return []Product.Product{}, err
	}
	defer rows.Close()

	products := []Product.Product{}

	fmt.Println(rows)

	for rows.Next() {
		cProduct := Product.Product{} // Current Product
		err := rows.Scan(&cProduct.Id, &cProduct.Name, &cProduct.Price, &cProduct.Unit.Id, &cProduct.Unit.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, cProduct)
	}

	return products, nil
}

func (r *repository) GetProduct(ctx context.Context, id int) (Product.Product, error) {
	return Product.Product{}, nil
}

func (r *repository) UpdateProduct(ctx context.Context, id int, product Product.Product) error {
	sql := `
			UPDATE products SET price = $1, unit_id = $2, name = $3 where id = $4
	`
	_, err := r.db.ExecContext(ctx, sql, product.Price, product.Unit.Id, product.Name)
	if err != nil {
		fmt.Println(err) // TODO log
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
		fmt.Println(err)
		return err
	}
	return nil
}
