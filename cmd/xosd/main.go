package main

import (
	"fmt"
	"os"

	"github.com/xos-labs/node/cmd/xosd/cmd"
	xosdconfig "github.com/xos-labs/node/cmd/xosd/config"
	appchain "github.com/xos-labs/node/xosd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	setupSDKConfig()

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "xosd", appchain.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}

func setupSDKConfig() {
	config := sdk.GetConfig()
	xosdconfig.SetBech32Prefixes(config)
	config.Seal()
}
