package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	RELEASE_VERSION = "1.0.0"
	RELEASE_DATE    = "2023-02-28"
)

func main() {
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s %s (by %s %s)", "pingalive", RELEASE_VERSION, "MBODM", RELEASE_DATE))
	fmt.Println()
	if runtime.GOOS != "windows" {
		printError("Non-Windows platforms are not supported by this executable.")
		os.Exit(3)
	}
	if len(os.Args) > 1 {
		printError("Command line arguments are not supported by this executable.")
		os.Exit(2)
	}
	fmt.Println("A tiny Windows command line tool, pinging Telekom´s main DNS server.")
	fmt.Println()
	params := strings.Fields("/C start ping.exe -t 194.25.2.129")
	if err := exec.Command("cmd.exe", params...).Start(); err != nil {
		printError("Could not execute cmd.exe to start the ping command.")
		os.Exit(1)
	}
	fmt.Println(" - Successfully started ping command (ping -t 194.25.2.129) in a separate window.")
	fmt.Println(" - The ping command is running until you press CTRL+C there or close that window.")
	fmt.Println()
	fmt.Println("Have a nice day.")
	os.Exit(0)
}

func printError(message string) {
	// Not printing to stderr is done on purpose here.
	fmt.Println("Error:", message)
	fmt.Println()
	fmt.Println("Take a look at https://github.com/mbodm/pingalive for more information.")
}
