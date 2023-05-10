package database

func (db *GORM) CreateMail(mail SendMail) error {
	mail.Status = false
	return db.Select("receiver_name", "receiver_email", "content", "template_title", "status").Create(&mail).Error
}

func (db *GORM) GetSendMail() (mails []*SendMail, err error) {
	return mails, db.Where("status=?", false).Find(&mails).Error
}

func (db *GORM) UpdateStatus(mail *SendMail) error {
	return db.Table("send_mail_tb").Select("status").Where(&SendMail{ID: mail.ID}).Updates(mail).Error
}
