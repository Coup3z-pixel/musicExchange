package templates

import (
	"fmt"
	"net/http"
	"music-exchange/templates"
	"music-exchange/templates/dashboard"
	"music-exchange/templates/leaderboard"
	"music-exchange/templates/stats"
	"music-exchange/util"

	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
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

	// SQL Query

	util.Render(ctx, 200, dashboard.Dashboard())
}

func Index(ctx *gin.Context) {
	util.Render(ctx, 200, templates.Index())
}

func Leaderboard(ctx *gin.Context) { util.Render(ctx, 200, leaderboard.Leaderboard()) }
func Stats(ctx *gin.Context) { util.Render(ctx, 200, stats.Stats()) }
