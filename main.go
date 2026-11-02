package main

func main() {
	state := config{state : "begun",
					Next : 0,
					Previous : 0}
	startRepl(&state)
}
