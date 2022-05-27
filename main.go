package raspberry

import "os"
import "fmt"

type Cli struct {
	AcceptedCommands []string
	HelpMsg          string
	Version          float64
}

var (
	Command   string
	argLength int
	Args      []string
)

type fn func()

func (c *Cli) SetHandler(cmd string, fun fn) {
	if Command == cmd {
		fun()
	}
}
func contains(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func (c *Cli) PrintHelp() {
	fmt.Println(c.HelpMsg)
}
func (c *Cli) PrintVersion() {
	fmt.Println(c.Version)
}
func shiftArgsDownward(arr []string) []string {
	arr = arr[2:]
	return arr
}
func (c *Cli) Setup() {
	argLength = len(os.Args)
	if argLength > 1 {
		if contains(os.Args[1], c.AcceptedCommands) {
			Command = os.Args[1]
			Args = os.Args
			shiftArgsDownward(Args)
		} else {
			fmt.Println("Command not found: ", os.Args[1])
		}
	} else {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	// set default cmd's
	c.SetHandler("-v", c.PrintVersion)
	c.SetHandler("version", c.PrintVersion)
	c.SetHandler("-h", c.PrintHelp)
	c.SetHandler("help", c.PrintHelp)
}
