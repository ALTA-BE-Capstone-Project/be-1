package entities

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Visit struct {
	ID               uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Visit_uid        string         `gorm:"index;type:varchar(22)"`
	Doctor_uid       string         `gorm:"index;type:varchar(22)"`
	Clinic_uid       string         `gorm:"index;type:varchar(22)"`
	Patient_uid      string         `gorm:"index;type:varchar(22)"`
	Date             datatypes.Date
	Status           string `gorm:"type:enum('pending', 'ready', 'cancelled');default:'pending'"`
	Complaint         string
	MainDiagnose     string
	AdditionDiagnose string
	Action           string
	Recipe           string
	BloodPressure    string
	HeartRate        string
	O2Saturate       string
	Weight           string
	Height           string
	Bmi              string
}