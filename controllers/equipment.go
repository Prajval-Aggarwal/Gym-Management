package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
)

func CreateEquipHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var equipment mod.Equipment
	json.NewDecoder(r.Body).Decode(&equipment)

	result := db.DB.Model(&mod.Equipment{}).Where("equip_name =?", equipment.Equip_Name).Updates(&equipment)
	if result.Error != nil {
		fmt.Println("error in DB")
	} else if result.RowsAffected == 0 { //if the equip_name is not in record then create new record
		db.DB.Create(&equipment)
		fmt.Fprint(w, "New Equipment added")
		json.NewEncoder(w).Encode(&equipment)

	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" Old Equipment updated successfully"))
		json.NewEncoder(w).Encode(&equipment)
	}

}

func GetEquipList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var equipments []mod.Equipment
	db.DB.Find(&equipments)

	//sort Equipments alphabetically
	query := "SELECT * FROM equipment ORDER BY equip_name ASC"
	db.DB.Raw(query).Scan(&equipments)
	json.NewEncoder(w).Encode(&equipments)
}