package queryBuilder

import (
	"database/sql"
	"fmt"
	"strings"
)

type DeleteBuilder struct {
	queryParts []string
	values     []interface{}
}

func NewDeleteBuilder() *DeleteBuilder {
	return &DeleteBuilder{}
}
func (dib *DeleteBuilder) DeleteValues(table string) *DeleteBuilder {
	dib.queryParts = append(dib.queryParts, fmt.Sprintf("DELETE FROM %s", table))
	return dib
}

func (dib *DeleteBuilder) Where(condition string) *DeleteBuilder {
	dib.queryParts = append(dib.queryParts, fmt.Sprintf("WHERE %s", condition))
	return dib
}

func (dib *DeleteBuilder) Build() (string, []interface{}) {
	query := strings.Join(dib.queryParts, " ")
	return query, dib.values
}
func (dib *DeleteBuilder) DeleteQuery(db *sql.DB) (bool, sql.Result, error) {
	query, values := dib.Build()
	rslt, err := db.Exec(query, values...)
	if err != nil {
		return false, nil, err
	}
	nbrRows, err := rslt.RowsAffected()
	if err != nil {
		return false, nil, err
	}
	return nbrRows > 0, rslt, nil
}
