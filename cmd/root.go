package cmd

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var (
	apiKey    string
	channelID string
)

var rootCmd = &cobra.Command{
	Use:   "slack-send-message [message]",
	Short: "Send a message to a Slack channel",
	Long:  `Send a message to a specific Slack channel using the Slack API and return the thread timestamp.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get message from positional argument
		message := args[0]

		// Get API key from parameter or environment variable (parameter has priority)
		key := apiKey
		if key == "" {
			key = os.Getenv("SLACK_API_KEY")
		}
		if key == "" {
			return fmt.Errorf("api-key is required (use --api-key flag or SLACK_API_KEY environment variable)")
		}

		if channelID == "" {
			return fmt.Errorf("channel-id is required")
		}

		// Create Slack client
		api := slack.New(key)

		// Send message
		_, timestamp, err := api.PostMessage(
			channelID,
			slack.MsgOptionText(message, false),
		)
		if err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}

		// Output thread-ts
		fmt.Println(timestamp)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&apiKey, "api-key", "", "Slack API key (can also use SLACK_API_KEY env var)")
	rootCmd.Flags().StringVar(&channelID, "channel-id", "", "Slack channel ID (required)")
	
	rootCmd.MarkFlagRequired("channel-id")
}
