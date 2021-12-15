/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package frp

import (
	"github.com/webx-top/echo"

	"github.com/admpub/nging/v4/application/handler"
	"github.com/admpub/nging/v4/application/registry/perm"

	"github.com/nging-plugins/frpmanager/pkg/dbschema"
	_ "github.com/nging-plugins/frpmanager/pkg/handler/plugins/multiuser"
	"github.com/nging-plugins/frpmanager/pkg/handler/proxy"
)

func init() {
	handler.RegisterToGroup(`/frp`, func(g echo.RouteRegister) {
		metaHandler := handler.IRegister().MetaHandler
		g.Route(`GET,POST`, `/server_index`, ServerIndex)
		g.Route(`GET,POST`, `/server_add`, ServerAdd)
		g.Route(`GET,POST`, `/server_edit`, ServerEdit)
		g.Route(`GET,POST`, `/server_delete`, ServerDelete)
		g.Route(`GET,POST`, `/server_log`, ServerLog)

		g.Route(`GET`, `/account`, AccountIndex)
		g.Route(`GET,POST`, `/account_add`, AccountAdd)
		g.Route(`GET,POST`, `/account_edit`, AccountEdit)
		g.Route(`GET,POST`, `/account_delete`, AccountDelete)

		g.Route(`GET`, `/client_index`, ClientIndex)
		g.Route(`GET,POST`, `/client_add`, ClientAdd)
		g.Route(`GET,POST`, `/client_edit`, ClientEdit)
		g.Route(`GET,POST`, `/client_delete`, ClientDelete)
		g.Route(`GET,POST`, `/client_log`, ClientLog)

		g.Route(`GET`, `/group_index`, GroupIndex)
		g.Route(`GET,POST`, `/group_add`, GroupAdd)
		g.Route(`GET,POST`, `/group_edit`, GroupEdit)
		g.Route(`GET,POST`, `/group_delete`, GroupDelete)
		g.Route(`GET,POST`, `/server_restart`, ServerRestart)
		g.Route(`GET,POST`, `/server_stop`, ServerStop)
		g.Route(`GET,POST`, `/client_restart`, ClientRestart)
		g.Route(`GET,POST`, `/client_stop`, ClientStop)
		g.Route(`GET`, `/addon_form`, metaHandler(echo.H{`name`: `FRP客户端配置表单`}, AddonForm))
	})
	handler.RegisterToGroup(`/frp/dashboard`, func(g echo.RouteRegister) {
		g.Get(`/server/:id`, ServerDashboard)
		g.Get(`/client/:id`, ClientDashboard)

		// - 代理方案 -
		g.Route(`GET,POST`, `/server/:id/*`, echo.NotFoundHandler, proxy.ProxyServer, proxy.Proxy())
		g.Route(`GET,POST`, `/client/:id/*`, echo.NotFoundHandler, proxy.ProxyClient, proxy.Proxy())
	})
	perm.AuthRegister(`/frp/dashboard/server/:id`, authDashboard)
	perm.AuthRegister(`/frp/dashboard/client/:id`, authDashboard)
	perm.AuthRegister(`/frp/dashboard/server/:id/*`, authDashboard)
	perm.AuthRegister(`/frp/dashboard/client/:id/*`, authDashboard)
}

func authDashboard(
	h echo.Handler,
	c echo.Context,
	user *dbschema.NgingUser,
	permission *perm.RolePermission,
) (ppath string, returning bool, err error) {
	ppath = `/frp/dashboard`
	return
}
