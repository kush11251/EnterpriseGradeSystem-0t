package controllers

import (
	"github.com/EnterpriseGradeSystem/pkg/models"
	"github.com/EnterpriseGradeSystem/pkg/services"
	"github.com/EnterpriseGradeSystem/pkg/utils"
	"net/http"
)

// RunServer starts the server
//
// @Description: RunServer is a function that starts the server.
//
func RunServer() {
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	userService := services.NewUserService()
	users, err := userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSONResponse(w, users)
}