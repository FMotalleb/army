/*
Copyright © 2024 Motalleb Fallahnezhad

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/FMotalleb/army/logutils"
	"github.com/spf13/cobra"
)

type CmdConfig struct {
	Verbose bool
	Trace   bool
}

var cmdConfig = &CmdConfig{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "army",
	Short: "A swiss-army knife",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		handler := logutils.NewLogHandler(
			"army",
			cmdConfig.Verbose,
			cmdConfig.Trace,
		)
		logger := slog.New(handler)

		slog.SetDefault(logger)
		slog.Debug("debug log enabled")
	},
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&cmdConfig.Verbose, "verbose", "v", false, "Set log level to verbose")
	rootCmd.PersistentFlags().BoolVar(&cmdConfig.Trace, "trace", false, "Set log level to verbose")
}
