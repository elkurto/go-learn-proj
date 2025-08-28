package music_catalog

type Band struct {
	Id   int
	Name string
}

type Musician struct {
	Id   int
	Name string
}

type BandMusician struct {
	Id       int
	BandId   int
	Musician int
}

type Song struct {
	Id                  int
	Name                string
	MapIdToBandMusician map[int]BandMusician
	ListBandMusician    []BandMusician
}
