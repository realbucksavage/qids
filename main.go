package main

import (
	"fmt"
	"os"

	"github.com/realbucksavage/qids/pkg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "qids",
		Short: "A quick ID generation utility",
	}

	ulidCmd = &cobra.Command{
		Use:   "ulid",
		Short: "Generate a ULID (https://github.com/ulid/spec)",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := pkg.NewULID()
			if err != nil {
				return err
			}

			fmt.Println(id)
			return nil
		},
	}

	uuidV1Cmd = &cobra.Command{
		Use:   "uuidv1",
		Short: "Generate a UUIDv1 (https://datatracker.ietf.org/doc/html/rfc4122)",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := pkg.NewV1UUID()
			if err != nil {
				return err
			}

			fmt.Println(id)
			return nil
		},
	}

	uuidV4Cmd = &cobra.Command{
		Use:   "uuid",
		Short: "Generate a UUIDv4 (https://datatracker.ietf.org/doc/html/rfc4122)",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := pkg.NewV4UUID()
			if err != nil {
				return err
			}

			fmt.Println(id)
			return nil
		},
	}
)

func main() {

	rootCmd.AddCommand(uuidV4Cmd)
	rootCmd.AddCommand(uuidV1Cmd)
	rootCmd.AddCommand(ulidCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "generation error: %v", err)
		os.Exit(1)
	}
}
