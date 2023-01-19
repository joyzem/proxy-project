package domain

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  Unit   `json:"unit"` // Единицы измерения
}