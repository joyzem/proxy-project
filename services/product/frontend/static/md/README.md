# Микросервис товаров

## Перед тем, как начать

- <a href="/about-golang/">Язык программирования Go</a>
- <a href="/about-arch/">Микросервисная архитектура</a>
- <a href="/about-sql/">SQL</a>
- <a href="/about/postgresql/">PostgreSQL</a>
- ER-модель

## Содержание
- [ER-Модель](#er-модель)
- [SQL-Запросы](#sql-запросы)
- [API](#api)

## ER-Модель
В данном микросервисе задействованы две таблицы:
- Справочник товаров;
- Справочник единиц измерения.

<img src="/product/static/assets/products-service.svg">

### Справочник товаров

Справочник товаров представлен таблицей ниже:

Код | Наименование | Цена | Код единицы измерения
-- | -- | -- | --
1 | Гвозди | 100 | 1
2 | Болты | 4 | 2

### Справочник единиц измерения

Справочник единиц измерения представлен таблицей ниже:

Код | Наименование
-- | --
1 | кг
2 | шт

## SQL запросы

### Справочник единиц измерения

- Создание
```
CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    name VARCHAR(15) NOT NULL
)
```

- Добавление 
```
INSERT INTO units (name) VALUES ($1)
```

- Получение
```
SELECT * FROM units
```

- Обновление
```
UPDATE units SET name = $1 WHERE id = $2
```

- Удаление
```
DELETE FROM units WHERE id = $1	
```	

### Справочник товаров

- Создание
```
CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    unit_id INT NOT NULL,
    FOREIGN KEY (unit_id) REFERENCES units (id)
)
```

- Добавление
```
INSERT INTO products (price, unit_id, name) VALUES ($1, $2, $3)
```

- Получение
```
SELECT p.id, p.name, p.price, u.id, u.name 
    FROM products p 
    INNER JOIN units u 
    ON u.id = p.unit_id 
ORDER BY p.name ASC
```

- Обновление
```
UPDATE products SET price = $1, unit_id = $2, name = $3 WHERE id = $4
```

- Удаление
```
DELETE FROM products WHERE id = $1
```

## API

