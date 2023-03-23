package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	cons "gym-api/Utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(1 * time.Minute)
	fmt.Println("expiration time is: ", expirationTime)
	var cred mod.Credential
	username := r.URL.Query().Get("username")
	err := db.DB.Where("user_name=?", username).First(&cred).Error
	if err != nil {
		w.Write([]byte("User with given username do not exists....."))
		return
	}
	//check if the user is valid then only create the token
	claims := mod.Claims{
		Username: cred.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(cons.SecretKey)
	if err != nil {
		fmt.Println("error is :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprint("Token is:", tokenString)))
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.URL.Query().Get("token")

	claims := &mod.Claims{}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return cons.SecretKey, nil
	})
	if err != nil || !parsedToken.Valid {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	fmt.Println(claims.Username)
	var password = make(map[string]string)

	var userCred mod.Credential
	err = db.DB.Where("user_name=?", claims.Username).Find(&userCred).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	json.NewDecoder(r.Body).Decode(&password)
	fmt.Println("Password: ", password["password"])
	bs, err := bcrypt.GenerateFromPassword([]byte(password["password"]), 8)
	if err != nil {
		panic(err)
	}
	userCred.Password = string(bs)
	userCred.Password = password["password"]
	err = db.DB.Where("user_name=?", claims.Username).Updates(userCred).Error
	if err != nil {
		http.Error(w, "Failed to update user password", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Password updated successfully"))

}
