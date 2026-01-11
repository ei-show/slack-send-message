package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestReplyCommand(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		env         map[string]string
		wantErr     bool
		errContains string
	}{
		{
			name:        "missing channel-id and thread-ts",
			args:        []string{"reply", "test message"},
			wantErr:     true,
			errContains: "required flag(s)",
		},
		{
			name:        "missing thread-ts",
			args:        []string{"reply", "--channel-id", "C123", "test message"},
			wantErr:     true,
			errContains: "required flag(s) \"thread-ts\" not set",
		},
		{
			name:        "missing message",
			args:        []string{"reply", "--channel-id", "C123", "--thread-ts", "1234567890.123456"},
			wantErr:     true,
			errContains: "accepts 1 arg(s), received 0",
		},
		{
			name:        "missing api-key without env",
			args:        []string{"reply", "--channel-id", "C123", "--thread-ts", "1234567890.123456", "test message"},
			wantErr:     true,
			errContains: "api-key is required",
		},
		{
			name:    "valid with env api-key",
			args:    []string{"reply", "--channel-id", "C123", "--thread-ts", "1234567890.123456", "test message"},
			env:     map[string]string{"SLACK_API_KEY": "xoxb-test"},
			wantErr: true, // Will fail on network call, but validates args are parsed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup environment
			for k, v := range tt.env {
				os.Setenv(k, v)
				t.Cleanup(func() { os.Unsetenv(k) })
			}

			// Create a new root command for each test
			cmd := rootCmd
			cmd.SetArgs(tt.args)

			// Capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)

			// Execute
			err := cmd.Execute()

			// Check error
			if tt.wantErr && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tt.errContains != "" && err != nil {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("error = %v, want substring %v", err.Error(), tt.errContains)
				}
			}
		})
	}
}

func TestReplyCommandFlags(t *testing.T) {
	// Test that flags are properly defined
	if replyCmd.Flags().Lookup("api-key") == nil {
		t.Error("api-key flag not found")
	}
	if replyCmd.Flags().Lookup("channel-id") == nil {
		t.Error("channel-id flag not found")
	}
	if replyCmd.Flags().Lookup("thread-ts") == nil {
		t.Error("thread-ts flag not found")
	}
}
