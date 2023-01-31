package main

import (
	"chapter7/my"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// main program.
func main() {
	my.Migrate()
}
