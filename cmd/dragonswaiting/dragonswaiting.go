/*
 * Import users from csv file to waiting list in Scoutnet
 *
 * For Sjöscoutkåren Drakarna, feel free to reuse.
 */

/*
MIT License

Copyright (c) 2019 Johan Manner

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mrManner/membernet-go/pkg/shared"
	"github.com/mrManner/membernet-go/pkg/waitinglist"
)

func main() {
	pathPtr := flag.String("infile", "", "path to csv file with members to add")
	groupPtr := flag.String("group", "", "group id of group to import to")
	keyPtr := flag.String("apikey", "", "api key for waitinglist API in selected group")
	hostPtr := flag.String("host", "scoutnet.se", "membernet host to use")

	flag.Parse()

	file, err := os.Open(*pathPtr)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(file))

	for {
		// read file line by line, post to Scoutnet
		// print all errors on stdout
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		ssno, _ := strconv.Atoi(line[4])
		var profile = waitinglist.Profile{
			Dob:       line[1] + "-" + line[2] + "-" + line[3],
			Ssno:      ssno,
			Note:      line[0],
			FirstName: line[6],
			LastName:  line[7],
			Gender:    line[5],
			Email:     line[8],
			Address: waitinglist.Address{
				AddressType: 1,
				Street:      line[9],
				ZipCode:     line[10],
				Country:     shared.Sweden,
				ZipName:     line[11],
			},
			Relative1: waitinglist.Contact{
				Name:   line[15],
				Phone:  line[18],
				Mobile: line[17],
				Email:  line[16],
			},
			Relative2: waitinglist.Contact{
				Name:   line[19],
				Phone:  line[22],
				Mobile: line[21],
				Email:  line[20],
			},
		}

		var leader bool
		if line[22] == "1" {
			leader = true
		} else {
			leader = false
		}

		waitinglist.Register(profile, leader, *groupPtr,
			*keyPtr,
			*hostPtr,
		)

	}
}
