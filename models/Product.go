package models

type Product struct {
	Id             int64   `json:"id"`
	ModelName      string  `json:"modelName"`
	Version        string  `json:"version"`
	Price          float64 `json:"price"`
	Description    string  `json:"description"`
	ProductionDate string  `json:"productionDate"`
}

type Products []Product
