package entity

type Transaction struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement"`
	ContractNumber string `gorm:"column:contract_number;size:50;uniqueIndex;not null"`
	CustomerID     int64  `gorm:"column:customer_id;not null;index"`
	OTR            int64  `gorm:"column:otr;not null"`
	AdminFee       int64  `gorm:"column:admin_fee;not null"`
	Installment    int64  `gorm:"column:installment_amount;not null"`
	Interest       int64  `gorm:"column:interest_amount;not null"`
	AssetName      string `gorm:"column:asset_name;size:150;not null"`
	TenorMonth     int    `gorm:"column:tenor_month;not null"`
	CreatedAt      string `gorm:"column:created_at"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
