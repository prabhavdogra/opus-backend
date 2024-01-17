package utils

import (
	"fmt"
	"opus-backend/store"
)

func PrintNoteGroups(NoteGroups [][]store.Note) {
	for i, group := range NoteGroups {
		notation := "["
		if i == 0 {
			continue
		}
		for j, note := range group {
			if j == len(group)-1 {
				notation += note.Name
				continue
			}
			notation += note.Name + " "
		}
		notation += "]"
		fmt.Println(notation)
	}
}
