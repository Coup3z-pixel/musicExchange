package models

type Track struct {
	Name string
	Artist Artist
	Elo int
}

type Artist struct {
	Name string
}

func CreateTrack(name string, artist string, elo int) *Track {
	song_elo := elo

	if elo == -1 {
		song_elo = 1000
	}
	
	return &Track {
		Name: name,
		Artist: Artist{ Name: artist },
		Elo: song_elo,
	}
}


