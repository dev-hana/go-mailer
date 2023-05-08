package smtp

import (
	"testing"
)

func TestSMTPConnection(t *testing.T) {
	t.Run("Connect to SMTP", func(t *testing.T) {
		_, err := ConnectSMTP()
		if err != nil {
			t.Fatal(err.Error())
		}
	})
}
