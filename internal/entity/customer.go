package entity

type Customer struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	NIK         string `gorm:"column:nik;uniqueIndex;size:16;not null"`
	FullName    string `gorm:"column:full_name;not null"`
	LegalName   string `gorm:"column:legal_name;not null"`
	BirthPlace  string `gorm:"column:birth_place;not null"`
	BirthDate   string `gorm:"column:birth_date;type:date;not null"`
	Salary      int64  `gorm:"column:salary;not null"`
	KTPPhoto    string `gorm:"column:ktp_photo;not null"`
	SelfiePhoto string `gorm:"column:selfie_photo;not null"`
	CreatedAt   string `gorm:"column:created_at"`
	UpdatedAt   string `gorm:"column:updated_at"`
	Token       string `gorm:"column:token"`
}

func (c *Customer) TableName() string {
	return "customers"
}
