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
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

func CreateReceipt(rw http.ResponseWriter, r *http.Request) {
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

	documentName := r.FormValue("document_name")
	documentType := r.FormValue("document_type")
	documentProperty := r.FormValue("document_property")
	documentInformation := r.FormValue("document_information")
	idUserSender, _ := strconv.Atoi(r.FormValue("id_user_sender"))
	idUserReceiver, _ := strconv.Atoi(r.FormValue("id_user_receiver"))
	date := r.FormValue("date")
	// date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	status := r.FormValue("status")

	receipt := models.Receipt{
		DocumentName:        documentName,
		DocumentType:        documentType,
		DocumentProperty:    documentProperty,
		DocumentInformation: documentInformation,
		IdUserSender:        idUserSender,
		IdUserReceiver:      idUserReceiver,
		Date:                date,
		Status:              status,
	}

	err := config.DB.Save(&receipt).Error

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	files := r.MultipartForm.File["images"]

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

		imageReceipt := models.ImageReceipt{
			IdReceipt: receipt.IdReceipt,
			Image:     f.Name(),
		}

		err = config.DB.Save(&imageReceipt).Error

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    receipt,
	})
}

func UpdateStatusReceipt(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	receiptId := r.URL.Query()["receiptId"]

	receipt := models.Receipt{}

	err := config.DB.Model(&receipt).Where("id_receipt = ?", receiptId).Update("status", "1").Error

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	intReceiptId, err := strconv.Atoi(receiptId[0])
	err = config.DB.Model(models.Receipt{IdReceipt: intReceiptId}).First(&receipt).Error
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
		"data":    receipt,
	})
}

func GetListReceiptMyTask(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	idUser := r.URL.Query()["idUser"]

	// receipts := []map[string]interface{}{}

	receipts := []models.Receipt{}

	config.DB.Model(models.Receipt{}).Where("id_user_receiver = ? AND status = 0", idUser).Find(&receipts)

	receiptResponse := []models.ReceiptResponse{}
	for _, v := range receipts {
		// get user for sender
		var userSender = models.User{}
		config.DB.Where("id_user = ?", v.IdUserSender).First(&userSender)

		// get user for receiver
		var userReceiver = models.User{}
		config.DB.Where("id_user = ?", v.IdUserReceiver).First(&userReceiver)

		receiptResponse = append(receiptResponse, models.ReceiptResponse{
			IdReceipt:           v.IdReceipt,
			DocumentName:        v.DocumentName,
			DocumentType:        v.DocumentType,
			DocumentProperty:    v.DocumentProperty,
			DocumentInformation: v.DocumentInformation,
			UserSender:          userSender,
			UserReceiver:        userReceiver,
			Date:                v.Date,
			Status:              v.Status,
		})
	}

	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    receiptResponse,
	})
}

func GetListReceiptWaiting(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	idUser := r.URL.Query()["idUser"]

	// receipts := []map[string]interface{}{}

	receipts := []models.Receipt{}

	config.DB.Model(models.Receipt{}).Where("id_user_sender = ? AND status = 0", idUser).Find(&receipts)

	receiptResponse := []models.ReceiptResponse{}
	for _, v := range receipts {
		// get user for sender
		var userSender = models.User{}
		config.DB.Where("id_user = ?", v.IdUserSender).First(&userSender)

		// get user for receiver
		var userReceiver = models.User{}
		config.DB.Where("id_user = ?", v.IdUserReceiver).First(&userReceiver)

		receiptResponse = append(receiptResponse, models.ReceiptResponse{
			IdReceipt:           v.IdReceipt,
			DocumentName:        v.DocumentName,
			DocumentType:        v.DocumentType,
			DocumentProperty:    v.DocumentProperty,
			DocumentInformation: v.DocumentInformation,
			UserSender:          userSender,
			UserReceiver:        userReceiver,
			Date:                v.Date,
			Status:              v.Status,
		})
	}

	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    receiptResponse,
	})
}

func GetHistory(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	idUser := r.URL.Query()["idUser"]

	// receipts := []map[string]interface{}{}

	receipts := []models.Receipt{}

	config.DB.Model(models.Receipt{}).Where("id_user_sender = ? OR id_user_receiver = ? AND status = 1", idUser, idUser).Find(&receipts)

	receiptResponse := []models.ReceiptResponse{}
	for _, v := range receipts {
		// get user for sender
		var userSender = models.User{}
		config.DB.Where("id_user = ?", v.IdUserSender).First(&userSender)

		// get user for receiver
		var userReceiver = models.User{}
		config.DB.Where("id_user = ?", v.IdUserReceiver).First(&userReceiver)

		receiptResponse = append(receiptResponse, models.ReceiptResponse{
			IdReceipt:           v.IdReceipt,
			DocumentName:        v.DocumentName,
			DocumentType:        v.DocumentType,
			DocumentProperty:    v.DocumentProperty,
			DocumentInformation: v.DocumentInformation,
			UserSender:          userSender,
			UserReceiver:        userReceiver,
			Date:                v.Date,
			Status:              v.Status,
		})
	}

	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    receiptResponse,
	})
}

