package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorings = 3
const delay = 5 * time.Second

func main() {
	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("I don't know this command")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	name := "Jon"
	version := 1.2
	fmt.Println("Hello,", name)
	fmt.Println("This is version", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)
	fmt.Println("")

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")

	// This is a slice
	/*sites := []string{"https://random-status-code.herokuapp.com/",
"https://www.alura.com.br", "https://www.caelum.com.br"}*/

	// The for loop can be done like this
	/*for i := 0; i < len(sites); i++ {
		fmt.Println(site[i])
	}*/

	sites := readSitesFile()

	for i := 0; i < monitorings; i++ {
		// Or like this
		for i, site := range sites {
			fmt.Println("Site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay)
		fmt.Println("")
	}

}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully!")
		writeLog(site, true)
	} else {
		fmt.Println("Site:", site, "is facing issues. Status Code:", resp.StatusCode)
		writeLog(site, false)
	}
}

func readSitesFile() []string {
	var sites []string

	// Reads the entire file
	// file, err := ioutil.ReadFile("sites.txt")

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		sites = append(sites, row)

		if err ==	io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func writeLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	fmt.Println("Showing Logs...")

	file, err := os.Open("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		fmt.Println(row)

		if err == io.EOF {
			break
		}
	}

	file.Close()
}