package cont

import (
	"encoding/json"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
)

func AddEquipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var equipment mod.Equipment
	json.NewDecoder(r.Body).Decode(&equipment)
	db.DB.Create(&equipment)
	json.NewEncoder(w).Encode(&equipment)
}

func GetEquipList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var equipments []mod.Equipment
	db.DB.Find(&equipments)
	json.NewEncoder(w).Encode(&equipments)
}
