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
	"fmt"
	"log/slog"
	"net"

	"github.com/spf13/cobra"
)

type ConnectParams struct {
	ip    string
	port  uint
	proto string
}

var dialParams = &ConnectParams{}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "dial <IP> <Port>",
	Short: "opens a connection to an <IP> <Port>",
	Run: func(cmd *cobra.Command, args []string) {
		slog.SetDefault(slog.Default().WithGroup("connect"))
		addr := fmt.Sprintf("%s:%d", dialParams.ip, dialParams.port)
		slog.Debug("dialing", addr)
		net.Dial(dialParams.proto, addr)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVarP(&dialParams.proto, "proto", "p", "tcp", "Help message for toggle")
}
