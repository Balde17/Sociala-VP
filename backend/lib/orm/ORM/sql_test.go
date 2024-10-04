package orm

import (
	"reflect"
	"testing"
)

func TestSQLBuilder_Select(t *testing.T) {
	builder := NewSQLBuilder().Select("id", "name").From(&Table{Name: "users"})
	query, params := builder.Build()

	expectedQuery := "SELECT id, name FROM users"
	if query != expectedQuery || len(params) != 0 {
		t.Errorf("Select failed. Expected %q, got %q", expectedQuery, query)
	}
}

func TestSQLBuilder_Where(t *testing.T) {
	builder := NewSQLBuilder().Select("name").From(&Table{Name: "users"}).Where("id", 1)
	query, params := builder.Build()

	expectedQuery := "SELECT name FROM users WHERE id = ?"
	if query != expectedQuery || !reflect.DeepEqual(params, []interface{}{1}) {
		t.Errorf("Where failed. Expected %q with params %v, got %q with params %v", expectedQuery, []interface{}{1}, query, params)
	}
}

// Vous pouvez ajouter d'autres tests pour Update, Delete, Join, etc.

func TestSQLBuilder_Update(t *testing.T) {
	updates := &Modifier{Model: &Table{Name: "users"}, field: "email", value: "new@example.com"}
	builder := NewSQLBuilder().Update(updates)
	query, params := builder.Build()

	expectedQuery := "UPDATE users SET email = ?"
	if query != expectedQuery || !reflect.DeepEqual(params, []interface{}{"new@example.com"}) {
		t.Errorf("Update failed. Expected %q with params %v, got %q with params %v", expectedQuery, []interface{}{"new@example.com"}, query, params)
	}
}

func TestSQLBuilder_Delete(t *testing.T) {
	builder := NewSQLBuilder().Delete().From(&Table{Name: "users"}).Where("id", 1)
	query, params := builder.Build()

	expectedQuery := "DELETE  FROM users WHERE id = ?"
	if query != expectedQuery || !reflect.DeepEqual(params, []interface{}{1}) {
		t.Errorf("Delete failed. Expected %q with params %v, got %q with params %v", expectedQuery, []interface{}{1}, query, params)
	}
}

func TestSQLBuilder_Join(t *testing.T) {
	builder := NewSQLBuilder().Select("users.name", "orders.amount").From(&Table{Name: "users"}).Join("orders", "users.id = orders.user_id")
	query, params := builder.Build()

	expectedQuery := "SELECT users.name, orders.amount FROM users JOIN orders ON users.id = orders.user_id"
	if query != expectedQuery || len(params) != 0 {
		t.Errorf("Join failed. Expected %q, got %q", expectedQuery, query)
	}
}

func TestSQLBuilder_OrderBy(t *testing.T) {
	builder := NewSQLBuilder().Select("name").From(&Table{Name: "users"}).OrderBy("name", "DESC")
	query, params := builder.Build()

	expectedQuery := "SELECT name FROM users ORDER BY name DESC" // Assuming `Order[1]` equals "ASC"
	if query != expectedQuery || len(params) != 0 {
		t.Errorf("OrderBy failed. Expected %q, got %q", expectedQuery, query)
	}
}

func TestSQLBuilder_Limit(t *testing.T) {
	builder := NewSQLBuilder().Select("name").From(&Table{Name: "users"}).Limit(10)
	query, params := builder.Build()

	expectedQuery := "SELECT name FROM users LIMIT ?"
	if query != expectedQuery || !reflect.DeepEqual(params, []interface{}{10}) {
		t.Errorf("Limit failed. Expected %q with params %v, got %q with params %v", expectedQuery, []interface{}{10}, query, params)
	}
}

func TestSQLBuilder_GroupBy(t *testing.T) {
	builder := NewSQLBuilder().Select("category", "COUNT(*)").From(&Table{Name: "products"}).GroupBy("category")
	query, params := builder.Build()

	expectedQuery := "SELECT category, COUNT(*) FROM products GROUP BY category"
	if query != expectedQuery || len(params) != 0 {
		t.Errorf("GroupBy failed. Expected %q, got %q", expectedQuery, query)
	}
}

func TestSQLBuilder_Having(t *testing.T) {
	builder := NewSQLBuilder().Select("category", "COUNT(*)").From(&Table{Name: "products"}).GroupBy("category").Having("COUNT(*) > 10")
	query, params := builder.Build()

	expectedQuery := "SELECT category, COUNT(*) FROM products GROUP BY category HAVING COUNT(*) > 10"
	if query != expectedQuery || len(params) != 0 {
		t.Errorf("Having failed. Expected %q, got %q", expectedQuery, query)
	}
}
