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

	"github.com/litencatt/uniar/repository"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup uniar",
	Long:  "Setup your member status and scene card collections for uniar",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := setup(); err != nil {
			return err
		}
		return nil
	},
}

func setup() error {
	fmt.Println("Start setup")
	ctx := context.Background()

	dbPath := GetDbPath()
	fmt.Println(dbPath)
	db, err := repository.NewConnection(dbPath)
	if err != nil {
		fmt.Println("setup error 1")
		fmt.Println(err)
		return err
	}
	q := repository.New()

	if err := setupMkdir(); err != nil {
		fmt.Println("setup error 2")
		return err
	}
	if err := migrate(ctx, db, dbPath); err != nil {
		fmt.Println("setup error 3")
		return err
	}
	if err := initProducerScene(ctx, db, q); err != nil {
		fmt.Println("setup error 4")
		return err
	}
	fmt.Println("End setup")
	return nil
}

func GetDbPath() string {
	if p, ok := os.LookupEnv("UNIAR_DB_PATH"); ok {
		return p
	}
	uniarPath := GetUniarPath()

	return uniarPath + "/uniar.db"
}

func GetUniarPath() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.HomeDir + "/.uniar"
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
