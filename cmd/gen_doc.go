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
	"os"
	"os/exec"
	"text/template"

	"github.com/spf13/cobra"
)

var genDocCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate uniar document",
	RunE: func(cmd *cobra.Command, args []string) error {
		readme := `
# uniar

{{ .bt }}uniar{{ .bt }} is database and management your scene collections CLI tool for [UNI'S ON AIR](https://keyahina-unisonair.com/).

## Usage		

{{ .bt3 }}
{{ .uniar }}
{{ .bt3 }}

## Install

{{ .bt3 }}
$ brew tap litencatt/tap
$ brew install litencatt/tap/uniar
{{ .bt3 }}

## Commands

### List

{{ .bt3 }}
$ uniar list
{{ .uniarList }}
{{ .bt3 }}

### Usage

{{ .bt3 }}
$ uniar list scene -h
{{ .uniarListScene }}
{{ .bt3 }}

{{ .bt3 }}
$ uniar list scene -f | head
{{ .uniarListSceneEg }}
{{ .bt3 }}

{{ .bt3 }}
$ uniar list scene -f -c Blue
{{ .uniarListSceneColor }}
{{ .bt3 }}

{{ .bt3 }}
$ uniar list scene -f -m 加藤史帆
{{ .uniarListSceneMember }}
{{ .bt3 }}

{{ .bt3 }}
$ uniar list scene -f -p キュン
{{ .uniarListScenePhoto }}
{{ .bt3 }}

{{ .bt3 }}
$ uniar list scene -f -c Blue -m 加藤史帆 -p キュン
{{ .uniarListSceneCombine }}
{{ .bt3 }}

{{ .bt3 }}
$ uniar list scene --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head
{{ .uniarListSceneIgnore }}
{{ .bt3 }}

{{ .bt3 }}
$ uniar list scene -d --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head
{{ .uniarListSceneDetail }}
{{ .bt3 }}
`
		ua, _ := exec.Command("/bin/bash", "-c", "./uniar").Output()
		ual, _ := exec.Command("/bin/bash", "-c", "./uniar", "list").Output()
		ualsh, _ := exec.Command("/bin/bash", "-c", "./uniar list scene -h").Output()
		ualseg, _ := exec.Command("/bin/bash", "-c", "./uniar list scene -f | head").Output()
		ualsc, _ := exec.Command("/bin/bash", "-c", "./uniar list scene -f -c Blue | head").Output()
		ualsm, _ := exec.Command("/bin/bash", "-c", "./uniar list scene -f -m 加藤史帆 | head").Output()
		ualsp, _ := exec.Command("/bin/bash", "-c", "./uniar list scene -f -p キュン | head").Output()
		ualscb, _ := exec.Command("/bin/bash", "-c", "./uniar list scene -f -c Blue -m 加藤史帆 -p キュン").Output()
		ualsi, _ := exec.Command("/bin/bash", "-c", "./uniar list scene --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head").Output()
		ualsd, _ := exec.Command("/bin/bash", "-c", "./uniar list scene -d --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head").Output()
		v := map[string]interface{}{
			"bt":                    "`",
			"bt3":                   "```",
			"uniar":                 string(ua),
			"uniarList":             string(ual),
			"uniarListScene":        string(ualsh),
			"uniarListSceneEg":      string(ualseg),
			"uniarListSceneColor":   string(ualsc),
			"uniarListSceneMember":  string(ualsm),
			"uniarListScenePhoto":   string(ualsp),
			"uniarListSceneCombine": string(ualscb),
			"uniarListSceneIgnore":  string(ualsi),
			"uniarListSceneDetail":  string(ualsd),
		}
		template, _ := template.New("").Parse(readme)
		if err := template.Execute(os.Stdout, v); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genDocCmd)
}
