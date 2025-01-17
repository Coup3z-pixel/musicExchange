package templates

import (
	addsong "rank-and-roll/templates/add-song"
	"rank-and-roll/util"

	"github.com/gin-gonic/gin"
)

func AddSong(ctx *gin.Context) {
	util.Render(ctx, 200, addsong.AddSong())
}
