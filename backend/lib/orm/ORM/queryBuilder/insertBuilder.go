package queryBuilder

import (
	"database/sql"
	"fmt"
	"strings"
)

type InsertBuilder struct {
	queryParts []string
	values     []interface{}
}

func (ib *InsertBuilder) Reinitialize() {
	ib.queryParts = []string{}
	ib.values = []interface{}{}
}

func NewInsertBuilder() *InsertBuilder {
	return &InsertBuilder{}
}
func (ib *InsertBuilder) InsertInto(table string, columns ...string) *InsertBuilder {
	ib.queryParts = append(ib.queryParts, fmt.Sprintf("INSERT INTO %s (%s)", table, strings.Join(columns, ", ")))
	return ib
}

func (ib *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	ib.values = append(ib.values, values...)
	valuePlaceholders := make([]string, len(ib.values))
	for i := range valuePlaceholders {
		valuePlaceholders[i] = "?"
	}
	valuesString := fmt.Sprintf("VALUES (%s)", strings.Join(valuePlaceholders, ", "))
	ib.queryParts = append(ib.queryParts, valuesString)
	return ib
}

func (ib *InsertBuilder) Build() (string, []interface{}) {
	return strings.Join(ib.queryParts, " "), ib.values
}

func (ib *InsertBuilder) InsertQuery(db *sql.DB) (bool, sql.Result, error) {
	query, values := ib.Build()

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
func (ib *InsertBuilder) InsertQueryLastID(db *sql.DB) (bool, sql.Result, int64, error) {
	query, values := ib.Build()

	rslt, err := db.Exec(query, values...)
	if err != nil {
		return false, nil, int64(0), err
	}

	lastInsertID, err := rslt.LastInsertId()
	if err != nil {
		return false, rslt, int64(0), err // Retourner le résultat avec l'erreur, et lastInsertID=0
	}

	nbrRows, err := rslt.RowsAffected()
	if err != nil {
		return false, rslt, lastInsertID, err // Retourner le résultat avec l'erreur
	}

	return nbrRows > 0, rslt, lastInsertID, nil // Tout s'est bien passé, retourner le résultat et lastInsertID
}
