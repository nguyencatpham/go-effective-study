package helper

import (
	"encoding/json"
	"fmt"
	"strings"

	merp "github.com/eneoti/merge-struct"
	"github.com/graphql-go/graphql"
)

func GetParams(key interface{}) []string {
	var queryArr []string
	var tsQueryArr []string

	if key == nil {
		return tsQueryArr
	}
	text := key.(string)
	tsQueryArr = append(tsQueryArr, `' '`)
	tsQueryArr = append(tsQueryArr, "topic.name")
	tsQueryArr = append(tsQueryArr, "topic.title")
	keyStr := `%s:*`
	//--
	f := func(c rune) bool {
		return c == ' '
	}
	temps := strings.FieldsFunc(text, f)
	if len(temps) > 0 {
		keyStr = strings.Join(temps, ` & `)
	}
	tempStr := fmt.Sprintf(`to_tsvector('english', f_concat_ws(%s))
				@@ to_tsquery('english', '%s')`,
		strings.Join(tsQueryArr, ", "), keyStr)

	queryArr = append(queryArr, tempStr)
	return queryArr
}
func MapResponse(conv merp.ConventionalMarshaller) (interface{}, error) {
	var mapValue interface{}

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
func NewSchema(queryType *graphql.Object, mutationType *graphql.Object) graphql.Schema {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
	return schema
}
