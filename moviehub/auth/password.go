package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword บดรหัสผ่านเป็นแฮช bcrypt (มีเกลือสุ่มในตัว)
func HashPassword(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword คืน true เมื่อรหัสที่พิมพ์ตรงกับแฮชที่เก็บไว้
func CheckPassword(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
