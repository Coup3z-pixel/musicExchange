package middleware

import (
	"fmt"
	"net/http"
	"music-exchange/util"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(ctx *gin.Context) {
	token_str, err := ctx.Cookie("music-exchange-token")

	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/sign")
		ctx.Abort()
		return
	}

	token, err := util.VerifyToken(token_str)

	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/sign")
		ctx.Abort()
		return
	}

	// Print information about the verified token
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)

	ctx.Next()
}
