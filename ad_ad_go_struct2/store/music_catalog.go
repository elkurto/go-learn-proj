package music_catalog

type Band struct {
	Id   int
	Name string
}

func NewBand(id int, name string) *Band {
	return &Band{id, name}
}

type Musician struct {
	Id   int
	Name string
}

func NewMusician(id int, name string) *Musician {
	return &Musician{id, name}
}

type BandMusician struct {
	Id       int
	Band     Band
	Musician Musician
}

func NewBandMusician(id int, band Band, musician Musician) *BandMusician {
	return &BandMusician{id, band, musician}
}

type Song struct {
	Id                  int
	Name                string
	MapIdToBandMusician map[int]BandMusician
	ListBandMusician    []BandMusician
}

func NewSong(id int, name string, listBandMusician []BandMusician) *Song {
	mapIdToBandMusician := make(map[int]BandMusician)

	for _, bandMusician := range listBandMusician {
		mapIdToBandMusician[bandMusician.id] = bandMusician
	}
	song := Song{}
	song.Id = id
	song.Name = name
	song.ListBandMusician = listBandMusician
	song.MapIdToBandMusician = mapIdToBandMusician

	return &song
}
