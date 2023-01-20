# Микросервис товаров
## Структура данных
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

### Создание таблиц

- Справочник единиц измерения:


```
CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    name VARCHAR(15) NOT NULL
)
```

- Справочник товаров:

```
CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    unit_id INT NOT NULL,
    FOREIGN KEY (unit_id) REFERENCES units (id)
)
```
