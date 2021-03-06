package monocle

import (
	"strings"

	"github.com/spf13/cobra"
)

var DefaultMonocle = &Monocle{}

// Enable enables monocle on a cobra command using the default monocle
func Enable(c *cobra.Command) {
	DefaultMonocle.Enable(c)
}

// Primary sets the commands a primary help topic
func Primary(cmds ...*cobra.Command) {
	DefaultMonocle.Primary(cmds...)
}

// Monocle represents the component that enables custom help on a cobra command
type Monocle struct {
	*cobra.Command

	primaries []*cobra.Command
}

// New returns a a new Monocle
func New() *Monocle {
	return &Monocle{}
}

// Enable enables monocle on a cobra command
func (m *Monocle) Enable(c *cobra.Command) {
	m.Command = c
	c.SetUsageFunc(m.UsageFunc())
}

// UsageFunc returns the usage function that'll be used for the command
func (m *Monocle) UsageFunc() (f func(*cobra.Command) error) {
	// if m.Command.HasParent() {
	// 	return m.Command.Parent().UsageFunc()
	// }
	return func(c *cobra.Command) error {
		if c == m.Command {
			return tmpl(c.OutOrStderr(), topicTemplate, m)
		}
		return tmpl(c.OutOrStderr(), usageTemplate, c)
	}
}

// Primary sets a command a primary help topic
func (m *Monocle) Primary(cmds ...*cobra.Command) {
	m.primaries = append(m.primaries, cmds...)
}

// PrimaryCommands returns a list primary commands enabled using the Primary function
func (m *Monocle) PrimaryCommands() []*cobra.Command {
	return m.primaries
}

func (m *Monocle) IsCommandTopic(cmd *cobra.Command) bool {
	return (cmd.HasSubCommands() || cmd.IsAvailableCommand())
}

// AdditionalCommands returns the a list addtional commands when a primary command is specified. Returns nil otherwise
func (m *Monocle) AdditionalCommands() []*cobra.Command {
	hasPrimary := len(m.primaries) > 0
	if hasPrimary {
		res := make([]*cobra.Command, 0)
		for _, cmd := range m.Command.Commands() {
			include := true
			for _, primaryCmd := range m.primaries {
				switch {
				case cmd == primaryCmd:
					include = false
					break
					// there is perhaps an elegant way to check to namespaced commands
				case strings.Contains(cmd.Name(), ":"):
					include = false
					break
				}
			}
			if include {
				res = append(res, cmd)
			}
		}
		return res
	}
	return make([]*cobra.Command, 0)
}
