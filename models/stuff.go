package models

type ImageStuff struct {
	ImageId int    `gorm:"primarykey" json:"image_id" form:"image_id"`
	IdStuff int    `json:"id_stuff" form:"id_stuff"`
	Image   string `json:"image" form:"image"`
}

type Stuff struct {
	IdStuff   int    `gorm:"primarykey" json:"id_stuff" form:"id_stuff"`
	StuffName string `json:"stuff_name" form:"stuff_name"`
	Stock     int    `json:"stock" form:"stock"`
	Type      string `json:"type" form:"type"`
}

type RequestStuff struct {
	IdRequestStuff     int    `gorm:"primarykey" json:"id_request_stuff" form:"id_request_stuff"`
	IdStuff            int    `json:"id_stuff" form:"id_stuff" gorm:"not null"`
	RequestInformation string `json:"request_information" form:"request_information"`
	IdUserRequest      int    `json:"id_user_request" form:"id_user_request"`
	StartTime          string `json:"start_date" form:"start_date"`
	EndTime            string `json:"end_date" form:"end_date"`
	TypeRequest        string `json:"type_request" form:"type_request"`
	Total              string `json:"total" form:"total"`
	Status             string `json:"status" form:"status"`
	Date               string `json:"date" form:"date"`
}
