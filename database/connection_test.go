package database

import "testing"

func TestSuccessConnectDatabase(t *testing.T) {
	dsn := "host=localhost user=postgres password=1234 dbname=mailer port=5432 sslmode=disable TimeZone=Asia/Seoul"
	_, err := ConnectDB("postgres", dsn)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestFailedConnectDatabase(t *testing.T) {
	dsn := "host=localhost user=postgres password=1234 dbname=mailer port=5432 sslmode=disable TimeZone=Asia/Seoul"
	_, err := ConnectDB("sqlite", dsn)
	if err != nil {
		t.Error(err.Error())
	}
}
