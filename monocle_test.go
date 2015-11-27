package monocle

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

var runFunc = func(cmd *cobra.Command, args []string) { fmt.Println("run", cmd.Name()) }

func TestPrimaryAdditional(t *testing.T) {
	root := &cobra.Command{Use: "root"}
	cmd1 := &cobra.Command{Use: "cmd1"}
	cmd1.AddCommand(&cobra.Command{Use: "cmd1.1", Run: runFunc})
	root.AddCommand(cmd1)
	cmd2 := &cobra.Command{Use: "cmd2"}
	cmd2.AddCommand(&cobra.Command{Use: "cmd2.1", Run: runFunc})
	root.AddCommand(cmd2)
	cmd3 := &cobra.Command{Use: "cmd3"}
	cmd3.AddCommand(&cobra.Command{Use: "cmd3.1", Run: runFunc})
	root.AddCommand(cmd3)

	m := New()
	m.Enable(root)
	m.Primary(cmd1, cmd2)

	if !reflect.DeepEqual(m.PrimaryCommands(), []*cobra.Command{cmd1, cmd2}) {
		t.Fatalf("expected primary commands to have cmd1")
	}

	if !reflect.DeepEqual(m.AdditionalCommands(), []*cobra.Command{cmd3}) {
		t.Fatalf("expected additonal commands to be cmd3, got ", m.AdditionalCommands())
	}
}
