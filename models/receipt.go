package models

type Receipt struct {
	IdReceipt           int    `gorm:"primarykey" json:"id_receipt" form:"id_receipt"`
	DocumentName        string `json:"document_name" form:"document_name"`
	DocumentType        string `json:"document_type" form:"document_type"`
	DocumentProperty    string `json:"document_property" form:"document_property"`
	DocumentInformation string `json:"document_information" form:"document_information"`
	Image1              string `json:"image_1" form:"image_1"`
	Image2              string `json:"image_2" form:"image_2"`
	Image3              string `json:"image_3" form:"image_3"`
	Image4              string `json:"image_4" form:"image_4"`
	IdUserSender        int    `json:"id_user_sender" form:"id_user_sender"`
	IdUserReceiver      int    `json:"id_user_receiver" form:"id_user_receiver"`
	Date                string `json:"date" form:"date"`
	Status              string `json:"status" form:"status"`
}
