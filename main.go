package main

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/lukewarlow/GoConsoleMenu"
	"math"
	"time"
)

func Break(length int) {
	fmt.Printf("Taking a %d minute break\n", length)

	now := time.Now()
	lengthInMinutes := time.Duration(length) * time.Minute
	breakOverTime := now.Add(lengthInMinutes)
	fmt.Printf("The Break will End at " + breakOverTime.Format(time.Kitchen))
	fmt.Println()
	fmt.Println()
	tmpl := `{{ red "Time Left:" }} {{string . "my_green_string" | green}}:{{string . "my_blue_string" | blue}} ⇒ {{ bar . "|" "█" (cycle . "⛦" "⛧" "⛦" "⛧" ) "░" "|"}}  {{percent .}}`

	// start bar based on our template
	bar := pb.ProgressBarTemplate(tmpl).Start(length * 60)

	// set values for string elements
	var secs int = 60
	for {

		time.Sleep(1000 * time.Millisecond)
		left := time.Now().Sub(breakOverTime)
		minutes := int32(math.Abs(math.Trunc(left.Minutes())))

		bar.Set("my_green_string", minutes).Set("my_blue_string", secs)
		bar.Increment()
		secs--
		if secs == 0 {
			secs = 60
		}
		if time.Now().After(breakOverTime) {
			bar.Finish()
			break
		}
	}
	fmt.Println("\n\nBreak is Over!")
}
func main() {

	subMenu := GoConsoleMenu.NewMenu("⛣ Super Deluxe Break Timer ⛳")
	subMenu.AddMenuItem(GoConsoleMenu.NewActionItem(0, "15 Minute Break", func() { Break(15) }))
	subMenu.AddMenuItem(GoConsoleMenu.NewActionItem(1, "10 Minute Break", func() { Break(10) }))
	subMenu.AddMenuItem(GoConsoleMenu.NewActionItem(2, "5 Minute Break", func() { Break(5) }))
	subMenu.AddMenuItem(GoConsoleMenu.NewActionItem(3, "1 Minute Break", func() { Break(1) }))

	subMenu.Display()
}
