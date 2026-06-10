package main

import "github.com/viktorbeznosov/job4j_go_lang_base/internal/tracker"

func main() {
	ui := tracker.UI{
		In:      tracker.ConsoleInput{},
		Out:     tracker.ConsoleOutput{},
		Tracker: tracker.NewTracker(),
	}
	ui.Run()
}
