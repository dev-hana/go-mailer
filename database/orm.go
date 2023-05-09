package database

func (db *GORM) GetSendMail() (mails []*SendMail, err error) {
	return mails, db.Preload("project_tb").Preload("template_tb").Where(SendMail{Status: false}).Find(&mails).Error
}
