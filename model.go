package main

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sqlx.DB) error {
	return errors.New("Not Implemented")
}

func (p *product) updateProduct(db *sqlx.DB) error {
	return errors.New("Not Implemented")
}

func (p *product) deleteProduct(db *sqlx.DB) error {
	return errors.New("Not Implemented")
}

func (p *product) createProduct(db *sqlx.DB) error {
	return errors.New("Not Implemented")
}

func getProducts(db *sqlx.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not Implemented")
}
