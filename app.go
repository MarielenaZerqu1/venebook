package main

import(
	"os"
	"fmt"
	"strconv"
	path "path/filepath"

	"github.com/gominima/minima"
	"github.com/hexacry/venebook/routes"
)

var(
	App 	*minima.Minima = minima.New()
	port 	int
)

func init() {
	semiport := os.Getenv("VENEPORT")

	if semiport != "" {
		temp, err := strconv.Atoi(semiport)
		if err != nil {
			panic(err)
		}
		port = temp
	} else {
		port = 9090
	}
}

func main() {
	App.UseGroup(routes.V1A())
	App.File("/", path.Join("public", "www", "index.html"))

	App.Listen(fmt.Sprintf(":%d", port))
}