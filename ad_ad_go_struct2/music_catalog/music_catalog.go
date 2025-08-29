package music_catalog

import "fmt"

// Band represents a collection of musicians
type Band struct {
	Id   int
	Name string
}

// create a new Music Band
func NewBand(id int, name string) *Band {
	return &Band{id, name}
}

// Musicians create songs and collect to form bands
type Musician struct {
	Id   int
	Name string
}

// create a new Musician
func NewMusician(id int, name string) *Musician {
	return &Musician{id, name}
}

// an association
type BandMusician struct {
	Id       int
	Band     Band
	Musician Musician
}

// create a new BandMusician object -
// that describes an association between Music,Band that created the Song
func NewBandMusician(id int, band Band, musician Musician) *BandMusician {
	return &BandMusician{id, band, musician}
}

// Song is created by one or MusicianBand(s)
type Song struct {
	Id                  int
	Name                string
	MapIdToBandMusician map[int]BandMusician
	ListBandMusician    []BandMusician
}

// create a new Song object
func NewSong(id int, name string, listBandMusician []BandMusician) *Song {
	mapIdToBandMusician := make(map[int]BandMusician)

	for _, bandMusician := range listBandMusician {
		mapIdToBandMusician[bandMusician.Id] = bandMusician
	}
	song := Song{}
	song.Id = id
	song.Name = name
	song.ListBandMusician = listBandMusician
	song.MapIdToBandMusician = mapIdToBandMusician

	return &song
}

// convert Song to string
func (s Song) String() string {
	return fmt.Sprintf("Song(%d,%s)", s.Id, s.Name)
}

// factory to create []*Song
func CreateSampleData02() []*Song {
	m1 := NewMusician(11, "Solveig")
	m2 := NewMusician(12, "Dragonette")

	b3 := NewBand(13, "DJ Solveig")
	b4 := NewBand(14, "Dragonette")

	mb31 := NewBandMusician(15, *b3, *m1)
	mb42 := NewBandMusician(16, *b4, *m2)

	var listPtrSong = []*Song{
		NewSong(1, "Big in Japan", []BandMusician{*mb31}),
		NewSong(2, "Hello", []BandMusician{*mb42}),
	}

	return listPtrSong
}
