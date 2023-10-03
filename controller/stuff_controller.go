package controller

import (
	"d_gita_be/config"
	"d_gita_be/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"d_gita_be/utils"
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
	intStock, err := strconv.Atoi(stock)
	stuff := models.Stuff{
		StuffName:	        stuffName,
		Stock:		        intStock,
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

	files := r.MultipartForm.File["stuff"]
	for _, fileHeader := range files {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(rw, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("./public", os.ModePerm)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("./public/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		imageStuff := models.ImageStuff{
			IdStuff:   stuff.IdStuff,
			Image:     f.Name(),
		}

		err = config.DB.Save(&imageStuff).Error

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"message": "Internal Server Error",
				"status":  http.StatusInternalServerError,
			})
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    stuff,
	})
}


func PostRequestStuff(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
 
	idStuff, err := strconv.Atoi(r.FormValue("id_stuff"))
	requestInformation := r.FormValue("request_information")
	idUserRequest, err := strconv.Atoi(r.FormValue("id_user_request"))
	typeRequest := r.FormValue("type_request")
	total := r.FormValue("total")
	startTime, err := utils.DateTimeFormatter(r.FormValue("start_time"))
	endTime, err := utils.DateTimeFormatter(r.FormValue("end_time"))
	date, err := utils.DateTimeFormatter(r.FormValue("date"))


	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	requestStuff := models.RequestStuff{
		IdStuff: idStuff,
		RequestInformation: requestInformation,
		IdUserRequest  : idUserRequest,
		StartTime     : startTime,
		EndTime       : endTime,
		TypeRequest   : typeRequest,
		Total         : total,
		Status        : "1",
		Date          : date,
	}


	err = config.DB.Save(&requestStuff).Error
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	
	// rw.WriteHeader(http.StatusOK)
	// json.NewEncoder(rw).Encode(map[string]interface{}{
	// 	"message": "success",
	// 	"data":    requestStuff,
	// })
}