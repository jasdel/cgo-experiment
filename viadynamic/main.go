package main

import "flag"

func main() {
	title := flag.String("t", "title", "The title of message to display")
	msg := flag.String("m", "hello", "The message to display")
	flag.Parse()

	defer cleanup()

	showDialog(*title, *msg)

	doCallback()

}
