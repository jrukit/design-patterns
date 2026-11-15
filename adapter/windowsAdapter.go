package main

type WindowsAdapter struct {
	widowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	w.widowMachine.insertIntoUSBPort()
}
