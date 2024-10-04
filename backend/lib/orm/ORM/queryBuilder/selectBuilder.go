package queryBuilder

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
)

type SelectBuilder struct {
	queryParts []string
	columns    string
	sqlQuery   *sql.Rows
	args       []interface{}
}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{}
}
func extractAlias(query string) string {
	re := regexp.MustCompile(`(?:\bAS\s+)(\w+)\b`)
	match := re.FindStringSubmatch(query)

	if len(match) > 1 {
		return match[1]
	}

	return query
}
func (sb *SelectBuilder) Select(columns ...string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "SELECT "+strings.Join(columns, ", "))
	var columnWithoutAliases []string
	for _, col := range columns {
		columnWithoutAliases = append(columnWithoutAliases, extractAlias(col))
	}
	sb.columns = strings.Join(columnWithoutAliases, ",")

	return sb
}

func (sb *SelectBuilder) SelectDistinct(columns ...string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "SELECT DISTINCT "+strings.Join(columns, ", "))
	sb.columns = strings.Join(columns, ",")
	return sb
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "FROM "+table)
	return sb
}

func (sb *SelectBuilder) Where(condition string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "WHERE "+condition)
	return sb
}
func (sb *SelectBuilder) Union() *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "UNION ")
	return sb
}
func (sb *SelectBuilder) EXCEPT() *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "EXCEPT ")
	return sb
}
func (sb *SelectBuilder) Wheres(condition string, args ...interface{}) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "WHERE "+condition)
	sb.args = append(sb.args, args...)
	return sb
}
func (sb *SelectBuilder) OrderBy(table string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "ORDER BY "+table)
	return sb
}

func (sb *SelectBuilder) Join(joinType, table, onCondition string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, fmt.Sprintf("%s JOIN %s ON %s", joinType, table, onCondition))
	return sb
}

func (sb *SelectBuilder) JoinL(table, onCondition string, joinType ...string) *SelectBuilder {
	joinString := "JOIN"
	if len(joinType) > 0 && strings.ToUpper(joinType[0]) == "LEFT" {
		joinString = "LEFT JOIN"
	}

	joinClause := fmt.Sprintf("%s %s ON %s", joinString, table, onCondition)
	sb.queryParts = append(sb.queryParts, joinClause)
	return sb
}

func (sb *SelectBuilder) GroupBy(table string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "GROUP BY "+table)
	return sb
}
func (sb *SelectBuilder) Having(cols string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "HAVING "+cols)
	return sb
}
func (sb *SelectBuilder) Between(value1, op, value2 string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "BETWEEN "+value1+" "+op+" "+value2)
	return sb
}

func (sb *SelectBuilder) Limit(limit string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "LIMIT "+limit)
	return sb
}

func (sb *SelectBuilder) Offset(offset string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "OFFSET "+offset)
	return sb
}

func (sb *SelectBuilder) In(subRequest *SelectBuilder) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, " IN ( ")
	sb.queryParts = append(sb.queryParts, subRequest.queryParts...)
	sb.queryParts = append(sb.queryParts, " )")
	return sb
}

func (sb *SelectBuilder) And(cond string) *SelectBuilder {
	sb.queryParts = append(sb.queryParts, "AND "+cond+" ")
	return sb
}

func (sb *SelectBuilder) Build() string {
	fmt.Println(strings.Join(sb.queryParts, " "))
	return strings.Join(sb.queryParts, " ")
}

// func (sb *SelectBuilder) SelectQuery(db *sql.DB) (rows [][]string, err error) {
// 	query := sb.Build()
// fmt.Println(query)

// 	sqlQuery, err := db.Query(query)
// 	if err == nil {
// 		sb.sqlQuery = sqlQuery
// 		rslt, err := sb.GetRslt()
// 		if err == nil {
// 			return rslt, nil
// 		}
// 	}
// 	return rows, err
// }

func (sb *SelectBuilder) SelectQuery(db *sql.DB) (rows [][]string, err error) {
	query := sb.Build()
	fmt.Println(query)
	sqlQuery, err := db.Query(query)
	if err == nil {
		sb.sqlQuery = sqlQuery
		rslt, err := sb.GetRslt()
		if err == nil {

			return rslt, nil
		}
	}
	return rows, err
}

// func (sb *SelectBuilder) GetRslt() (rows [][]string, err error) {
// 	columns := strings.Split(sb.columns, ",")
// 	var stringPointers []interface{}
// 	for range columns {
// 		var variable string
// 		stringPointers = append(stringPointers, &variable)
// 	}
// 	defer sb.sqlQuery.Close()

// 	for sb.sqlQuery.Next() {
// 		err = sb.sqlQuery.Scan(stringPointers...)
// 		if err != nil {
// 			return [][]string{}, err
// 		}
// 		row := []string{}
// 		for _, val := range stringPointers {
// 			if ptr, ok := val.(*string); ok {
// 				row = append(row, (*ptr))
// 			}
// 		}

//			rows = append(rows, row)
//		}
//		return rows, nil
//	}
func (sb *SelectBuilder) GetRslt() (rows [][]string, err error) {
	columns := strings.Split(sb.columns, ",")

	var nullStringPointers []interface{}
	for range columns {
		var variable sql.NullString
		nullStringPointers = append(nullStringPointers, &variable)
	}
	defer sb.sqlQuery.Close()

	for sb.sqlQuery.Next() {
		err = sb.sqlQuery.Scan(nullStringPointers...)
		if err != nil {
			return [][]string{}, err
		}
		row := []string{}
		for _, val := range nullStringPointers {
			if ptr, ok := val.(*sql.NullString); ok {
				if ptr.Valid {
					row = append(row, ptr.String)
				} else {
					row = append(row, "") // Replace NULL with empty string
				}
			}
		}

		rows = append(rows, row)
	}
	return rows, nil
}
