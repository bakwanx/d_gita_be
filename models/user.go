package models

type User struct {
	IdUser   int    `gorm:"primarykey" json:"id_user" form:"id_user"`
	Nik      string `json:"nik" form:"nik"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Jabatan  string `json:"jabatan" form:"jabatan"`
	Lokasi   string `json:"lokasi" form:"lokasi"`
	Profile  string `json:"profile" form:"profile"`
}

type UserResponse struct {
	IdUser  int    `json:"id_user" form:"id_user"`
	Nik     string `json:"nik" form:"nik"`
	Token   string `json:"token" form:"token"`
	Name    string `json:"name" form:"name"`
	Jabatan string `json:"jabatan" form:"jabatan"`
	Lokasi  string `json:"lokasi" form:"lokasi"`
	Profile string `json:"profile" form:"profile"`
}
