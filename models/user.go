package models

type User struct {
	Id       int    `json:"id" form:"id"`
	Nik      string `json:"nik" form:"nik"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Jabatan  string `json:"jabatan" form:"jabatan"`
	Lokasi   string `json:"lokasi" form:"lokasi"`
}

type UserResponse struct {
	Id      int    `json:"id" form:"id"`
	Email   string `json:"email" form:"email"`
	Token   string `json:"token" form:"token"`
	Name    string `json:"name" form:"name"`
	Jabatan string `json:"jabatan" form:"jabatan"`
	Lokasi  string `json:"lokasi" form:"lokasi"`
}
