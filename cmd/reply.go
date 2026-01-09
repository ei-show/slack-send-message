package cmd

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var (
	replyApiKey  string
	replyThreadTS string
	replyMessage  string
)

var replyCmd = &cobra.Command{
	Use:   "reply",
	Short: "Reply to a thread in Slack",
	Long:  `Send a reply message to an existing thread using the Slack API.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get API key from parameter or environment variable (parameter has priority)
		key := replyApiKey
		if key == "" {
			key = os.Getenv("SLACK_API_KEY")
		}
		if key == "" {
			return fmt.Errorf("api-key is required (use --api-key flag or SLACK_API_KEY environment variable)")
		}

		if replyThreadTS == "" {
			return fmt.Errorf("thread-ts is required")
		}

		if replyMessage == "" {
			return fmt.Errorf("message is required")
		}

		// Create Slack client
		api := slack.New(key)

		// Extract channel ID from the thread-ts context
		// In Slack API, to reply to a thread, we need both channel ID and thread_ts
		// For simplicity, we'll use the same channel-id flag from root command
		// But we need to get the channel ID somehow - let's add it as a flag
		channelIDForReply := channelID
		if channelIDForReply == "" {
			return fmt.Errorf("channel-id is required for reply")
		}

		// Send reply message
		_, timestamp, err := api.PostMessage(
			channelIDForReply,
			slack.MsgOptionText(replyMessage, false),
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

	replyCmd.Flags().StringVar(&replyApiKey, "api-key", "", "Slack API key (can also use SLACK_API_KEY env var)")
	replyCmd.Flags().StringVar(&channelID, "channel-id", "", "Slack channel ID (required)")
	replyCmd.Flags().StringVar(&replyThreadTS, "thread-ts", "", "Thread timestamp to reply to (required)")
	replyCmd.Flags().StringVar(&replyMessage, "message", "", "Message to send as reply (required)")
	
	replyCmd.MarkFlagRequired("channel-id")
	replyCmd.MarkFlagRequired("thread-ts")
	replyCmd.MarkFlagRequired("message")
}
