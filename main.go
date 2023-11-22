package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Printf(url.QueryEscape("    "))
}
