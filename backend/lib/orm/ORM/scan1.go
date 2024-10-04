package orm

import (
	"fmt"
	"reflect"
	"strings"
)

func (o *ORM) Scan1(table interface{}, columns ...string) (interface{}, error) {
	var (
		__BUILDER__ = NewSQLBuilder()
	)
	_, __table := InitTable(table)
	__BUILDER__.custom = o.Custom
	query, param := __BUILDER__.Select(columns...).From(__table).Build()
	rows, err := o.Db.Query(query, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var fields []reflect.StructField
	for _, col := range columns {
		fieldName := col
		if dotIndex := strings.LastIndex(col, "."); dotIndex != -1 {
			fieldName = col[dotIndex+1:]
		}
		field, exists := UniqueField[fieldName]

		if !exists {
			return nil, fmt.Errorf("field %s does not exist in table", fieldName)
		}
		fields = append(fields, reflect.StructField{Name: fieldName, Type: field.Type})
	}
	resultType := reflect.SliceOf(reflect.StructOf(fields))
	results := reflect.MakeSlice(resultType, 0, 0)
	for rows.Next() {
		values := make([]interface{}, len(fields))
		for i := range fields {
			values[i] = reflect.New(fields[i].Type).Interface()
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}
		newStruct := reflect.New(reflect.StructOf(fields)).Elem()
		for i, value := range values {
			newStruct.Field(i).Set(reflect.ValueOf(value).Elem())
		}
		results = reflect.Append(results, newStruct)
	}
	o.Custom.parameters = nil
	o.Custom.query = ""
	return results.Interface(), nil
}
