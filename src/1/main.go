package main

import (
	"log"
	_ "1/matchers"
	"os"
	"1/search"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("president")
}
