package main

func main() {
	trackers := Trackers{}
	library := NewLibrary[Trackers]("trackers.json")
	library.Load(&trackers)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&trackers)
	library.Save(trackers)
}
