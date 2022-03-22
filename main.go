package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_DATABASE"))
	a.Run(":8010")
}
