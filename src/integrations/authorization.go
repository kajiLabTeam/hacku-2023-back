package integrations

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
	"github.com/kajiLabTeam/hacku-2023-back/lib"
)

func GetUserByID(idToken string) (*auth.Token, error) {
	client := lib.AuthorizationConnect()
	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Fatalf("Error verifying ID token: %v\n", err)
	}

	return token,nil
}
