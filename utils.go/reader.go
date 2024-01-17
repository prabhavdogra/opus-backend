package utils

import (
	"opus-backend/store"

	"gitlab.com/gomidi/midi"
	. "gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/reader"
)

func MIDIReader(in midi.In, ActiveNotes map[int64][]store.Note) {
	rd := reader.New(
		reader.NoLogger(),
		reader.Each(func(pos *reader.Position, msg midi.Message) {
			switch v := msg.(type) {
			case NoteOn:
				currentTime := GetCurrentTimestamp()
				newNote := store.Note{Key: int(v.Key()), TimeStamp: currentTime, Name: store.MidiNoteMap[int(v.Key())], Velocity: int(v.Velocity())}
				ActiveNotes[currentTime] = append(ActiveNotes[currentTime], newNote)
				ParseInput(ActiveNotes)
			case NoteOff:
				// note := ActiveNotes[int(v.Key())]
				// delete(ActiveNotes, int(v.Key()))
			}
		}),
	)
	err := rd.ListenTo(in)
	Must(err)
}
