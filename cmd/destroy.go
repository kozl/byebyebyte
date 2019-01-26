// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"math/rand"
	"os"
	"time"
)

const bufSize = 1024 * 4

// destroyCmd represents the destroy command
var (
	prob float32
	err  error

	destroyCmd = &cobra.Command{
		Use:   "destroy",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			prob, err = flags.GetFloat32("probability")
			if err != nil {
				return errors.New("Missing probability flag")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			inF, err := os.Open(inFile)
			if err != nil {
				return err
			}
			defer inF.Close()

			outF, err := os.Create(outFile)
			if err != nil {
				return err
			}
			defer outF.Close()

			buf := make([]byte, bufSize)
			rand.Seed(int64(time.Now().Unix()))

			var currentPos int64

			for {
				n, err := io.ReadFull(inF, buf)
				if err == io.EOF {
					break
				}
				if currentPos > startByte && currentPos < endByte {
					for i := range buf[:n] {
						if rand.Float32() < prob {
							rand.Read(buf[i : i+1])
						}
					}
				}
				outF.Write(buf)
				currentPos += int64(n)
			}
			return nil
		},
	}
)

func init() {
	destroyCmd.Flags().Float32P("probability", "p", 0.0, "Probability of destroyng a byte")
	RootCmd.AddCommand(destroyCmd)
}
