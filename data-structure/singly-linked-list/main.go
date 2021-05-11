package main

import "fmt"

type song struct {
	name   string
	artist string
	next   *song
}

type playlist struct {
	name       string
	start      *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) {
	s := &song{
		name:   name,
		artist: artist,
	}
	if p.start == nil {
		p.start = s
	} else {
		current := p.start
		for current.next != nil {
			current = current.next
		}
		current.next = s
	}
}

func (p *playlist) playNextSong() {
	start := p.start
	if start == nil {
		fmt.Println("no available songs")
		return
	}
	if start.next == nil {
		fmt.Println("out of playlist")
		return
	}
	p.start = start.next
	fmt.Printf("playing song: name-%s | artist-%s\n", p.start.name, p.start.artist)
	p.nowPlaying = start
}

func (p *playlist) showPlaylist() {
	start := p.start
	if start == nil {
		fmt.Println("no available songs")
		return
	}
	for start.next != nil {
		fmt.Printf("song: name-%s | artist-%s\n", start.name, start.artist)
		start = start.next
	}
}

func main() {
	plist := createPlaylist("Freak D Love.Song")
	plist.addSong("Chang the tim duoc em", "hkz")
	plist.addSong("Tinh dep nhat khi tan vo", "phongle")
	plist.addSong("no one love me", "lucifer")
	plist.playNextSong()
	plist.playNextSong()
}
