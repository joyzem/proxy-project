package implementation

import (
	"database/sql"

	"github.com/joyzem/proxy-project/services/product/backend/repo"
	"github.com/joyzem/proxy-project/services/product/domain"
)

type repository struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) repo.ProductRepo {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProduct(p domain.Product) (*domain.Product, error) {
	sql := `
			INSERT INTO products (name, price, unit_id)
			VALUES ($1, $2, $3) RETURNING id
			`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(p.Name, p.Price, p.Unit.Id).Scan(&insertedId); err != nil {
		return nil, err
	}
	getProductSql := `
			SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id WHERE p.id = $1	
	`
	insertedProduct := domain.Product{}
	if err := r.db.QueryRow(getProductSql, insertedId).Scan(
		&insertedProduct.Id,
		&insertedProduct.Name,
		&insertedProduct.Price,
		&insertedProduct.Unit.Id,
		&insertedProduct.Unit.Name); err != nil {

		return nil, err
	}
	return &insertedProduct, nil
}

func (r *repository) GetProducts() ([]domain.Product, error) {
	sql := `
	SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id ORDER BY p.name ASC
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		return []domain.Product{}, err
	}
	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{} // Current Product
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Unit.Id, &product.Unit.Name)
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) UpdateProduct(p domain.Product) (*domain.Product, error) {
	sql := `
			UPDATE products SET price = $1, unit_id = $2, name = $3 WHERE id = $4 RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var id int
	if err := result.QueryRow(p.Price, p.Unit.Id, p.Name, p.Id).Scan(&id); err != nil {
		return nil, err
	}
	getProductSql := `
			SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id WHERE p.id = $1	
	`
	updatedProduct := domain.Product{}
	if err := r.db.QueryRow(getProductSql, id).Scan(
		&updatedProduct.Id,
		&updatedProduct.Name,
		&updatedProduct.Price,
		&updatedProduct.Unit.Id,
		&updatedProduct.Unit.Name); err != nil {

		return nil, err
	}
	return &updatedProduct, nil
}

func (r *repository) DeleteProduct(id int) error {
	sql := `
			DELETE FROM products WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
