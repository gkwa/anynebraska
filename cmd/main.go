package main

import (
	"os"

	"github.com/taylormonacelli/anynebraska"
)

func main() {
	code := anynebraska.Execute()
	os.Exit(code)
}
