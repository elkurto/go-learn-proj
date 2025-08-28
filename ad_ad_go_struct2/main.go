package main

import (
	"ad_ad_go_struct2/music_catalog"
	"fmt"
)

func createSampleData() []*music_catalog.Song {
	m1 := music_catalog.NewMusician(1, "Hiatt")
	m2 := music_catalog.NewMusician(2, "DjMarsh")

	b3 := music_catalog.NewBand(3, "Hiatt and the Combo")
	b4 := music_catalog.NewBand(4, "DJ Marsh")

	mb13 := music_catalog.NewBandMusician(5, *b3, *m1)
	mb24 := music_catalog.NewBandMusician(6, *b4, *m2)

	var listSong = []*music_catalog.Song{
		music_catalog.NewSong(1, "Master Of Disaster", []music_catalog.BandMusician{*mb13}),
		music_catalog.NewSong(2, "Live1", []music_catalog.BandMusician{*mb24}),
	}

	return listSong
}

func main() {
	fmt.Println("hello")
	listPtrSong := createSampleData()

	for i, songPtr := range listPtrSong {
		fmt.Printf("i =%d ::: songName =%s\n", i, songPtr.Name)
	}

}
