package orm

import (
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type TestTable struct {
	ID   int    `orm:"PRIMARY KEY"`
	Name string `orm:""`
}

func TestInitDB(t *testing.T) {
	testORM := &ORM{}

	dbName := "test.db"
	dirMigration := "./migrations"

	defer os.Remove(dbName)
	defer os.RemoveAll(dirMigration)

	err := testORM.InitDB(dbName, dirMigration)
	if err != nil {
		t.Errorf("InitDB() failed: %s", err)
	}

	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		t.Errorf("Database file %s was not created.", dbName)
	}

	if _, err := os.Stat(dirMigration); os.IsNotExist(err) {
		t.Errorf("Migration directory %s was not created.", dirMigration)
	}
}

func TestCreateTable(t *testing.T) {
	tableName := "test_table"
	fields := []*Field{
		NewField("ID", reflect.TypeOf(0), "PRIMARY KEY"),
		NewField("Name", reflect.TypeOf("test"), ""),
	}
	expectedSQL := "CREATE TABLE IF NOT EXISTS test_table (\n\tID INTEGER PRIMARY KEY,\n\tName TEXT \n)"

	sql := CreateTable(tableName, fields...)

	if sql != expectedSQL {
		t.Errorf("CreateTable() = %v and want %v", sql, expectedSQL)
	}
}

func TestAutoMigrate(t *testing.T) {
	dirMigration := "./test_migrations"
	defer os.RemoveAll(dirMigration)

	orm := &ORM{}
	dbName := "test.db"
	orm.InitDB(dbName, dirMigration)

	defer os.Remove(dbName)

	err := orm.AutoMigrate(dirMigration, TestTable{})
	if err != nil {
		t.Errorf("AutoMigrate() failed: %s", err)
	}

}
