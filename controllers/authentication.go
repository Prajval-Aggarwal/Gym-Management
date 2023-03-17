package cont

import "golang.org/x/crypto/bcrypt"

func hashPass(password string) string {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bs)
}
