/*
Copyright © 2022 Cameron Esfahani
*/
package day12

import (
	"io"
	"log"
	"os"

	"github.com/d1r7y/advent_2023/cmd"
	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use:   "day12",
	Short: "Day 12 of Advent of Code 2023",
	Long:  `Hill Climbing Algorithm`,
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
		err = day12(string(fileContent))
		if err != nil {
			log.Fatal(err)
		}
	},
}

var inputPath string

func init() {
	cmd.RootCmd.AddCommand(day13Cmd)

	day13Cmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input file")
}
