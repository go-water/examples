package main

import (
	"github.com/go-water/water"
)

func main() {
	r := water.New()

	r.GET("/", func(c *water.Context) {
		c.Text(200, "Hello, World!")
	})

	r.Run(":8080")
}
