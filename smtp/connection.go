package smtp

import (
	"github.com/dev-hana/go-mailer/conf"
	"gopkg.in/gomail.v2"
)

type SMTP struct {
	*gomail.Dialer
}

func ConnectSMTP() (*SMTP, error) {
	smtp, err := conf.GetSMTPConfig()
	if err != nil {
		return nil, err
	}

	dialer := gomail.NewDialer(smtp.Host, smtp.Port, smtp.User, smtp.Password)
	return &SMTP{
		Dialer: dialer,
	}, nil
}
