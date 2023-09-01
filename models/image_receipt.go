package models

type ImageReceipt struct {
	ImageId   int    `gorm:"primarykey" json:"image_id" form:"image_id"`
	IdReceipt int    `json:"id_receipt" form:"id_receipt"`
	Image     string `json:"image" form:"image"`
}
