package conf

import (
	"fmt"
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
		debug, port, err := GetServerConfig()
		if err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println(debug, port)
	})
}
