/*
Copyright © 2024 Nev3r Kn0wn <nn0wn@proton.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/pilab-dev/ion/pkg/client"
	v1 "github.com/pilab-dev/ion/pkg/gen/pb/room/v1"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

var (
	password string
	lock     bool
	maxPeers uint32
	sid      string
)

// roomCmd represents the serve command
var roomCreateCmd = &cobra.Command{
	Use:     "create [name] [description]",
	Aliases: []string{"c"},
	Short:   "Create a room",
	Long:    `Create a room in the ion stack. If a description is provided, it will be displayed in the room list.`,
	Args:    cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewRoomClient(serverAddr)
		if err != nil {
			return err
		}

		description := ""
		if len(args) == 2 {
			description = args[1]
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		fmt.Printf("Creating room...  Server [%s]\n", serverAddr)

		if sid == "" {
			sid = xid.New().String()
		}

		_, err = c.CreateRoom(ctx, &v1.CreateRoomRequest{
			Room: &v1.Room{
				Sid:         sid,
				Name:        args[0],
				Lock:        lock,
				Password:    password,
				Description: description,
				MaxPeers:    uint32(maxPeers),
			},
		})
		if err != nil {
			return fmt.Errorf("failed to create room: %w", err)
		}

		return nil
	},
}

func init() {
	roomCmd.AddCommand(roomCreateCmd)

	roomCreateCmd.Flags().StringVarP(&password, "password", "p", "", "password for the room")
	roomCreateCmd.Flags().BoolVarP(&lock, "lock", "l", false, "lock the room")
	roomCreateCmd.Flags().Uint32VarP(&maxPeers, "max", "m", 0, "maximum number of peers allowed in the room")
	roomCreateCmd.Flags().StringVarP(&sid, "sid", "s", "", "room id")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
