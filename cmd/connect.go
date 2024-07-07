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
	"net"

	"github.com/spf13/cobra"
)

type ConnectParams struct {
	ip    string
	port  uint
	proto string
}

var connectParams = &ConnectParams{}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "dial <IP> <Port>",
	Short: "opens a connection to an <IP> <Port>",
	Run: func(cmd *cobra.Command, args []string) {

		net.Dial(connectParams.proto, fmt.Sprintf("%s:%d", connectParams.ip, connectParams.port))
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&connectParams.proto, "proto", "p", "tcp", "Help message for toggle")
}
