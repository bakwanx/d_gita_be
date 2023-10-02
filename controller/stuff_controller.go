package controller

import (
	"d_gita_be/config"
	"d_gita_be/models"
	"encoding/json"
	"d_gita_be/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func PostStuff(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	stuffName := r.FormValue("stuff_name")
	stock := r.FormValue("stock")
	typeStuff := r.FormValue("type_stuff")
	file, fileHeader, err := r.FormFile("stuff")

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
	intStock, err := strconv.Atoi(stock)
	stuff := models.Stuff{
		StuffName:	        stuffName,
		Stock:		        intStock,
		Image:			    utils.ImageUrlProvider(f.Name(), r),
		Type: 				typeStuff,
	}

	err = config.DB.Save(&stuff).Error
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
		"data":    stuff,
	})

}
