package auth

import "music-exchange/db"

type OAuthHandlers struct {
	DB *db.UserDB
	SongDB *db.SongDB
}
