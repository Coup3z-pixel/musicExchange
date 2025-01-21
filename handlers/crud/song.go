package crud

import (
	"net/http"
	"music-exchange/util"
)

type AddSongForm struct {
	ID string `form:"song-id"`
	Service string `form:"streaming-service"`
}

func fetchSongById(song_details AddSongForm, bearer_token string) {
	_, err := util.CreateHttpRequest(
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
