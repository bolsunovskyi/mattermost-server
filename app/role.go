// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"github.com/bolsunovskyi/mattermost-server/model"
	"github.com/bolsunovskyi/mattermost-server/utils"
)

func (a *App) Role(id string) *model.Role {
	return a.roles[id]
}

func (a *App) setDefaultRolesBasedOnConfig() {
	a.roles = utils.DefaultRolesBasedOnConfig(a.Config(), a.License() != nil)
}
