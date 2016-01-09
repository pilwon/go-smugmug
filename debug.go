package smugmug

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/fatih/color"
)

var debug = true

func doDebug(data []byte, err error) {
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	fmt.Printf("%s\n\n", data)
}

func debugRequest(req *http.Request) {
	if debug {
		color.Set(color.FgMagenta)
		doDebug(httputil.DumpRequestOut(req, true))
		color.Unset()
	}
}

func debugResponse(res *http.Response) {
	if debug {
		color.Set(color.FgCyan)
		doDebug(httputil.DumpResponse(res, true))
		color.Unset()
	}
}
