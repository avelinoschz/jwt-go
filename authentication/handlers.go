package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/avelinoschz/jwt-go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// Login returns a JWT with the same user info received and standard claims
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error reading user: %s\n", err)
		fmt.Fprintf(w, "Error reading user: %s\n", err)
	}

	if user.Name == "avelino" && user.Password == "password" {
		user.Password = "" // clean password to re-use model
		user.Role = "admin"

		token := GenerateJWT(user)
		respToken := models.ResponseToken{
			Token: token,
		}

		jsonResp, err := json.Marshal(respToken)
		if err != nil {
			log.Println("Error marshaling json response token")
			fmt.Fprintln(w, "Error marshaling json response token: ", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintln(w, "Invalid user or password")
	return

}

// ValidateToken valites if the token is a valid one and haven't expired
func ValidateToken(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequestWithClaims(
		r,                       // takes the request token with all the claims included
		request.OAuth2Extractor, // type of extraction. OAuth2Extractor looks in 'Authorization' header
		&models.Claim{},         // struct or model of the claims to be extracted
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		},
	)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Fprintln(w, "Token expired")
				return
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Fprintln(w, "Signature doesn't match")
				return
			default:
				fmt.Fprintln(w, "Invalid token")
				return
			}
		default:
			fmt.Fprintln(w, "Unknown token error")
			return
		}

	}

	if token.Valid {
		log.Println("is token valid?: ", token.Valid)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Token is no valid")
	}

	return
}
