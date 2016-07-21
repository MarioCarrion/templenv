// Copyright © 2016 Mario Carrion <mario@carrion.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/MarioCarrion/dynamic-docker-composer/ddclib"
	"github.com/spf13/cobra"
)

var initName string
var initOverwrite bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes and selects the workspace to use",
	Long: `Selects or initializes a workspace, if the workspace does not exist then
it is created. Provide the "--overwrite" argument to overwrite previous work.

Examples:

ddc-mc init --name workspace1
ddc-mc init --name workspace1 --overwrite
`,
	Run: func(cmd *cobra.Command, args []string) {
		ddclib.NewInitCommand(initName).Execute(initOverwrite)
	},
}

func init() {
	initCmd.Flags().StringVarP(&initName, "name", "n", "", "workspace name")
	initCmd.Flags().BoolVar(&initOverwrite, "overwrite", false, "overwrite existing workspace, if any")

	rootCmd.AddCommand(initCmd)
}
