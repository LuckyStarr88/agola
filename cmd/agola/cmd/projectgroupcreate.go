// Copyright 2019 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sorintlab/agola/internal/services/gateway/api"

	"github.com/spf13/cobra"
)

var cmdProjectGroupCreate = &cobra.Command{
	Use:   "create",
	Short: "create a project",
	Run: func(cmd *cobra.Command, args []string) {
		if err := projectGroupCreate(cmd, args); err != nil {
			log.Fatalf("err: %v", err)
		}
	},
}

type projectGroupCreateOptions struct {
	name       string
	parentPath string
}

var projectGroupCreateOpts projectGroupCreateOptions

func init() {
	flags := cmdProjectGroupCreate.Flags()

	flags.StringVarP(&projectGroupCreateOpts.name, "name", "n", "", "project name")
	flags.StringVar(&projectGroupCreateOpts.parentPath, "parent", "", `parent project group path (i.e "org/org01" for root project group in org01, "/user/user01/group01/subgroub01") or project group id where the project group should be created`)

	cmdProjectGroupCreate.MarkFlagRequired("name")
	cmdProjectGroupCreate.MarkFlagRequired("parent")

	cmdProjectGroup.AddCommand(cmdProjectGroupCreate)
}

func projectGroupCreate(cmd *cobra.Command, args []string) error {
	gwclient := api.NewClient(gatewayURL, token)

	req := &api.CreateProjectGroupRequest{
		Name:     projectGroupCreateOpts.name,
		ParentID: projectGroupCreateOpts.parentPath,
	}

	log.Infof("creating project group")

	project, _, err := gwclient.CreateProjectGroup(context.TODO(), req)
	if err != nil {
		return errors.Wrapf(err, "failed to create project")
	}
	log.Infof("project %s created, ID: %s", project.Name, project.ID)

	return nil
}