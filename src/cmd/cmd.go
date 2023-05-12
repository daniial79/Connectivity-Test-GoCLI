package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/daniial79/Connectivity-Test-GoCLI/src/utils"
)

func LoadProgram() {
	fmt.Println("Connetion Check CLI")
	fmt.Println("COMMAND: checkinfyconnection <ip-address> | 'exit' to quit the program")
	fmt.Println("-------------------")
	fmt.Print("=> ")
}

func FetchIp() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		userInput := scanner.Text()

		if utils.IsExit(userInput) {
			fmt.Println("Bye :)")
			os.Exit(0)
		}

		slicedInputData := strings.Split(userInput, " ")

		if utils.CommandIsValid(slicedInputData[0]) {
			pingResult := execPingCommand(slicedInputData[1])
			fmt.Println(pingResult)
		} else {
			fmt.Println("command is not valid")
		}
		fmt.Print("=> ")
	}
}

func execPingCommand(cmdParam string) string {
	data, err := exec.Command("cmd", "/c", "ping", "-n", "3").CombinedOutput()

	if err != nil {
		return "pinging has failed!"
	}

	return string(data)
}
