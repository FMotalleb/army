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
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/FMotalleb/army/log"
	"github.com/spf13/cobra"
)

type ConnectParams struct {
	ip    string
	port  uint
	proto string
	zero  bool
}

var dialParams = &ConnectParams{}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "dial <IP> <Port>",
	Short: "opens a connection to an <IP> <Port>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			log.Warnf("excpected to have 2 args but received %d from: %v", len(args), args)
			return errors.New("unexcpected arguments")
		}
		port, err := strconv.ParseInt(args[1], 10, 16)
		if err != nil {
			log.Warnf("port number is invalid: %d, %s", args[1], err)
			return err
		}
		dialParams.ip = args[0]
		dialParams.port = uint(port)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		slog.SetDefault(slog.Default().WithGroup("dialer"))
		addr := fmt.Sprintf("%s:%d", dialParams.ip, dialParams.port)
		log.Debugf("dialing %s", addr)
		conn, err := net.Dial(dialParams.proto, addr)
		if err != nil {
			log.Warnf("dial failed: %v", err)
			return
		}
		slog.Debug("Connection established")
		if dialParams.zero {
			wg := new(sync.WaitGroup)
			wg.Add(1)
			go func() {
				_, err = conn.Read([]byte{0})
				log.Debugf("received single byte of data from socket, errors: %v", err)
				err = conn.Close()
				if err != nil {
					log.Warnf("an error accrued when tried to close the connection: %s", err)
				} else {
					slog.Info("connection closed")
				}
				wg.Done()
			}()
			time.Sleep(time.Second * 1)
			wg.Add(1)
			go func() {
				_, err := conn.Write([]byte{0})
				log.Debugf("sent single byte of data to socket, errors: %v", err)
				wg.Done()
			}()
			wg.Wait()

		} else {
			wg := new(sync.WaitGroup)
			slog.Debug("socket's output and input are attached to stdout and stdin")
			wg.Add(1)
			go func() {

				count, err := io.Copy(conn, os.Stdin)
				log.Debugf("wrote %d bytes from stdin onto the connection's input, errors: %#v", count, err)
				wg.Done()
			}()
			wg.Add(1)
			go func() {
				count, err := io.Copy(os.Stdout, conn)
				log.Debugf("wrote %d bytes from connection's output into stdout, errors: %#v", count, err)
				wg.Done()
			}()
			wg.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVarP(&dialParams.proto, "proto", "p", "tcp", "protocol used to access address")
	connectCmd.Flags().BoolVarP(&dialParams.zero, "zero", "z", false, "Help message for toggle")
}
