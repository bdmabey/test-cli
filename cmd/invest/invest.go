package invest

import "github.com/spf13/cobra"

type Invest struct {
	cost        float64
	numberOwned int
}

func NewInvestCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "Used to tell the program how to invest",
		Short: "This command tells the program how much a stock costs.",
		Long: `This command tells allows you to tell the program how much a stock costs.
Usage is: test-cli invest <amount>`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
