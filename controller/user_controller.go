package controller

import (
	"d_gita_be/config"
	"d_gita_be/models"
	"encoding/json"
	"net/http"
)

func Register(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = config.DB.Save(&user).Error
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}
