package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var cred mod.Credential
	json.NewDecoder(r.Body).Decode(&cred)

	err := db.DB.Where("user_name=?", cred.UserName).First(&mod.Credential{}).Error
	if err == nil {
		fmt.Println("User already exist please login to move forward...")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bs, err := bcrypt.GenerateFromPassword([]byte(cred.Password), 8)
	if err != nil {
		panic(err)
	}
	cred.Password = string(bs)
	db.DB.Create(&cred)
	w.Write([]byte("User Registerd sucessfully"))
	json.NewEncoder(w).Encode(cred)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var cred mod.Credential
	json.NewDecoder(r.Body).Decode(&cred)
	var existCred mod.Credential
	err := db.DB.Where("user_name=?", cred.UserName).First(&existCred).Error
	if err != nil {
		fmt.Println("User do not exists please register first...")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(existCred.Password), []byte(cred.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Incorrect Password")
		return
	}
	fmt.Println("Logged In Successfully....")

}
