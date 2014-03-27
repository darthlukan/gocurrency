/*
	GoCurrency v1.0

	Copyright (C) 2014, Brian Tomlinson <brian.tomlinson@linux.com>

	This program is free software; you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation; either version 2 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License along
    with this program; if not, write to the Free Software Foundation, Inc.,
    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	version        string = "GoCurrency version 1.0 \nWritten by Brian Tomlinson <brian.tomlinson@linux.com>\n"
	usage          string = "\ngocurrency <amount> <from currency> <to currency>\n\nExample:\n\tgocurrency 20 usd gbp\n\n"
	displayVersion *bool  = flag.Bool("version", false, "Output version information and exit.")
	displayUsage   *bool  = flag.Bool("usage", false, "Display usage and exit")
	apiKey         string = "0a9f4758fe7a4f5c8fddfb2ff6b926530b3f2ebe"
	baseUrl        string = "http://currency-api.appspot.com/api/"
	requestUrl     string
	data           map[string]interface{}
)

func main() {
	flag.Parse()

	if *displayVersion {
		os.Stdout.WriteString(version)
	}

	if *displayUsage {
		os.Stdout.WriteString(usage)
	}

	if flag.NArg() == 3 {
		amount := flag.Args()[0]
		from := flag.Args()[1]
		to := flag.Args()[2]

		endUrl := fmt.Sprintf("%s/%s.json?key=%s&amount=%s", from, to, apiKey, amount)
		requestUrl := fmt.Sprintf("%s%s", baseUrl, endUrl)

		response, err := http.Get(requestUrl)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		fmt.Printf("%v %v equals %v %v\n", amount, from, data["amount"], to)
	} else {
		err := errors.New("You supplied invalid input, please run gocurrency with the --usage flag for an example of proper input.")
		fmt.Println(err)
	}

}
