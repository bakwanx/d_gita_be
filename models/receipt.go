package models

type Receipt struct {
	IdReceipt           int    `gorm:"primarykey" json:"id_receipt" form:"id_receipt"`
	DocumentName        string `json:"document_name" form:"document_name"`
	DocumentType        string `json:"document_type" form:"document_type"`
	DocumentProperty    string `json:"document_property" form:"document_property"`
	DocumentInformation string `json:"document_information" form:"document_information"`
	IdUserSender        int    `json:"id_user_sender" form:"id_user_sender"`
	IdUserReceiver      int    `json:"id_user_receiver" form:"id_user_receiver"`
	Date                string `json:"date" form:"date"`
	Status              string `json:"status" form:"status"`
}

type ReceiptResponse struct {
	IdReceipt           int    `json:"id_receipt" form:"id_receipt"`
	DocumentName        string `json:"document_name" form:"document_name"`
	DocumentType        string `json:"document_type" form:"document_type"`
	DocumentProperty    string `json:"document_property" form:"document_property"`
	DocumentInformation string `json:"document_information" form:"document_information"`
	UserSender          User   `json:"user_sender" form:"user_sender"`
	UserReceiver        User   `json:"user_receiver" form:"user_receiver"`
	Date                string `json:"date" form:"date"`
	Status              string `json:"status" form:"status"`
}
