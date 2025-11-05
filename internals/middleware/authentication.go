package middleware

import (
	"base_crud_api/internals/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication(g *gin.Context) {

	auth_header := g.Request.Header.Get("Authorization")
	// fmt.Println("header a ", auth_header)
	if auth_header == "" {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "auth headers not provided"})
		g.Abort()
		return

	}
	// make sure that the
	split_strings := strings.Split(auth_header, " ")
	if len(split_strings) != 2 {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect auth headers format provided"})
		g.Abort()
		return
	}
	token := split_strings[1]

	mapClaims, err := utils.VerifyJWT(token)

	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		g.Abort()
		return
	}
	g.Set("claims", mapClaims)

	g.Next()
}
