package pkg1

import (
	"log"

	"github.com/sbinet-vgo/api"
)

func Do(name string) {
	var g api.Greeter
	log.Printf("expects a greeting: %q", g.Greet(name))
}
