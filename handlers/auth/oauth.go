package auth

import "rank-and-roll/db"

type OAuthHandlers struct {
	DB *db.UserDB
	SongDB *db.SongDB
}
