package orm

import (
	"reflect"
)

// The `Scan` function is a method of the `ORM` struct. It takes in a `table` interface and a variadic
// parameter `columns` of type string. It returns a map with string keys and slice of interface{}
// values.
func (o *ORM) Scan(table interface{}, columns ...string) (interface{}, error) {
	_, __table := InitTable(table)
	__BUILDER__ := NewSQLBuilder()
	__BUILDER__.custom = o.Custom
	query, param := __BUILDER__.Select(columns...).From(__table).Build()

	__BUILDER__.Clear()

	if __BUILDER__.custom != nil {
		__BUILDER__.custom.Clear()
	}

	rows, err := o.Db.Query(query, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []reflect.StructField
	for _, namefield := range columns {
		f := __table.GetField(namefield)
		newField := reflect.StructField{Name: namefield, Type: f.Type}
		fields = append(fields, newField)
	}

	__results := reflect.MakeSlice(reflect.SliceOf(reflect.StructOf(fields)), 0, 0)

	for rows.Next() {
		values := make([]interface{}, 0)
		for _, name := range columns {
			for _, v := range __table.AllFields {
				if name == v.Name {
					values = append(values, reflect.New(v.Type).Interface())
				}
			}
		}
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		newStruct := reflect.New(reflect.StructOf(fields)).Elem()

		for i, value := range values {
			val := reflect.ValueOf(value)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}

			newStruct.FieldByName(columns[i]).Set(val)
		}

		__results = reflect.Append(__results, newStruct)
	}

	return __results.Interface(), nil
}
