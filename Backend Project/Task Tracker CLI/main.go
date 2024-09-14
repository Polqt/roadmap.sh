package main


func main() {
	trackers := Trackers{}
	trackers.add("Task 1")
	trackers.add("Task 2")

	trackers.toggle(0)
	trackers.print()
}