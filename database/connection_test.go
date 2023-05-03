package database

import "testing"

func TestSuccessConnectDatabase(t *testing.T) {
	// t.Run("Connect to MySQL", func(t *testing.T) {
	// 	dbms := "mysql"
	// 	dsn := "user:password@tcp(localhost:3306)/dbname"

	// 	gorm, err := ConnectDB(dbms, dsn)
	// 	if err != nil {
	// 		t.Fatalf("unexpected error: %v", err)
	// 	}
	// 	if gorm == nil {
	// 		t.Fatal("GORM is nil")
	// 	}
	// })

	t.Run("Connect to PostgreSQL", func(t *testing.T) {
		dbms := "postgres"
		dsn := "host=localhost user=postgres password=1234 dbname=mailer port=5432 sslmode=disable"

		gorm, err := ConnectDB(dbms, dsn)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if gorm == nil {
			t.Fatal("GORM is nil")
		}
	})

	t.Run("Unsupported DBMS", func(t *testing.T) {
		dbms := "sqlite"
		dsn := "host=localhost user=postgres password=1234 dbname=mailer port=5432 sslmode=disable"

		gorm, err := ConnectDB(dbms, dsn)
		if err == nil {
			t.Fatal("expected error, but got nil")
		}
		if gorm != nil {
			t.Fatal("GORM should be nil")
		}
	})
}
