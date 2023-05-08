package conf

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Run("get db config", func(t *testing.T) {
		dbms, dsn, err := GetDBConfig()
		if err != nil {
			t.Fatal(err.Error())
		}
		if dbms == "" || dsn == "" {
			t.Fatal("nil")
		}
	})

	t.Run("get server config", func(t *testing.T) {
		_, _, err := GetServerConfig()
		if err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("get smtp config", func(t *testing.T) {
		smtp, err := GetSMTPConfig()
		if err != nil {
			t.Fatal(err.Error())
		}
		if smtp == nil {
			t.Fatal("smtp info was nil")
		}
	})
}
