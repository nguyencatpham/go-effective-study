package pgsql

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-pg/pg"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/helper"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
)

// NewProduct returns a new product database instance
func NewProduct() *Product {
	return &Product{}
}

// Product represents the client for product table
type Product struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Product name  already exists.")
)

// Create creates a new product on database
func (self *Product) Create(db orm.DB, productModel model.Product) (*model.Product, error) {
	fmt.Printf("xxxxxxxxxxxxxxxxxxxxxxx")
	var product = new(model.Product)
	err := db.Model(product).
		Where("lower(name) = ?", strings.
			ToLower(productModel.Name)).
		Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, ErrAlreadyExists

	}

	if err := db.Insert(&productModel); err != nil {
		return nil, err
	}
	return &productModel, nil
}

// View returns single product by ID
func (self *Product) View(db orm.DB, qp *model.FilterQuery) (*model.Product, error) {
	fmt.Printf("pgsql view")
	var product = new(model.Product)
	q := db.Model(product).
		Column("product.*").
		Where("deleted_at is null")
	if qp != nil && qp.Query != "" {
		q.Where(qp.Query, qp.Params...)
	}
	err := q.Order("product.id desc").Limit(1).Select()
	if err == pg.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, helper.HandleError("web:product:view:query.error")
	}
	return product, nil
}

// Update updates product's contact info
func (self *Product) Update(db orm.DB, product *model.Product) error {
	return db.Update(product)
}

// List returns list of all products retrievable for the current product, depending on role
func (self *Product) List(db orm.DB, qp *model.FilterQuery, p *model.Pagination) ([]model.Product, int, error) {
	fmt.Printf("pgsql list")
	var products []model.Product
	q := db.Model(&products).
		Column("product.*").
		Where("deleted_at is null").
		Limit(p.Limit).
		Offset(p.Offset).
		Order("product.id desc")
	if qp != nil && qp.Query != "" {
		q.Where(qp.Query, qp.Params...)
	}
	totalItems, err := q.Count()
	if err != nil {
		// return nil, 0, helper.HandleError("web:product:list:query.error")
		return nil, 0, err
	}
	if err := q.Limit(p.Limit).Offset(p.Offset).Order("product.created_at desc").Select(); err != nil {
		return nil, 0, err
	}

	return products, totalItems, nil
}

// Delete sets deleted_at for a product
func (self *Product) Delete(db orm.DB, product *model.Product) error {
	return db.Delete(product)
}
