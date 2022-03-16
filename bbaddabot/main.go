package main

import (
	// "bbaddabot/presentation"
	"bbaddabot/persistence"
)

func main() {
	persistence.InsertHistory()
	persistence.SelectTodayHistory("hj")
	// presentation.PresentationTest()
	// presentation.Bbaddabot()
}
