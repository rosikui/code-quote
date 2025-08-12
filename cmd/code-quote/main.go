package cmd

import (
	"embed"
	"encoding/json"
	"log"

	"github.com/rosikui/code-quote/internal/quote"
	"github.com/rosikui/code-quote/internal/term"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fs embed.FS

func Execute(quotesFS embed.FS) error {
	fs = quotesFS
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "code-quote",
	Short: "Display random or daily programming quotes in your terminal",
	Long: `code-quote is a CLI tool that displays inspiring programming quotes.

Your terminal deserves more than just errors and logs — a fresh code quote for every session.`,
	RunE: runQuote,
}

func init() {
	// Local flags
	rootCmd.Flags().String("file", "", "Additional quotes file(s) (JSON/YAML, comma-separated)")
	rootCmd.Flags().String("tag", "", "Filter by tags (comma-separated)")
	rootCmd.Flags().String("lang", "", "Filter by language")
	rootCmd.Flags().Bool("daily", false, "Daily quote (same quote for the same day)")
	rootCmd.Flags().Bool("no-color", false, "Disable ANSI colors")
	rootCmd.Flags().Bool("markdown", false, "Output in Markdown format")
	rootCmd.Flags().Bool("show-tags", false, "Show tags after author")

	// Bind flags to viper
	viper.BindPFlag("file", rootCmd.Flags().Lookup("file"))
	viper.BindPFlag("tag", rootCmd.Flags().Lookup("tag"))
	viper.BindPFlag("lang", rootCmd.Flags().Lookup("lang"))
	viper.BindPFlag("daily", rootCmd.Flags().Lookup("daily"))
	viper.BindPFlag("no-color", rootCmd.Flags().Lookup("no-color"))
	viper.BindPFlag("markdown", rootCmd.Flags().Lookup("markdown"))
	viper.BindPFlag("show-tags", rootCmd.Flags().Lookup("show-tags"))
}

func runQuote(cmd *cobra.Command, args []string) error {
	// Get values from viper
	file := viper.GetString("file")
	tags := viper.GetString("tag")
	lang := viper.GetString("lang")
	daily := viper.GetBool("daily")
	noColor := viper.GetBool("no-color")
	markdown := viper.GetBool("markdown")
	showTags := viper.GetBool("show-tags")

	// Load embedded quotes
	data, err := fs.ReadFile("quotes/quotes.en.json")
	if err != nil {
		return err
	}

	var quotes []quote.Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return err
	}

	// Load additional files if specified
	if file != "" {
		loader := quote.NewLoader()
		additionalQuotes, err := loader.LoadFromFiles(file)
		if err != nil {
			return err
		}
		quotes = append(quotes, additionalQuotes...)
	}

	// Create picker and apply filters
	picker := quote.NewPicker(quotes)

	if tags != "" {
		picker = picker.FilterByTags(tags)
	}

	if lang != "" {
		picker = picker.FilterByLang(lang)
	}

	// Check if we have any quotes after filtering
	if picker.Count() == 0 {
		log.Fatal("No quotes available after filtering")
	}

	// Select quote
	var selectedQuote *quote.Quote
	if daily {
		selectedQuote, err = picker.PickDaily()
	} else {
		selectedQuote, err = picker.PickRandom()
	}

	if err != nil {
		return err
	}

	// Render and output
	renderer := term.NewRenderer(noColor, markdown, showTags)
	output := renderer.Render(selectedQuote)
	cmd.Println(output)
	return nil
}
