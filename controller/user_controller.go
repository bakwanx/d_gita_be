package controller

import (
	"d_gita_be/config"
	"d_gita_be/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Register(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	nik := r.FormValue("nik")
	password := r.FormValue("password")
	name := r.FormValue("name")
	jabatan := r.FormValue("jabatan")
	lokasi := r.FormValue("lokasi")
	file, _, err := r.FormFile("profile")

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

	// Create a temporary file within our public directory that follows
	// a particular naming pattern
	// randomImageName := strconv.Itoa(int(rand.NewSource(time.Now().UnixMicro()).Int63()))

	tempFile, err := ioutil.TempFile("public", "*.png")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	user := models.User{
		Nik:      nik,
		Password: password,
		Name:     name,
		Jabatan:  jabatan,
		Lokasi:   lokasi,
		Profile:  tempFile.Name(),
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
