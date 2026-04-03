package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell autocompletion",
	Long: `Generate autocompletion script for the specified shell.

To load completions:

Bash:
  $ source <(fluence completion bash)
  # To load completions for each session, execute once:
  # Linux:
  $ fluence completion bash > /etc/bash_completion.d/fluence
  # macOS:
  $ fluence completion bash > $(brew --prefix)/etc/bash_completion.d/fluence

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it. You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  # To load completions for each session, execute once:
  $ fluence completion zsh > "${fpath[1]}/_fluence"
  # You will need to start a new shell for this setup to take effect.

Fish:
  $ fluence completion fish | source
  # To load completions for each session, execute once:
  $ fluence completion fish > ~/.config/fish/completions/fluence.fish

PowerShell:
  PS> fluence completion powershell | Out-String | Invoke-Expression
  # To load completions for every new session, run:
  PS> fluence completion powershell > fluence.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			return cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			return cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
		return nil
	},
}

func init() {
	completionCmd.GroupID = "utility"
	rootCmd.AddCommand(completionCmd)
}
