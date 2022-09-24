/*
Copyright © 2022 Kosuke Nakamura <ncl0709@gmail.com>

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
	"os"
	"os/user"

	"github.com/k0kubun/sqldef"
	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

var (
	options = sqldef.Options{}
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup uniar",
	Long:  "Setup your member status and scene card collections for uniar",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		uniarPath := user.HomeDir + "/.uniar"
		dbPath := uniarPath + "/uniar.db"

		if _, err := os.Stat(uniarPath); err != nil {
			if err := os.Mkdir(uniarPath, 0750); err != nil {
				fmt.Println(err)
			}
		}

		setupMigrate(dbPath)
		setupSeed(dbPath)

		ctx := context.Background()
		db, err := repository.NewConnection()
		if err != nil {
			fmt.Println(err)
			return
		}
		q := repository.New()

		setupMember(ctx, db, q)
		setupOffice(ctx, db, q)
		setupScene(ctx, db, q)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
