package orm

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestORM_Insert_Scan(t *testing.T) {
	o := NewORM()
	dbName, dbDir := "testdb", "test"
	defer os.Remove(dbName)
	defer os.RemoveAll(dbDir)

	o.InitDB(dbName, dbDir)
	defer o.Db.Close()

	err := o.AutoMigrate(dbDir, TestTable{})
	if err != nil {
		t.Fatalf("Failed to automigrte : %s", err)
	}

	test := TestTable{Name: "Test Name"}

	err = o.Insert(test)
	if err != nil {
		t.Fatalf("Failed to insert test data: %s", err)
	}

	result, err := o.Scan(TestTable{}, "ID", "Name")
	if err != nil {
		t.Fatalf("Scan failed: %s", err)
	}

	results := result.([]struct {
		ID   int
		Name string
	})

	if len(results) != 1 || results[0].Name != "Test Name" {
		t.Errorf("Scan did not return the expected results: got %+v", results)
	}

}
