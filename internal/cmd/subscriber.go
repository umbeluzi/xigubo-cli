// Copyright 2023 Edson Michaque
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/getumbeluzi/plumber-cli/internal"
	"github.com/getumbeluzi/plumber-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdSubscriber(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "subscriber",
		Short:   "Manage subscribers",
		Aliases: []string{"subscribers"},
	}

	cmd.AddCommand(NewCmdSubscriberList(opts))
	cmd.AddCommand(NewCmdSubscriberDelete(opts))
	cmd.AddCommand(NewCmdSubscriberCreate(opts))
	cmd.AddCommand(NewCmdSubscriberGet(opts))

	return cmd
}

func NewCmdSubscriberList(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List subscribers",
		Example: heredoc.Doc(`
			plumber subscriber list
			plumber subscriber list --sandbox
		`),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			internal.SetupIO(cmd, opts)

			cfg, err := config.New()
			if err != nil {
				return err
			}

			cmd.Println(cfg)

			return nil
		},
	}

	return cmd
}

func NewCmdSubscriberDelete(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a subscriber",
		Args:  cobra.NoArgs,
		Example: heredoc.Doc(`
			plumber subscriber delete --subscriber example.com
			plumber subscriber delete --subscriber example.com --sandbox
		`),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			internal.SetupIO(cmd, opts)

			cfg, err := config.New()
			if err != nil {
				return err
			}

			cmd.Println(cfg)

			return nil
		},
	}

	return cmd
}

func NewCmdSubscriberCreate(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a subscriber",
		Example: heredoc.Doc(`
			plumber subscriber new --subscriber example.com
			plumber subscriber new --subscriber example.com --sandbox
		`),
		Args: cobra.MaximumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			internal.SetupIO(cmd, opts)

			cfg, err := config.New()
			if err != nil {
				return err
			}

			cmd.Println(cfg)

			return nil
		},
	}

	return cmd
}

func NewCmdSubscriberGet(opts *internal.CommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Retrieve a subscriber",
		Example: heredoc.Doc(`
			plumber subscriber show --subscriber example.com
			plumber subscriber show --subscriber example.com --sandbox
		`),
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				panic(err)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			internal.SetupIO(cmd, opts)

			cfg, err := config.New()
			if err != nil {
				return err
			}

			cmd.Println(cfg)

			return nil
		},
	}

	return cmd
}
