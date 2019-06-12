// Package product contains product application services
package product

import (
	"log"

	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/structs"
)

// Create creates a new product account
func (u *Product) Create(c echo.Context, req model.Product) (*model.Product, error) {
	log.Println("aaaaaaaaaa")
	return u.udb.Create(u.db, req)
}

// List returns list of products
func (u *Product) List(c echo.Context, p *model.Pagination, query []string) ([]model.Product, int, error) {
	// TODO: []string to filter query
	q := &model.FilterQuery{}
	log.Println("Params", query)
	return u.udb.List(u.db, q, p)
}

// View returns single product
func (u *Product) View(c echo.Context, query []string) (*model.Product, error) {
	q := &model.FilterQuery{
		// Query:  "product.id=?",
		// Params: []interface{}{query},
	}
	return u.udb.View(u.db, q)
}

// Delete deletes a product
func (u *Product) Delete(c echo.Context, id int) error {
	q := &model.FilterQuery{Query: "product.id=?", Params: []interface{}{id}}
	product, err := u.udb.View(u.db, q)
	if err != nil {
		return err
	}
	return u.udb.Delete(u.db, product)
}

// Update updates product's contact information
func (u *Product) Update(c echo.Context, req model.UpdateReq) (*model.Product, error) {
	q := &model.FilterQuery{Query: "product.id=?", Params: []interface{}{req.ID}}
	product, err := u.udb.View(u.db, q)
	if err != nil {
		return nil, err
	}

	structs.Merge(product, req)
	if err := u.udb.Update(u.db, product); err != nil {
		return nil, err
	}

	return product, nil
}
