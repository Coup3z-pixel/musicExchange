package db

import (
	"context"
	"music-exchange/models"

	"github.com/jackc/pgx/v5"
)

type UserDB struct {
	Postgres *Postgres
}

func (userDB *UserDB) DoesUserExist(email string, service string) bool {

	var user_country string
	row := userDB.Postgres.db.QueryRow(
		context.Background(), 
		"SELECT country FROM users WHERE email=$1 AND service=$2", 
		email, service,
	)

	if row.Scan(&user_country) != pgx.ErrNoRows {
		return true
	}

	return false
}

func (userDB *UserDB) InsertUser(user models.UserProfile) bool {
	_, err := userDB.Postgres.db.Exec(
		context.Background(), 
		"INSERT INTO users (email, username, country, service) VALUES ($1, $2, $3, $4)",
		user.Email, user.Username, user.Country, user.Service,
	)

	if err != nil {
		return false
	}

	 return true
}
