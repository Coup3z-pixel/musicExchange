package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rank-and-roll/models"
	"rank-and-roll/util"

	"github.com/gin-gonic/gin"
)



func (oAuthHandler *OAuthHandlers) CreateSpotifyOAuth(ctx *gin.Context) {
	redirect_uri := "http://localhost:8080/spotify-callback"

	spotify_oauth_url := fmt.Sprintf(
		"https://accounts.spotify.com/authorize?response_type=%s&client_id=%s&scope=%s&redirect_uri=%s&state=%s", 
			"code", // response_type
			os.Getenv("SPOTIFY_ID"), // client_id
			"user-read-private user-read-email user-top-read", // scope
			redirect_uri, // redirect_uri
			generateRandomString(16), // state
		)

	ctx.Redirect(http.StatusFound, spotify_oauth_url)
}

func (oAuthHandler *OAuthHandlers) SpotifyCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")
	if state == "" { ctx.Redirect(http.StatusPermanentRedirect, "/sign") }

	access_token_struct, err := fetchAccessToken(code)
	if err != nil { ctx.Redirect(http.StatusPermanentRedirect, "/sign") }

	profile_response, err := fetchUserProfile(access_token_struct.AccessToken)
	userExist := oAuthHandler.DB.DoesUserExist(profile_response.Email, "Spotify");

	if !userExist {
		oAuthHandler.DB.InsertUser(*profile_response)
		user_tracks, err := fetchUserTopItems(access_token_struct.AccessToken)

		if err != nil {
		} else {
			oAuthHandler.SongDB.AddUserTracks(user_tracks.Tracks)
		}
	}

	token_str, err := util.CreateToken(profile_response.Email, profile_response.Service)
	ctx.SetCookie("rank-and-roll-token", token_str, 3600, "/", "localhost", false, true)
	ctx.Redirect(http.StatusPermanentRedirect, "/dashboard")
}



type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func fetchAccessToken(code string) (*AccessToken, error) {
	access_token, err := util.CreateHttpRequest(
		http.MethodPost, 
		"https://accounts.spotify.com/api/token", 
		map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("SPOTIFY_ID") + ":" + os.Getenv("SPOTIFY_SECRET"))),
			"content-type": "application/x-www-form-urlencoded",
		},
		map[string]string{
			"code": code,
			"redirect_uri": "http://localhost:8080/spotify-callback",
			"grant_type": "authorization_code",
		})

	response, err := io.ReadAll(access_token.Body)
	access_token_data := AccessToken{}	
	json.Unmarshal(response, &access_token_data)

	return &access_token_data, err
}

func fetchUserProfile(access_token string) (*models.UserProfile, error) {
	profile_response, err := util.CreateHttpRequest(http.MethodGet, "https://api.spotify.com/v1/me", map[string]string{
		"Authorization": "Bearer " + access_token,
	}, nil)

	user_profile := models.UserProfile{}
	profile_bytes, err := io.ReadAll(profile_response.Body)
	json.Unmarshal(profile_bytes, &user_profile)
	user_profile.Service = "Spotify"

	return &user_profile, err
}

type UserTopItems struct {
	Tracks []map[string]interface{} `json:"items"`
}

func fetchUserTopItems(access_token string) (*UserTopItems, error) {
	track_http, err := util.CreateHttpRequest(http.MethodGet, "https://api.spotify.com/v1/me/top/tracks", map[string]string {
		"Authorization": "Bearer " + access_token,
	}, nil)

	if err != nil {
		return nil, err
	}

	track_response, err := io.ReadAll(track_http.Body)
	user_top_items := UserTopItems{}	
	json.Unmarshal(track_response, &user_top_items)

	return &user_top_items, nil
}

func generateRandomString(length int) string {
	b := make([]byte, length)
   _, err := rand.Read(b)
   if err != nil {
      panic(err)
   }
   return base64.StdEncoding.EncodeToString(b) 
}
