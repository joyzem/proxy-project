package implementation

import (
	"database/sql"

	"github.com/joyzem/proxy-project/services/product/backend/repo"
	"github.com/joyzem/proxy-project/services/product/domain"
)

// Реализация репозитория товаров
type repository struct {
	db *sql.DB
}

// Возвращает репозиторий товаров
func NewProductRepo(db *sql.DB) repo.ProductRepo {
	return &repository{
		db: db,
	}
}

// Обращается к базе данных и добавляет новый товар
func (r *repository) CreateProduct(p domain.Product) (*domain.Product, error) {
	// SQL-запрос для добавления товара с возвратом id
	sql := `
			INSERT INTO products (name, price, unit_id)
			VALUES ($1, $2, $3) RETURNING id
			`

	// Подготовка запроса
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	// Выполнение запроса и сканирование результата (ID созданного товара)
	var insertedId int
	if err := result.QueryRow(p.Name, p.Price, p.Unit.Id).Scan(&insertedId); err != nil {
		return nil, err
	}

	// SQL-запрос для получения данных нового товара
	getProductSql := `
			SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id WHERE p.id = $1	
	`

	// Инициализация переменной с новым товаром
	insertedProduct := domain.Product{}
	
	// Выполнение запроса и сканирование результата (данные нового товара)
	if err := r.db.QueryRow(getProductSql, insertedId).Scan(
		&insertedProduct.Id,
		&insertedProduct.Name,
		&insertedProduct.Price,
		&insertedProduct.Unit.Id,
		&insertedProduct.Unit.Name); err != nil {

		return nil, err
	}
	// Возврат указателя на новый товар
	return &insertedProduct, nil
}

// GetProducts выполняет запрос в базу данных для получения списка всех продуктов.
// Возвращает массив типа domain.Product и ошибку, если она произошла.
func (r *repository) GetProducts() ([]domain.Product, error) {
	// SQL-запрос для получения информации о всех продуктах и их единицах измерения
	sql := `
	SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id ORDER BY p.name ASC
	` 
	
	// Выполнение запроса
	rows, err := r.db.Query(sql)
	if err != nil {
		// Возврат пустого списка и ошибки, если запрос не может быть выполнен
		return []domain.Product{}, err
	}

	// Освобождение ресурсов
	defer rows.Close()

	// Создание пустого списка продуктов
	products := []domain.Product{}

	// Перебор всех строк в ответе на запрос
	for rows.Next() {
		// Создание пустой структуры продукта
		product := domain.Product{}
		// Заполнение структуры данными из текущей строки
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Unit.Id, &product.Unit.Name)
		// Добавление продукта в список продуктов
		products = append(products, product)
	}

	// Возврат продуктов
	return products, nil
}

// Обновляет информацию о товаре в базе данных. 
// Принимает параметр p типа domain.Product, который содержит новые данные для товара.
func (r *repository) UpdateProduct(p domain.Product) (*domain.Product, error) {
	// SQL-запрос для обновления цены, единицы измерения и имени товара
	sql := `
			UPDATE products SET price = $1, unit_id = $2, name = $3 WHERE id = $4 RETURNING id
	`
	// Подготовка запроса
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	// Выполнение запроса и сканирование результата (ID обновленного товара)
	var id int
	if err := result.QueryRow(p.Price, p.Unit.Id, p.Name, p.Id).Scan(&id); err != nil {
		return nil, err
	}

	// SQL-запрос для получения данных обновленного товара
	getProductSql := `
			SELECT p.id, p.name, p.price, u.id, u.name FROM products p INNER JOIN units u ON u.id = p.unit_id WHERE p.id = $1	
	`

	// Инициализация переменной с обновленным товаром
	updatedProduct := domain.Product{}

	// Выполнение запроса и сканирование результата (данные обновленного товара)
	if err := r.db.QueryRow(getProductSql, id).Scan(
		&updatedProduct.Id,
		&updatedProduct.Name,
		&updatedProduct.Price,
		&updatedProduct.Unit.Id,
		&updatedProduct.Unit.Name); err != nil {
		return nil, err
	}

	// Возврат указателя на обновленный товар
	return &updatedProduct, nil
}

// DeleteProduct выполняет удаление продукта из базы данных по его ID.
// ID продукта передается в качестве параметра.
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
