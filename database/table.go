package database

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID          uint64    `gorm:"primaryKey;autoIncrement column:id"`
	Name        string    `gorm:"not null column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"default:now()"`
}

func (Project) TableName() string {
	return "project_tb"
}

type Template struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey;autoIncrement column:id"`
	Title     string    `gorm:"not null column:title"`
	Header    string    `gorm:"column:header"`
	Footer    string    `gorm:"column:footer"`
	ProjectID uint64    `gorm:"column:project_id"`
	Project   Project   `gorm:"foreignkey:ProjectID"`
	CreatedAt time.Time `gorm:"default:now()"`
}

func (Template) TableName() string {
	return "template_tb"
}

type SendMail struct {
	gorm.Model
	ID            uint64    `gorm:"primaryKey;autoIncrement column:id"`
	Status        bool      `gorm:"not null default:false column:status"`
	ReceiverName  string    `gorm:"column:receiver_name"`
	ReceiverEmail string    `gorm:"not null column:receiver_email"`
	Content       string    `gorm:"not null column:content"`
	TemplateID    uint64    `gorm:"column:tempalte_id"`
	Template      Template  `gorm:"foreignkey:TemplateID"`
	CreatedAt     time.Time `gorm:"default:now()"`
}

func (SendMail) TableName() string {
	return "send_mail_tb"
}

type SMTP struct {
	gorm.Model
	Host          string    `gorm:"not null column:host"`
	Port          int       `gorm:"not null column:port"`
	User          string    `gorm:"not null column:user"`
	Password      string    `gorm:"not null column:password"`
	Vertification bool      `gorm:"not null default:false column:vertification"`
	CreatedAt     time.Time `gorm:"default:now()" column:"created_at"`
}

func (SMTP) TableName() string {
	return "smtp_tb"
}

func (db *GORM) InitTable() error {
	return db.AutoMigrate(&SendMail{}, &Project{}, &Template{}, &SMTP{})
}
