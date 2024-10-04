package queryBuilder

import (
	"database/sql"
	"fmt"
	"strings"
)

type UpdateBuilder struct {
	queryParts []string
	values     []interface{}
}

func NewUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{}
}

func (ub *UpdateBuilder) Update(table string) *UpdateBuilder {
	ub.queryParts = append(ub.queryParts, fmt.Sprintf("UPDATE %s", table))
	return ub
}

func (ub *UpdateBuilder) Set(columns ...string) *UpdateBuilder {
	setPairs := []string{}
	for _, column := range columns {
		setPairs = append(setPairs, fmt.Sprintf("%s = ?", column))
	}
	ub.queryParts = append(ub.queryParts, fmt.Sprintf("SET %s", strings.Join(setPairs, ", ")))
	return ub
}

func (ub *UpdateBuilder) Values(values ...interface{}) *UpdateBuilder {
	ub.values = append(ub.values, values...)
	return ub
}

func (ub *UpdateBuilder) Where(condition string) *UpdateBuilder {
	ub.queryParts = append(ub.queryParts, fmt.Sprintf("WHERE %s", condition))
	return ub
}

func (ub *UpdateBuilder) Build(db *sql.DB) (sql.Result, error) {
	query := strings.Join(ub.queryParts, " ")
	fmt.Println(query)
	return db.Exec(query, ub.values...)

}
