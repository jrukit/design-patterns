package main

func main() {
	client := Client{}

	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	window := &Windows{}
	windowsAdapter := &WindowsAdapter{
		widowMachine: window,
	}
	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
