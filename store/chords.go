package store

import (
	"fmt"
	"os"
	"sort"
)

func sortByKeys(notes []Note) {
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Key < notes[j].Key
	})
}

func getChord(notes []Note) string {
	sortByKeys(notes)
	chordName := ""
	diff1 := notes[1].Key - notes[0].Key
	diff2 := notes[2].Key - notes[1].Key
	if diff1 == 4 && diff2 == 3 {
		chordName += notes[0].Name + "maj"
	} else if diff1 == 3 && diff2 == 4 {
		chordName += notes[0].Name + "min"
	} else if diff1 == 3 && diff2 == 3 {
		chordName += notes[0].Name + "dim"
	} else if diff1 == 7 && diff2 == 5 {
		chordName += notes[0].Name + "pow"
	} else {
		chordName = ""
	}
	return chordName
}

func getGap(notes []Note) string {
	sortByKeys(notes)
	chordName := ""
	chordName += notes[0].Name + notes[1].Name
	return chordName
}

func getOctave(token string) (string, int) {
	octave := 0
	filteredToken := ""
	for _, char := range token {
		if char >= '0' && char <= '9' {
			if octave != 0 {
				continue
			}
			octave = int(char - '0')
		} else {
			filteredToken += string(char)
		}
	}
	return filteredToken, octave
}

func getToken(notes []Note) string {
	sortByKeys(notes)
	chord := ""
	octave := 0

	if len(notes) == 3 {
		chord, octave = getOctave(getChord(notes))
	} else if len(notes) == 2 {
		chord, octave = getOctave(getGap(notes))
	}

	if octave > 2 {
		if len(chord) == 2 && chord[0] == chord[1] {
			return fmt.Sprint(int(chord[0] - 'A'))
		}
		if len(chord) == 2 && chord[0] == 'B' {
			return string(chord[1]) + "mem"
		}
		tokenValue := ""
		tokenValue = tokenMap[chord]
		return tokenValue
	}
	return ""
}

func overwriteFile(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
