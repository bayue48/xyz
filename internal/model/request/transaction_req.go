package request

type TransactionReq struct {
	ContractNumber string `json:"contract_number" gorm:"column:contract_number;uniqueIndex;not null" validate:"required"`
	CustomerID     int64  `json:"customer_id" gorm:"column:customer_id;not null;index"`
	OTR            int64  `json:"otr" gorm:"column:otr;not null" validate:"required,gt=0"`
	AdminFee       int64  `json:"admin_fee" gorm:"column:admin_fee;not null" validate:"required,gt=0"`
	Installment    int    `json:"installment_amount" gorm:"column:installment_amount;not null" validate:"required,oneof=1 2 3 6"`
	Interest       int64  `json:"interest_amount" gorm:"column:interest_amount;not null" validate:"required,gt=0"`
	AssetName      string `json:"asset_name" gorm:"column:asset_name;not null" validate:"required"`
	TenorMonth     int    `json:"tenor_month" gorm:"column:tenor_month;not null" validate:"required,oneof=1 2 3 6"`
}

func (t *TransactionReq) TableName() string {
	return "transactions"
}
