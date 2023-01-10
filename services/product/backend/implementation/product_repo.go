package implementation

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"
	product "github.com/joyzem/proxy-project/services/product/backend"
	"github.com/joyzem/proxy-project/services/utils"
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewProductRepo(db *sql.DB, logger log.Logger) (product.ProductRepo, error) {
	return &repository{
		db:     db,
		logger: log.With(logger, "rep"),
	}, nil
}

func (r *repository) CreateProduct(ctx context.Context, p product.Product) (*product.Product, error) {
	sql := `
			INSERT INTO products (price, unit_id, name)
			VALUES ($1, $2, $3) RETURNING id
			`
	result, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(p.Price, p.Unit.Id, p.Name).Scan(&insertedId); err != nil {
		return nil, err
	}
	getProductSql := `
			SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id WHERE p.id = $1	
	`
	insertedProduct := product.Product{}
	if err := r.db.QueryRowContext(ctx, getProductSql, insertedId).Scan(
		&insertedProduct.Id,
		&insertedProduct.Name,
		&insertedProduct.Price,
		&insertedProduct.Unit.Id,
		&insertedProduct.Unit.Name); err != nil {

		utils.LogError(&r.logger, err)
		return nil, err
	}
	return &insertedProduct, nil
}

func (r *repository) GetProducts(ctx context.Context) ([]product.Product, error) {
	sql := `
			SELECT * FROM products INNER JOIN units ON units.id = products.unit_id
	`
	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		utils.LogError(&r.logger, err)
		return []product.Product{}, err
	}
	defer rows.Close()

	products := []product.Product{}

	for rows.Next() {
		product := product.Product{} // Current Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Unit.Id, &product.Unit.Id, &product.Unit.Name)
		if err != nil {
			utils.LogError(&r.logger, err)
			continue
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) UpdateProduct(ctx context.Context, p product.Product) (*product.Product, error) {
	sql := `
			UPDATE products SET price = $1, unit_id = $2, name = $3 WHERE id = $4 RETURNING id
	`
	result, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	defer result.Close()

	var id int
	if err := result.QueryRow(p.Price, p.Unit.Id, p.Name, p.Id).Scan(&id); err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	getProductSql := `
			SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id WHERE p.id = $1	
	`
	updatedProduct := product.Product{}
	if err := r.db.QueryRowContext(ctx, getProductSql, id).Scan(
		&updatedProduct.Id,
		&updatedProduct.Name,
		&updatedProduct.Price,
		&updatedProduct.Unit.Id,
		&updatedProduct.Unit.Name); err != nil {

		utils.LogError(&r.logger, err)
		return nil, err
	}
	return &updatedProduct, nil
}

func (r *repository) DeleteProduct(ctx context.Context, id int64) error {
	sql := `
			DELETE FROM products WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, sql, id)
	if err != nil {
		utils.LogError(&r.logger, err)
		return err
	}
	return nil
}
