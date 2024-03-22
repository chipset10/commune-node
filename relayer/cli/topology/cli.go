// The Licensed Work is (c) 2024 Commune
// SPDX-License-Identifier: LGPL-3.0-only

package topology

import "github.com/spf13/cobra"

var TopologyCLI = &cobra.Command{
	Use:   "topology",
	Short: "utility commands that helps to encrypt and test p2p TopologyMap",
}

func init() {
	TopologyCLI.AddCommand(encryptTopologyCMD)
	TopologyCLI.AddCommand(testTopologyCMD)
}
