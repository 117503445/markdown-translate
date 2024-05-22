package mdtran

import (
	"os"
	"time"

	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/117503445/markdown-translate/pkg/translator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type AppConfig struct {
	InputFile string
	Provider  string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mdtran",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal().Msg("No input file, see --help")
		}

		if len(args) > 1 {
			log.Fatal().Msg("Too many arguments, see --help")
		}

		input := args[0]

		log.Debug().Str("input", input).Msg("rootCMD")

		cfg := &AppConfig{
			InputFile: input,
			Provider:  "mock",
		}

		// output, err translator.Translate(input)
		inputText, err := os.ReadFile(cfg.InputFile)
		if err != nil {
			log.Fatal().Err(err).Msg("read input file")
		}

		// first try to use input file name + .out.md
		// if exists, use input file name + date + .out.md
		// if still exists, error
		var output string
		if _, err := os.Stat(cfg.InputFile + ".out.md"); os.IsNotExist(err) {
			output = cfg.InputFile + ".out.md"
		} else {
			date := time.Now().Format("20060102-150405")
			output = cfg.InputFile + "." + date + ".out.md"
			if _, err := os.Stat(output); err == nil {
				log.Fatal().Strs("output", []string{cfg.InputFile + ".out.md", output}).Msg("output file exists")
			}
		}

		provider, err := provider.GetProvider(cfg.Provider)
		if err != nil {
			log.Fatal().Err(err).Msg("provider not found")
		}

		translator := translator.NewTranslator(provider)

		outputText, err := translator.Translate(string(inputText))
		if err != nil {
			log.Fatal().Err(err).Msg("translate failed")
		}

		err = os.WriteFile(output, []byte(outputText), 0644)
		if err != nil {
			log.Fatal().Err(err).Msg("write output file")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mdtran.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// input file as parameter
	// rootCmd.Flags().StringP("input", "i", "", "input file")
}
