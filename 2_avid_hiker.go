package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

// AvidHiker calculate number of valleys Gary walked through during his hike.
func AvidHiker() {
	reader := bufio.NewReader(os.Stdin)
	eachSteps, totalSteps := make([]string, 0), 0

	getString := func(reader *bufio.Reader) string {
		input, err := reader.ReadString('\n')
		HandleError(err)
		input = strings.Replace(input, "\n", "", -1)
		return input
	}
	for i := 0; i < 2; i++ {

		// Get total of steps
		if i == 0 {
			fmt.Print("total steps : ")
			ns, err := strconv.Atoi(getString(reader))
			HandleError(err)
			totalSteps = ns
		}

		// Get each steps
		if i == 1 {
			fmt.Print("each step separated by space: ")
			sSteps := strings.Split(strings.TrimSpace(getString(reader)), " ")
			if totalSteps != len(sSteps) {
				panic("invalid input for eachSteps and totalSteps")
			}

			for _, sStep := range sSteps {
				if sStep == "U" || sStep == "D" {
					eachSteps = append(eachSteps, sStep)
				}
			}
		}
	}

	// Calculate each steps, is it on valley or mountain
	currentAltitude, onValley := 0, 0
	for _, v := range eachSteps {
		switch v {
		// step up
		case "U":
			currentAltitude += 1
		// step down
		case "D":
			// count onValley if step down below sea level and start from sea level and will end count when exit sea level
			if currentAltitude-1 < 0 && currentAltitude >= 0 {
				onValley += 1
			}
			currentAltitude -= 1
		}
	}
	if currentAltitude != 0 {
		panic("invalid eachSteps: not end at sea level")
	}

	fmt.Println(eachSteps, onValley)
}
