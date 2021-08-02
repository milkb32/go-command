package cmds

import "github.com/urfave/cli"

var (
	commands []cli.Command
)

func Register(cmd cli.Command) {
	commands = append(commands, cmd)
}

func GetCommands() []cli.Command {
	return commands
}
