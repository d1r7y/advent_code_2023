/*
Copyright © 2022 Cameron Esfahani
*/
package day02

import (
	"io"
	"log"
	"os"

	"github.com/d1r7y/advent_2023/cmd"
	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "Day 2 of Advent of Code 2023",
	Long:  `Cube Conundrum`,
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
		err = day02(string(fileContent))
		if err != nil {
			log.Fatal(err)
		}
	},
}

var inputPath string

func init() {
	cmd.RootCmd.AddCommand(day02Cmd)

	day02Cmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input file")
}
