package cmd

import (
	"fmt"
	"github.com/perottobc/mvn-pom-mutator/pkg/pom"
	"github.com/spf13/cobra"
	"log"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "parses and writes a maven pom.xml file",
	Long:  `parses and writes a maven pom.xml file`,
	Run: func(cmd *cobra.Command, args []string) {
		targetDirectory, err := cmd.Flags().GetString("target")
		if err != nil {
			log.Fatalln(err)
		}

		model, err := pom.GetModelFrom(fmt.Sprintf("%s/pom.xml", targetDirectory))

		if err != nil {
			log.Fatalln(err)
		}

		err = model.WriteToFile(fmt.Sprintf("%s/pom.xml.new", targetDirectory))

		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("target", ".", "Optional target directory")
}
