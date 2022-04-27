package claim

import (
	"errors"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tapiaw38/irrigation-api/models/user"
)

// Claim is the custom claim
type Claim struct {
	ID    uint   `json:"id,omitempty"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token
func GenerateJWT(user user.User) (string, error) {

	myKey := []byte(os.Getenv("JWT_SECRET"))

	payload := jwt.MapClaims{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":   user.Username,
		"picture":    user.Picture,
		"is_active":  user.IsActive,
		"exp":        user.CreatedAt.Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

// ValidateJWT validates a JWT token
func ValidateJWT(tk string) (*Claim, error) {

	claim := &Claim{}
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claim, errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(tk, claim, verifyToken)

	if err != nil {
		return claim, err
	}

	if !token.Valid {
		return claim, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*Claim)

	if !ok {
		return claim, errors.New("token claims invalid")
	}

	return claim, nil
}

// verifyToken is the token verification function
func verifyToken(token *jwt.Token) (interface{}, error) {
	myKey := []byte(os.Getenv("JWT_SECRET"))

	return myKey, nil
}
