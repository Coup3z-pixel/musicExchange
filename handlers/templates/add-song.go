package templates

import (
	addsong "music-exchange/templates/add-song"
	"music-exchange/util"

	"github.com/gin-gonic/gin"
)

func AddSong(ctx *gin.Context) {
	util.Render(ctx, 200, addsong.AddSong())
}
