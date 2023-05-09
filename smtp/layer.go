package smtp

type SMTPLayer interface {
	// Check Connection
	CheckSMTPConnection() error
}
