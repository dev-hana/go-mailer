package database

import (
	"fmt"
	"time"
)

type Project struct {
	Name        string    `gorm:"primaryKey;autoIncrement:false column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"default:now()"`
}

func (Project) TableName() string {
	return "project_tb"
}

type Template struct {
	Title       string    `gorm:"primaryKey;autoIncrement:false column:title"`
	Header      string    `gorm:"column:header"`
	Footer      string    `gorm:"column:footer"`
	ProjectName string    `gorm:"column:project_name"`
	Project     Project   `gorm:"foreignkey:ProjectName;references:Name"`
	CreatedAt   time.Time `gorm:"default:now()"`
}

func (Template) TableName() string {
	return "template_tb"
}

type SendMail struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement column:id"`
	Status        bool      `gorm:"not null default:false column:status"`
	ReceiverName  string    `gorm:"column:receiver_name"`
	ReceiverEmail string    `gorm:"not null column:receiver_email"`
	Content       string    `gorm:"not null column:content"`
	TemplateTitle string    `gorm:"column:tempalte_title"`
	Template      Template  `gorm:"foreignkey:TemplateTitle;references:Title"`
	CreatedAt     time.Time `gorm:"default:now()"`
}

func (SendMail) TableName() string {
	return "send_mail_tb"
}

func (db *GORM) InitTable() error {
	// drop table when exists
	err := db.Migrator().DropTable("send_mail_tb", "template_tb", "project_tb")
	if err != nil {
		return err
	}

	// create table
	err = db.AutoMigrate(&SendMail{}, &Template{}, &Project{})
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
			fmt.Println("Init DB!!")
		}
	}()

	// insert example project
	example := Project{
		Name: "example",
	}
	err = tx.Select("name").Create(&example).Error
	if err != nil {
		return err
	}

	// insert example template
	template := Template{
		Title:       "회원가입 인증코드 발송",
		Header:      `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html></head><body>`,
		Footer:      `</body></html>`,
		ProjectName: example.Name,
	}
	err = tx.Select("title", "header", "footer", "project_name").Create(&template).Error
	return err
}
