package product


// terminal: go run cmd/product/product.go -config ./config/config.yaml

import (
	"github.com/jmoiron/sqlx"
	"github.com/seminarioGo/internal/config"	
)

type product struct {
	ID int64
	Name string
	Price float64
}

type Service interface {
	InsertProduct(product) (*product,error)
	FindProducts() ([]*product, error)
	UpdateProduct(product) (*product,error)
	DeleteProduct(int) (bool,error)
	FindProductID(int) (*product, error)
} 

type service struct {
	db *sqlx.DB
	conf *config.Config
}

func New(db *sqlx.DB,c *config.Config) (Service, error){
	return service{db,c}, nil
}

func (s service) FindProductID(ID int) (*product, error) {
	var product product
	query := "SELECT * FROM product WHERE ID=?"
	if err := s.db.Get(&product,query, ID); 
	err != nil {
		return nil, err
	}	
	return &product, nil
}

func (s service) FindProducts() ([]*product, error) {
	var list []*product
	query := "SELECT * FROM product"
	if err := s.db.Select(&list,query); 
	err != nil {		
		return nil, err
	}
	return list, nil
}

func (s service) InsertProduct(p product) (*product,error) {
	query := "INSERT INTO product (name, price) VALUES (?,?)"
	res, err := s.db.Exec(query, p.Name, p.Price)
	if err != nil {
		return nil,err
	}
	id,_ := res.LastInsertId()	
	var product product
	query = "SELECT * FROM product WHERE ID=?"
	if err := s.db.Get(&product,query, id); 
	err != nil {
		return nil, err
	}	
	return &product, nil	
}

func (s service) UpdateProduct(p product) (*product,error) {
	query := "UPDATE product SET name = ?, price = ? WHERE ID = ?"	
	_, err := s.db.Exec(query, p.Name, p.Price, p.ID)
	if err != nil {
		return nil,err
	}
	var product product
	query = "SELECT * FROM product WHERE ID=?"
	if err := s.db.Get(&product,query, p.ID); 
	err != nil {
		return nil, err
	}	
	return &product, nil	
}

func (s service) DeleteProduct(ID int) (bool,error) {
	query := "DELETE FROM product WHERE ID = ?"
	_,err := s.db.Exec(query, ID) 
	if err != nil {
		return false, err
	}
	return true, nil
}