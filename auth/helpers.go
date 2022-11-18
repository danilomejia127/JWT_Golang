package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Here, we are declaring a secret key that will be used later for generating JWTs. For now, the key is “supersecretkey”.
// You should be ideally storing this value outside the code. But for the sake of simplicity, let’s proceed as it is.
var jwtKey = []byte("supersecretkey")

// JWTClaim - We define a custom struct for JWT Claims which will ultimately become the payload of the JWT
type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT - function, which takes in email and username as parameters, would return the generated JWT string.
// Here we set a default expiration time as 1 Hour, which can be (and should be) made configurable. From there on we
// create a new claim variable with the available data and expiration time. Finally, we generate the token using the
// HS256 Signing Algorithm by passing the previously created claims.
func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

// ValidateToken -- we would take in the token string coming from the client’s HTTP request header and validate it.
// we will be using this function in our Authentication middleware to verify if the incoming client request is
// authenticated. Clear, yeah? So, here we will try to parse the JWT into claims using the JWT package’s helper method
// “ParseWithClaims”. From the parsed token, we extract the claims. Using these claims, we check if the token is
// actually expired or not.
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
