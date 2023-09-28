package services

import (
	"api/auth/models"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user models.UserAuth) (string, error) {
	//el hs256 es porque lo vamos verificar nosostros mismos
	tokenjwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user": user.ID,
	})

	return tokenjwt.SignedString([]byte("elefante"))
}

func ExtracClaimsJWT(token_jwt string) int {

	//estructura del objwto que vamos a retornar
	token, err := jwt.ParseWithClaims(token_jwt, &models.ClaimsJWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("elefante"), nil
	})

	if claims, ok := token.Claims.(*models.ClaimsJWT); ok && token.Valid {
		fmt.Println("aca el id")
		fmt.Printf("%v %v", claims.Id, claims.RegisteredClaims.Issuer)
		return claims.Id
	} else {
		fmt.Println(err)
		return 0
	}
}
