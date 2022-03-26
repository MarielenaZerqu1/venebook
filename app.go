package main

import(
	"os"
	"fmt"
	"strconv"

	"github.com/gominima/minima"
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
	App.Listen(fmt.Sprintf(":%d", port))
}