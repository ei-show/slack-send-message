package cmd

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var (
	replyAPIKey    string
	replyChannelID string
	replyThreadTS  string
)

var replyCmd = &cobra.Command{
	Use:   "reply [message]",
	Short: "Reply to a thread in Slack",
	Long:  `Send a reply message to an existing thread using the Slack API.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get message from positional argument
		message := args[0]

		// Get API key from parameter or environment variable (parameter has priority)
		key := replyAPIKey
		if key == "" {
			key = os.Getenv("SLACK_API_KEY")
		}
		if key == "" {
			return fmt.Errorf("api-key is required (use --api-key flag or SLACK_API_KEY environment variable)")
		}

		if replyChannelID == "" {
			return fmt.Errorf("channel-id is required for reply")
		}

		if replyThreadTS == "" {
			return fmt.Errorf("thread-ts is required")
		}

		// Create Slack client
		api := slack.New(key)

		// Send reply message
		_, timestamp, err := api.PostMessage(
			replyChannelID,
			slack.MsgOptionText(message, false),
			slack.MsgOptionTS(replyThreadTS),
		)
		if err != nil {
			return fmt.Errorf("failed to send reply: %w", err)
		}

		// Output thread-ts
		fmt.Println(timestamp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(replyCmd)

	replyCmd.Flags().StringVar(&replyAPIKey, "api-key", "", "Slack API key (can also use SLACK_API_KEY env var)")
	replyCmd.Flags().StringVar(&replyChannelID, "channel-id", "", "Slack channel ID (required)")
	replyCmd.Flags().StringVar(&replyThreadTS, "thread-ts", "", "Thread timestamp to reply to (required)")

	_ = replyCmd.MarkFlagRequired("channel-id")
	_ = replyCmd.MarkFlagRequired("thread-ts")
}
