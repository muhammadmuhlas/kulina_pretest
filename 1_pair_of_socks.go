package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

// PairOfSocks calculate how many pair of socks could sell
func PairOfSocks() {
	reader := bufio.NewReader(os.Stdin)
	sockColorMap, totalSock := make(map[int]int), 0

	getString := func(reader *bufio.Reader) string {
		input, err := reader.ReadString('\n')
		HandleError(err)
		input = strings.Replace(input, "\n", "", -1)
		return input
	}
	for i := 0; i < 2; i++ {

		// Get total of socks available
		if i == 0 {
			fmt.Print("total sock: ")
			ns, err := strconv.Atoi(getString(reader))
			HandleError(err)
			totalSock = ns
		}

		// Get color of sock in pile
		if i == 1 {
			fmt.Print("sock color separated by space: ")
			sSocks := strings.Split(strings.TrimSpace(getString(reader)), " ")
			if totalSock != len(sSocks) {
				panic("invalid input for sockColorMap and totalSock")
			}

			for _, sSock := range sSocks {
				if sSock == "" {
					continue
				}
				sock, err := strconv.Atoi(sSock)
				HandleError(err)

				// Add sock to map to count socks by color later
				if _, ok := sockColorMap[sock]; ok {
					sockColorMap[sock] += 1
				} else {
					sockColorMap[sock] = 1
				}
			}
		}
	}

	// count sellable socks by pairing each color of socks available
	var sellable int
	for _, v := range sockColorMap {
		sellable += v / 2
	}

	fmt.Println(sellable)
}
