/*
Copyright © 2022 Cameron Esfahani
*/
package day03

import (
	"io"
	"log"
	"os"

	"github.com/d1r7y/advent_2023/cmd"
	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Day 3 of Advent of Code 2023",
	Long:  `Gear Ratios`,
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
		err = day03(string(fileContent))
		if err != nil {
			log.Fatal(err)
		}
	},
}

var inputPath string

func init() {
	cmd.RootCmd.AddCommand(day03Cmd)

	day03Cmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input file")
}
