package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thg021/go-hexagonal/application"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(ID string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(ID).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
