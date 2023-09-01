package controller

import (
	"d_gita_be/config"
	"d_gita_be/models"
	"d_gita_be/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Register(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	nik := r.FormValue("nik")
	password := r.FormValue("password")
	name := r.FormValue("name")
	jabatan := r.FormValue("jabatan")
	lokasi := r.FormValue("lokasi")
	file, fileHeader, err := r.FormFile("profile")

	// error when retrieving image
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./public", os.ModePerm)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	// Create a new file in the uploads directory
	f, err := os.Create(fmt.Sprintf("./public/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	user := models.User{
		Nik:      nik,
		Password: password,
		Name:     name,
		Jabatan:  jabatan,
		Lokasi:   lokasi,
		Profile:  utils.ImageUrlProvider(f.Name(), r),
	}

	err = config.DB.Save(&user).Error
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	user := models.User{}
	dbUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = config.DB.First(&dbUser).Error
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "user not found",
			"status":  http.StatusNotFound,
		})
		return
	}

	err = config.DB.Where("nik = ?", user.Nik).First(&dbUser).Error
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "wrong email or password",
			"status":  http.StatusNotFound,
		})
		return
	}

	if dbUser.Password != user.Password {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "wrong email or password",
			"status":  http.StatusNotFound,
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data": models.UserResponse{
			IdUser:  dbUser.IdUser,
			Nik:     dbUser.Nik,
			Token:   "ini token",
			Name:    dbUser.Name,
			Jabatan: dbUser.Jabatan,
			Lokasi:  dbUser.Lokasi,
			Profile: dbUser.Profile,
		},
	})
}

func GetUserDetailById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	idUser := r.URL.Query()["idUser"]

	intIdUser, _ := strconv.Atoi(idUser[0])
	user := models.User{}

	err := config.DB.Where("id_user = ? ", intIdUser).First(&user).Error
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "user not found",
			"status":  http.StatusNotFound,
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func GetUserDetailByNik(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	nik := r.URL.Query()["nik"]

	user := models.User{}
	err := config.DB.Where("nik = ? ", nik).First(&user).Error
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "user not found",
			"status":  http.StatusNotFound,
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}
