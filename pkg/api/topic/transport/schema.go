package transport

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	dynamicGraphql "github.com/eneoti/dynamic-struct-graphql"
	merp "github.com/eneoti/merge-struct"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/api/topic"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// List get list by graphql
func List(c echo.Context, svc topic.Service, page *model.PaginationReq) graphql.Schema {
	objectType := dynamicGraphql.DynamicGraphql("ListResponse", reflect.TypeOf(listResponse{}))
	resolver := func(p graphql.ResolveParams) (interface{}, error) {
		var queryArr []string
		var tsQueryArr []string

		//-
		key := p.Args["key"]
		//-- key not empty
		if key != nil && strings.TrimSpace(key.(string)) != "" {
			tmpStr := key.(string)
			tsQueryArr = append(tsQueryArr, `' '`)
			tsQueryArr = append(tsQueryArr, "topic.name")
			tsQueryArr = append(tsQueryArr, "topic.title")
			keyStr := `%s:*`
			//--
			f := func(c rune) bool {
				return c == ' '
			}
			temps := strings.FieldsFunc(tmpStr, f)
			if len(temps) > 0 {
				keyStr = strings.Join(temps, ` & `)
			}
			tempStr := fmt.Sprintf(`to_tsvector('english', f_concat_ws(%s))
				@@ to_tsquery('english', '%s')`,
				strings.Join(tsQueryArr, ", "), keyStr)

			queryArr = append(queryArr, tempStr)

		}

		name := p.Args["name"]
		if name != nil {
			tempStr := fmt.Sprintf(`topic.name = '%s'`, name.(string))
			queryArr = append(queryArr, tempStr)
		}
		// if queryArr == nil || len(queryArr) == 0 {
		// 	return nil, helper.HandleError("web:topic:list:invalidParams")
		// }
		result, totalItems, err := svc.List(c, page.Transform(), queryArr)
		resultOutput := listResponse{result, page.Page, totalItems}

		if err != nil {
			return nil, err
		}
		var mapValue interface{}
		conv := merp.ConventionalMarshaller{Value: resultOutput}

		b, err := conv.MarshalJSON()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, &mapValue)
		if err != nil {
			return nil, err
		}
		return mapValue, nil
	}
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"result": &graphql.Field{
					Type: objectType,
					Args: graphql.FieldConfigArgument{
						"key": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver,
				},
			},
		})
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	return schema
}
