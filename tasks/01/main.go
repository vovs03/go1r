package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// format: 28.12.1844 17:45

	timeFormat := fmt.Sprintf("%d.%d.%d %d:%d",
		t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(),
	)

	fmt.Println("Привет, вот время, в котором ты находишься: ", timeFormat)
}
