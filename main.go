// http://en.wikipedia.org/wiki/Conway's_Game_of_Life

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	here := filepath.Dir(file)
	static := filepath.Join(here, "/client")
	http.Handle("/", http.FileServer(http.Dir(static)))

	grid := new(Grid)
	printer := NewHtmlStringer(grid)
	grid.Seed(gliderGun)

	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, printer.String())
		grid.Scan()
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

const gliderGun = `
------------------------x--------------------
----------------------x-x--------------------
------------xx------xx------------xx---------
-----------x---x----xx------------xx---------
xx--------x-----x---xx-----------------------
xx--------x---x-xx----x-x--------------------
----------x-----x-------x--------------------
-----------x---x-----------------------------
------------xx-------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
---------------------------------------------
`
