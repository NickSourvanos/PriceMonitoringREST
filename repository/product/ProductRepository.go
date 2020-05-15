package product

import (
	"PriceMonitoringService/models"
	"database/sql"
	"fmt"
)

type productRepositoryPostgres struct {
	db *sql.DB
}

func NewProductRepositoryPostgres(db *sql.DB) *productRepositoryPostgres {
	return &productRepositoryPostgres{db}
}

func (p *productRepositoryPostgres) Save(product *models.Product) error {
	var query = `insert into public.products(modelname, version, price, description, production_date) 
				values($1, $2, $3, $4, $5)`
	statement, err := p.db.Prepare(query)
	if err != nil {	return err	}
	defer statement.Close()

	_, err = statement.Exec(product.ModelName, product.Version, product.Price, product.Description, product.ProductionDate)
	if err != nil {	return err	}

	return nil
}

func (p *productRepositoryPostgres) Update(productId string, product *models.Product) error {
	query := `UPDATE public.products SET modelname=$1, version=$2, price=$3, description=$4, production_date=$5 WHERE id = $6`

	statement, err := p.db.Prepare(query)
	if err != nil { return err }
	defer statement.Close()

	_, err = statement.Exec(product.ModelName, product.Version, product.Price, product.Description, product.ProductionDate, productId)
	if err != nil { return err }

	return nil
}

func (p *productRepositoryPostgres) UpdatePrice(productId string, product *models.Product) error {
	query := `UPDATE public.products SET price=$1 WHERE id = $2`

	statement, err := p.db.Prepare(query)
	if err != nil { return err }
	defer statement.Close()

	_, err = statement.Exec(product.Price, productId)
	if err != nil { return err }

	return nil
}

func (p *productRepositoryPostgres) Delete(productId string) error {
	query := `DELETE FROM public.products WHERE id = $1`

	statement, err := p.db.Prepare(query)
	if err != nil { return err }
	defer statement.Close()

	_, err = statement.Exec(productId)
	if err != nil { return err }

	return nil
}


func (p *productRepositoryPostgres) FindAll() (models.Products, error) {
	query := `SELECT id, modelname, version, price, description, production_date FROM products`

	var products models.Products
	rows, err := p.db.Query(query)

	if err != nil {	return nil, err	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.ModelName, &product.Version, &product.Price,
			&product.Description, &product.ProductionDate)

		if err != nil {	return nil, err	}
		products = append(products, product)
	}
	return products, nil
}
//Save(*models.Product) error
//Update(string, *models.Product) error
//Delete(string) error
//FindByID(string) (*models.Product, error)
//FindProductsByUserId(string) (*models.Product, error)
//FindAll() (models.Products, error)

func (p *productRepositoryPostgres) FindProductsByUserId(userId string) (models.Products, error) {

	query := `select p.id, p.modelname, p.version, p.price, p.description, p.production_date from users u
inner join users_products up on (u.id = up.users_id)
inner join products p on (p.id = up.products_id) where u.id = $1`

	var products models.Products

	statement, err := p.db.Prepare(query)
	if err != nil { return nil, err	}

	defer statement.Close()
	rows, err := statement.Query(userId)
	if err != nil {	return nil, err	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.ModelName, &product.Version, &product.Price, &product.Description, &product.ProductionDate)

		if err != nil {	return nil, err	}
		products = append(products, product)
	}

	fmt.Println(products)
	return products, nil
}
