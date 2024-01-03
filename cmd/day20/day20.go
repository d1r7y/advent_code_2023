/*
Copyright © 2023 Cameron Esfahani <dirty@mac.com>
*/
package day20

import (
	"io"
	"log"
	"os"

	"github.com/d1r7y/advent_2023/cmd"
	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use:   "day20",
	Short: "Day 20 of Advent of Code 2023",
	Long:  `Grove Positioning System`,
	Run: func(cmd *cobra.Command, args []string) {
		df, err := os.Open(inputPath)
		if err != nil {
			log.Fatal(err)
		}

		defer df.Close()

		fileContent, err := io.ReadAll(df)
		if err != nil {
			log.Fatal(err)
		}
		err = day20(string(fileContent))
		if err != nil {
			log.Fatal(err)
		}
	},
}

var inputPath string

func init() {
	cmd.RootCmd.AddCommand(day20Cmd)

	day20Cmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input file")
}
