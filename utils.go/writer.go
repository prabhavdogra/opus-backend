package utils

import (
	"fmt"
	"os"

	"opus-backend/repl"
	"opus-backend/store"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
)

func ParseInput(ActiveNotes map[int64][]store.Note) {
	noteGroups := store.ParseNotes(ActiveNotes)
	store.ParseChords(noteGroups)
}

func InputChannel(in midi.In, ActiveNotes map[int64][]store.Note) {
	input := ""
	for {
		fmt.Scan(&input)
		if input == "exit" {
			err := in.StopListening()
			fmt.Println(err)
			fmt.Printf("closing MIDI Port %v.....\n", in)
			os.Exit(0)
		} else if input == "compile" {
			err := in.StopListening()
			Must(err)
			codestring := string(ReadFile("test.opus"))
			repl.Start(codestring)
			fmt.Printf("closing MIDI Port %v.....\n", in)
			os.Exit(0)
		}
	}
}

func Writer() {
	drv, err := rtmididrv.New()
	Must(err)

	defer drv.Close()

	ins, err := drv.Ins()
	Must(err)

	// takes the first MIDI input channel
	in := ins[0]

	fmt.Printf("opening MIDI Port %v\n", in)
	Must(in.Open())

	defer in.Close()

	ActiveNotes := make(map[int64][]store.Note)

	MIDIReader(in, ActiveNotes)

	InputChannel(in, ActiveNotes)
}
