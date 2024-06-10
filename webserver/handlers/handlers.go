package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gitgub.com/emersonary/gilasw/go/global"
	"gitgub.com/emersonary/gilasw/go/infra/database"
	"gitgub.com/emersonary/gilasw/go/internal/dto"
	"gitgub.com/emersonary/gilasw/go/internal/model"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func NewDBHandler(db *gorm.DB) *DBHandler {
	return &DBHandler{
		DB: db,
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(global.UsersComplete)

}

func GetOptions(w http.ResponseWriter, r *http.Request) {

	log.Println("GetOptions executed")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)

}

func GetCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(global.Categories)

}

func GetChannels(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(global.Channels)

}

type Error struct {
	Message string `json:"message"`
}

func (d *DBHandler) GetLastMessages(w http.ResponseWriter, r *http.Request) {

	messageDB := database.NewMessage(d.DB)

	messages, err := messageDB.GetMessagesLastRows(10)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
	w.WriteHeader(http.StatusCreated)

}

func (d *DBHandler) GetLastMessagesNotifications(w http.ResponseWriter, r *http.Request) {

	messageDB := database.NewMessage(d.DB)

	messages, err := messageDB.GetMessagesNotificationsLastRows(10)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
	w.WriteHeader(http.StatusCreated)

}

func (d *DBHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {

	var messagedto dto.CreateMessage
	err := json.NewDecoder(r.Body).Decode(&messagedto)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	categoryDB := database.NewCategory(d.DB)
	category, err := categoryDB.FindByID(messagedto.CategoryID)
	if err != nil {
		error := Error{Message: "Inexistent Category ID"}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messagetoinsert, err := model.NewMessage(messagedto.CategoryID, messagedto.MessageText)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messageDB := database.NewMessage(d.DB)

	messagereturn, err := messageDB.Create(messagetoinsert)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	messagereturn.Category = *category
	json.NewEncoder(w).Encode(messagereturn)
	w.WriteHeader(http.StatusCreated)

}
