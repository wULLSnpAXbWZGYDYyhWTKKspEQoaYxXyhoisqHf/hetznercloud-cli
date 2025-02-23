package firewall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/hetznercloud/hcloud-go/hcloud/schema"
	"github.com/spf13/cobra"

	"github.com/hetznercloud/cli/internal/cmd/cmpl"
	"github.com/hetznercloud/cli/internal/cmd/util"
	"github.com/hetznercloud/cli/internal/state"
)

func newReplaceRulesCommand(cli *state.State) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "replace-rules FIREWALL FLAGS",
		Short:                 "Replaces all rules from a Firewall from a file",
		Args:                  cobra.ExactArgs(1),
		ValidArgsFunction:     cmpl.SuggestArgs(cmpl.SuggestCandidatesF(cli.FirewallNames)),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               util.ChainRunE(cli.EnsureToken),
		RunE:                  cli.Wrap(runFirewallReplaceRules),
	}
	cmd.Flags().String("rules-file", "", "JSON file containing your routes (use - to read from stdin). The structure of the file needs to be the same as within the API: https://docs.hetzner.cloud/#firewalls-get-a-firewall")
	cmd.MarkFlagRequired("rules-file")
	return cmd
}

func runFirewallReplaceRules(cli *state.State, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	firewall, _, err := cli.Client().Firewall.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}
	if firewall == nil {
		return fmt.Errorf("Firewall not found: %v", idOrName)
	}

	opts := hcloud.FirewallSetRulesOpts{}

	rulesFile, _ := cmd.Flags().GetString("rules-file")

	var data []byte
	if rulesFile == "-" {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = ioutil.ReadFile(rulesFile)
	}
	if err != nil {
		return err
	}
	var rules []schema.FirewallRule
	err = json.Unmarshal(data, &rules)
	if err != nil {
		return err
	}
	for _, rule := range rules {
		d := hcloud.FirewallRuleDirection(rule.Direction)
		r := hcloud.FirewallRule{
			Direction:   d,
			Protocol:    hcloud.FirewallRuleProtocol(rule.Protocol),
			Port:        rule.Port,
			Description: rule.Description,
		}
		switch d {
		case hcloud.FirewallRuleDirectionOut:
			r.DestinationIPs = make([]net.IPNet, len(rule.DestinationIPs))
			for i, ip := range rule.DestinationIPs {
				_, n, err := net.ParseCIDR(ip)
				if err != nil {
					return fmt.Errorf("invalid CIDR on index %d : %s", i, err)
				}
				r.DestinationIPs[i] = *n
			}
		case hcloud.FirewallRuleDirectionIn:
			r.SourceIPs = make([]net.IPNet, len(rule.SourceIPs))
			for i, ip := range rule.SourceIPs {
				_, n, err := net.ParseCIDR(ip)
				if err != nil {
					return fmt.Errorf("invalid CIDR on index %d : %s", i, err)
				}
				r.SourceIPs[i] = *n
			}
		}
		opts.Rules = append(opts.Rules, r)
	}

	actions, _, err := cli.Client().Firewall.SetRules(cli.Context, firewall, opts)
	if err != nil {
		return err
	}
	if err := cli.ActionsProgresses(cli.Context, actions); err != nil {
		return err
	}
	fmt.Printf("Firewall Rules for Firewall %d updated\n", firewall.ID)

	return nil
}
