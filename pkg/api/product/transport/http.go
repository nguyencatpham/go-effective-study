package transport

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/api/product"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/helper"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// HTTP represents product http service
type HTTP struct {
	svc product.Service
}

// NewHTTP creates new product http service
func NewHTTP(svc product.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/products")

	// swagger:operation GET /v1/products products listProducts
	// ---
	// summary: Returns list of products.
	// description: Returns list of products. Depending on the product role requesting it, it may return all products for SuperAdmin/Admin products, all company/location products for Company/Location admins, and an error for non-admin products.
	// parameters:
	// - name: query
	//   in: query
	//   description:  graphql query.To get list :{list(key:""){page,totalItems, data{name,description,type}}}. To get 1 item {result(key:""){name,description,type}}
	//   type: string
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/productListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.query)

	// swagger:operation POST /v1/products products productCreate
	// ---
	// summary: Returns product created.
	// description: Returns list of products. Depending on the product role requesting it, it may return all products for SuperAdmin/Admin products, all company/location products for Company/Location admins, and an error for non-admin products.
	// parameters:
	// - name: query
	//   in: query
	//   description:  graphql query.To get list :{create(name:"",description:"",type:""){name,description,type}}.
	//   type: string
	//   required: false
	// - name: mutation
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/productCreate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/productListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.POST("", h.create)

	// swagger:operation PATCH /v1/products/{id} products productUpdate
	// ---
	// summary: Updates product's contact information
	// description: Updates product's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/productUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/productResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("", h.update)

	// swagger:operation DELETE /v1/products/{id} products productDelete
	// ---
	// summary: Deletes a product
	// description: Deletes a product with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of product
	//   type: number
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.DELETE("/:id", h.delete)
	// swagger:operation GET /v1/products/test products productDelete
	// ---
	// summary: Deletes a product
	// description: Deletes a product with requested ID.
	// parameters:
	// - name: query
	//   in: query
	//   description:  graphql query.To get list :{list(key:""){page,totalItems, data{name,description,type}}}. To get 1 item {result(key:""){name,description,type}}
	//   type: string
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/test", h.test)
}

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

func (h *HTTP) test(c echo.Context) error {
	log.Println("params", c.QueryParams())
	// r := new(model.UpdateReq)
	// if err := c.Bind(r); err != nil {
	// 	return err
	// }
	r := &model.UpdateReq{Name: "1", Description: "2", Type: 3}
	schema := Mutation(c, h.svc, r)
	requestString := c.QueryParam("query")
	result, err := mutation(schema, requestString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *HTTP) query(c echo.Context) error {
	p := new(model.PaginationReq)
	if err := c.Bind(p); err != nil {
		return err
	}

	schema := Query(c, h.svc, p)
	log.Println(schema)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: c.QueryParam("query"),
	})
	if len(result.Errors) > 0 {
		errors := ""
		for _, b := range result.Errors {
			errors += b.Message
		}
		return helper.HandleError(errors)
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) create(c echo.Context) error {
	r := new(model.UpdateReq)
	if err := c.Bind(r); err != nil {
		return err
	}
	schema := Mutation(c, h.svc, r)
	log.Println(schema)
	requestString := fmt.Sprintf("mutation _{create(name:\"%s\",description:\"%s\",type:%d){name,description,type}}", r.Name, r.Description, r.Type)
	result, err := mutation(schema, requestString)
	log.Println(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *HTTP) update(c echo.Context) error {
	r := new(model.UpdateReq)
	if err := c.Bind(r); err != nil {
		return err
	}
	schema := Mutation(c, h.svc, r)
	requestString := fmt.Sprintf("{update(id:\"%d\",name:\"%s\",description:\"%s\",type:\"%d\"){id,name,description,type}}", r.ID, r.Name, r.Description, r.Type)
	result, err := mutation(schema, requestString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
func (h *HTTP) delete(c echo.Context) error {
	r := new(model.UpdateReq)
	if err := c.Bind(r); err != nil {
		return err
	}
	schema := Mutation(c, h.svc, r)
	requestString := c.QueryParam("id")
	result, err := mutation(schema, requestString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
func mutation(schema graphql.Schema, requestString string) (*graphql.Result, error) {

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestString,
	})
	if len(result.Errors) > 0 {
		errors := ""
		for _, b := range result.Errors {
			errors += b.Message
		}
		return nil, helper.HandleError(errors)
	}
	return result, nil
}
