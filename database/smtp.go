package database

func (db *GORM) CreateSMTP(smtp *SMTP) error {
	return db.Create(&smtp).Error
}
