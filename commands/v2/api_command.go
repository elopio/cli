package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

type ApiCommand struct {
	OptionalArgs      flags.APITarget `positional-args:"yes"`
	SkipSSLValidation bool            `long:"skip-ssl-validation" description:"Skip verification of the API endpoint. Not recommended!"`
	Unset             bool            `long:"unset" description:"Remove all api endpoint targeting"`
	usage             interface{}     `usage:"CF_NAME api [URL]"`
	relatedCommands   interface{}     `related_commands:"auth, login, target"`
}

func (_ ApiCommand) Setup(config commands.Config, ui commands.UI) error {
	return nil
}

func (_ ApiCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
