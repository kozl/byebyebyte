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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	flag "github.com/spf13/pflag"
)

// RootCmd represents the base command when called without any subcommands
var (
	requiredFlagsMissing bool
	startByte, endByte   int64
	inFile, outFile      string

	RootCmd = &cobra.Command{
		Use:   "byebyebyte",
		Short: "Destroy your files in the name of art",
		Long:  `A command line tool that replaces random bytes in a file with random values.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			validateFlags(flags)

			inFile, _ = flags.GetString("input")
			outFile, _ = flags.GetString("output")
			st, err := os.Stat(inFile)
			if err != nil {
				return errors.New("Input file not found")
			}
			size := st.Size()
			startByte, _ = flags.GetInt64("start")
			min, _ := flags.GetFloat32("min")
			if startByte == 0 && min != 0 {
				startByte = int64(min * float32(size))
			}
			endByte, _ = flags.GetInt64("stop")
			max, _ := flags.GetFloat32("max")
			if endByte == 0 && max != 0 {
				endByte = int64(max * float32(size))
			}

			if (startByte > endByte) || (endByte > size) {
				return errors.New("Alter range out of bounds")
			}

			return nil
		},
	}
)

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringP("input", "i", "", "input file path")
	RootCmd.MarkFlagFilename("input")
	RootCmd.PersistentFlags().StringP("output", "o", "", "output file path")
	RootCmd.PersistentFlags().Float32("min", 0, `the lower bound of the file to alter bytes in - use percentage 0 to 1 (ex: 0.15 = 15%, 1 = 100%).
If specified, you cannot use --start or --stop`)
	RootCmd.PersistentFlags().Float32("max", 0, `the upper bound of the file to alter bytes in - use percentage from 0 to 1 (ex: 0.15 = 15%, 1 = 100%).
If specified, you cannot use --start or --stop`)
	RootCmd.PersistentFlags().Int64("start", 0, `a specific point at the file, in bytes, at which to begin altering bytes.
If specified, you cannot use --min or --max`)
	RootCmd.PersistentFlags().Int64("stop", 0, `a specific point at the file, in bytes, at which to stop altering bytes.
If specified, you cannot use --min or --max`)
}

func validateFlags(flagset *flag.FlagSet) error {
	if (flagset.Changed("min")) && (flagset.Changed("start")) {
		return errors.New("Conflicting flags --min and --start")
	}
	if (flagset.Changed("max")) && (flagset.Changed("stop")) {
		return errors.New("Conflicting flags --end and --stop")
	}
	if (!flagset.Changed("input")) || (!flagset.Changed("output")) {
		requiredFlagsMissing = true
	}
	if (!flagset.Changed("min")) && (!flagset.Changed("start")) {
		requiredFlagsMissing = true
	}
	if (!flagset.Changed("min")) && (!flagset.Changed("start")) {
		requiredFlagsMissing = true
	}
	if requiredFlagsMissing {
		return errors.New("Missing required flags")
	}
	return nil
}