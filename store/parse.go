package store

import (
	"sort"
)

func ParseNotes(ActiveNotes map[int64][]Note) [][]Note {
	totalNotes := 0
	for _, notes := range ActiveNotes {
		totalNotes += len(notes)
	}
	song := make([]Note, totalNotes)
	for _, notes := range ActiveNotes {
		for _, note := range notes {
			if note.Key < 49 {
				continue
			}
			song = append(song, note)
		}
	}

	sort.Slice(song, func(i, j int) bool {
		return song[i].TimeStamp < song[j].TimeStamp
	})

	NoteGroups := [][]Note{}
	LocalGroup := []Note{}
	for i, note := range song {
		if i == 0 {
			LocalGroup = append(LocalGroup, note)
		} else {
			if note.TimeStamp-song[i-1].TimeStamp < 100 {
				LocalGroup = append(LocalGroup, note)
			} else {
				NoteGroups = append(NoteGroups, LocalGroup)
				LocalGroup = []Note{}
				LocalGroup = append(LocalGroup, note)
			}
		}
	}
	NoteGroups = append(NoteGroups, LocalGroup)
	return NoteGroups
}

func ParseChords(NoteGroups [][]Note) {
	codeString := ""
	for i, group := range NoteGroups {
		token := getToken(group)
		if token != "" {
			NoteGroups[i][0].Name = token
			NoteGroups[i] = NoteGroups[i][:1]
			codeString += NoteGroups[i][0].Name
		}
	}
	overwriteFile("test.opus", codeString)
}
