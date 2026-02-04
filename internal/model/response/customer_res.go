package response

import "xyz-multifinance/internal/entity"

type CustomerRes struct {
	ID          int64  `json:"id"`
	NIK         string `json:"nik"`
	FullName    string `json:"full_name"`
	LegalName   string `json:"legal_name"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"birth_date"`
	Salary      int64  `json:"salary"`
	KTPPhoto    string `json:"ktp_photo"`
	SelfiePhoto string `json:"selfie_photo"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToCustomerResponse(c *entity.Customer) *CustomerRes {
	return &CustomerRes{
		ID:          c.ID,
		NIK:         c.NIK,
		FullName:    c.FullName,
		LegalName:   c.LegalName,
		BirthPlace:  c.BirthPlace,
		BirthDate:   c.BirthDate,
		Salary:      c.Salary,
		KTPPhoto:    c.KTPPhoto,
		SelfiePhoto: c.SelfiePhoto,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

type LoginRes struct {
	Token string `json:"token"`
}

type Auth struct {
	// Login user id
	ID int64
}
