package crud

import (
	"fmt"
	"net/http"
	"rank-and-roll/util"

	"github.com/gin-gonic/gin"
)

type AddSongForm struct {
	ID string `form:"song-id"`
	Service string `form:"streaming-service"`
}


func AddSongById(ctx *gin.Context) {
	
	fetchSongById(postForm, "")

	var postForm AddSongForm
	ctx.Bind(&postForm)

	fmt.Println(postForm)
}

func fetchSongById(song_details AddSongForm, bearer_token string) {
	track_request, err := util.CreateHttpRequest(
		http.MethodGet, 
		"https://api.spotify.com/v1/tracks/"+song_details.ID, 
		map[string]string{
			"Authorization": "Bearer: " + bearer_token,
		}, 
		map[string]string{},
	)

	if err != nil {
	}
}
