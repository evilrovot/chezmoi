package cmd

import (
	"github.com/absfs/afero"
	"github.com/spf13/cobra"
)

var forgetCommand = &cobra.Command{
	Use:   "forget",
	Args:  cobra.MinimumNArgs(1),
	Short: "Forget a file",
	Run:   makeRun(runForgetCommand),
}

func init() {
	rootCommand.AddCommand(forgetCommand)
}

func runForgetCommand(command *cobra.Command, args []string) error {
	// FIXME support directories
	fs := afero.NewOsFs()
	targetState, err := getTargetState(fs)
	if err != nil {
		return err
	}
	sourceFileNames, err := getSourceFileNames(targetState, args)
	if err != nil {
		return err
	}
	for _, sourceFileName := range sourceFileNames {
		if err := fs.Remove(sourceFileName); err != nil {
			return err
		}
	}
	return nil
}