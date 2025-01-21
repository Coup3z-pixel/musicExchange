package db

import (
	"context"
	"fmt"
	"music-exchange/models"

	"github.com/jackc/pgx/v5"
)

type SongDB struct {
	Postgres *Postgres
}

func (songDB *SongDB) AddUserTracks(tracks []map[string]interface{}) bool {
	// track format: [{track artist}, {track artist}, ...]

	sql_cmd := "INSERT INTO songs (name, artist, elo) VALUES "
	insert_vals := []interface{}{}

	for i := 0; i < len(tracks); i++ {

		row_values := "(?, ?, ?),"
		sql_cmd += row_values

		name, track_entry := tracks[i]["name"].(string)
		i++
		fmt.Println(tracks[i])
		artist_name, artist_entry := tracks[i]["artist"].(string)

		track_exists := songDB.getSong(name, artist_name)

		

		if track_entry && artist_entry && track_exists == nil {
			insert_vals = append(insert_vals, name, artist_name, 100)
		} else {
			return false
		}
	}

	fmt.Println("Loop Exits")

	sql_cmd = sql_cmd[0:len(sql_cmd)-1]
	_, err := pgInstance.db.Exec(context.Background(), sql_cmd)

	if err != nil {
		return false
	}
	 
	return true
}

func (SongDB *SongDB) AddSong(track models.Track) {
}

func (songDB *SongDB) getSong(name string, artist string) *models.Track {
	var song models.Track

	row := songDB.Postgres.db.QueryRow(
		context.Background(),
		"SELECT * FROM songs WHERE name=$1 AND artist=$2",
		name, artist,
	)

	if row.Scan(&song) == pgx.ErrNoRows {
		return nil
	}

	return &song	
}
