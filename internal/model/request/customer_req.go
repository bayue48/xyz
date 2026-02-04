package request

type CustomerReq struct {
	NIK         string `json:"nik" gorm:"column:nik;uniqueIndex;size:16;not null" validate:"required,len=16"`
	FullName    string `json:"full_name" gorm:"column:full_name;not null" validate:"required"`
	LegalName   string `json:"legal_name" gorm:"column:legal_name;not null" validate:"required"`
	BirthPlace  string `json:"birth_place" gorm:"column:birth_place;not null" validate:"required"`
	BirthDate   string `json:"birth_date" gorm:"column:birth_date;type:date;not null" validate:"required,datetime=2006-01-02"`
	Salary      int64  `json:"salary" gorm:"column:salary;not null" validate:"required,gt=0"`
	SelfiePhoto string `json:"selfie_photo" gorm:"column:selfie_photo;not null" validate:"required"`
	KTPPhoto    string `json:"ktp_photo" gorm:"column:ktp_photo;not null" validate:"required"`
}

type LoginReq struct {
	NIK string `json:"nik" validate:"required,len=16"`
}

type VerifyReq struct {
	Token string `json:"token" validate:"required,uuid4"`
}

func (c *CustomerReq) TableName() string {
	return "customers"
}
