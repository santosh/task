package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/santosh/cotu/integration"
	"github.com/spf13/cobra"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "interface with cotu s3 bucket",
}

var uploadS3Cmd = &cobra.Command{
	Use:   "upload",
	Short: "uploads a file to cotu bucket",
	Run: func(cmd *cobra.Command, args []string) {
		err := integration.UploadFile(args[0])
		if err != nil {
			panic(err)
		}
	},
}

var listS3Cmd = &cobra.Command{
	Use:   "list",
	Short: "list items from cotu bucket",
	Run: func(cmd *cobra.Command, args []string) {
		err := integration.ListFiles()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	s3Cmd.AddCommand(uploadS3Cmd, listS3Cmd)
	RootCmd.AddCommand(s3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
