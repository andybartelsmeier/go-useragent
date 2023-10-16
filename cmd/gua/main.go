package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mileusna/useragent"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println()
		ua := useragent.Parse(scanner.Text())
		fmt.Println(ua.String)
		fmt.Println(strings.Repeat("=", len(ua.String)))
		fmt.Println("Browser:", ua.Name, "v", ua.Version)
		fmt.Println("OS:", ua.OS, "v", ua.OSVersion)
		fmt.Print("Device: ", ua.Device)
		if ua.Mobile {
			fmt.Println(" (Mobile)")
		}
		if ua.Tablet {
			fmt.Println(" (Tablet)")
		}
		if ua.Desktop {
			fmt.Println(" (Desktop)")
		}
		if ua.Bot {
			fmt.Println(" (Bot)")
		}
		if ua.URL != "" {
			fmt.Println(ua.URL)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
