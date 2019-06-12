package transport

import (
	"log"
	"reflect"

	"github.com/nguyencatpham/go-effective-study/pkg/utl/helper"

	dynamicGraphql "github.com/eneoti/dynamic-struct-graphql"
	merp "github.com/eneoti/merge-struct"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/api/product"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// object type
var listType = dynamicGraphql.DynamicGraphql("ListResponse", reflect.TypeOf(listResponse{}))
var itemType = dynamicGraphql.DynamicGraphql("ItemResponse", reflect.TypeOf(model.Product{}))
var updateType = dynamicGraphql.DynamicGraphql("UpdateResponse", reflect.TypeOf(updateReq{}))

// Query query data by graphql
func Query(c echo.Context, svc product.Service, page *model.PaginationReq) graphql.Schema {
	var viewResolver = func(p graphql.ResolveParams) (interface{}, error) {
		key := p.Args["key"]
		result, err := svc.View(c, helper.GetParams(key))
		if err != nil {
			return nil, err
		}
		conv := merp.ConventionalMarshaller{Value: &result}

		return helper.MapResponse(conv)
	}
	var listResolver = func(p graphql.ResolveParams) (interface{}, error) {
		key := p.Args["key"]
		result, totalItems, err := svc.List(c, page.Transform(), helper.GetParams(key))
		if err != nil {
			return nil, err
		}
		resultOutput := listResponse{result, page.Page, totalItems}
		conv := merp.ConventionalMarshaller{Value: resultOutput}

		return helper.MapResponse(conv)
	}
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"result": &graphql.Field{
					Type: itemType,
					Args: graphql.FieldConfigArgument{
						"key": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: viewResolver,
				},
				"list": &graphql.Field{
					Type: listType,
					Args: graphql.FieldConfigArgument{
						"key": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: listResolver,
				},
			},
		})
	log.Println("queryType", queryType)
	log.Println("helper.NewSchema(queryType, mutationType)", helper.NewSchema(queryType, nil))
	return helper.NewSchema(queryType, nil)
}

// Mutation data
func Mutation(c echo.Context, svc product.Service, req *model.UpdateReq) graphql.Schema {
	log.Println("............Mutation.......", req)

	var createResolver = func(p graphql.ResolveParams) (interface{}, error) {
		item := model.Product{
			Name:        req.Name,
			Description: req.Description,
		}
		result, err := svc.Create(c, item)
		log.Println("...................", result)
		if err != nil {
			return nil, err
		}
		conv := merp.ConventionalMarshaller{Value: &result}

		return helper.MapResponse(conv)
	}
	var updateResolver = func(p graphql.ResolveParams) (interface{}, error) {
		item := model.UpdateReq{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Type:        req.Type,
		}
		result, err := svc.Update(c, item)
		if err != nil {
			return nil, err
		}
		conv := merp.ConventionalMarshaller{Value: &result}

		return helper.MapResponse(conv)
	}
	var deleteResolver = func(p graphql.ResolveParams) (interface{}, error) {
		id := req.ID
		if id == 0 {
			return nil, helper.HandleError("mising param id")
		}
		err := svc.Delete(c, id)
		if err != nil {
			return nil, err
		}
		conv := merp.ConventionalMarshaller{Value: nil}

		return helper.MapResponse(conv)
	}
	var mutationType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"create": &graphql.Field{
					Type: itemType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"type": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: createResolver,
				},
				"update": &graphql.Field{
					Type: itemType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"type": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: updateResolver,
				},
				"delete": &graphql.Field{
					Type: itemType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: deleteResolver,
				},
			},
		})

	return helper.NewSchema(mutationType, mutationType)
}
