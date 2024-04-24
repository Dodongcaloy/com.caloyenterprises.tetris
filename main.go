package main

import (
	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
	// define the portal route

	portalRoute := sdkhttp.VuePortalRoute{
		RouteName: "portal.tetris",
		RoutePath: "/tetris",
		Component: "portal/tetris.vue",
	}

	// register portal route
	api.Http().VueRouter().RegisterPortalRoutes(portalRoute)

	api.Http().VueRouter().PortalItemsFunc(func(clnt sdkconnmgr.ClientDevice) []sdkhttp.VuePortalItem {
		portalItem := sdkhttp.VuePortalItem{
			IconPath:  "icons/tetris.png",
			Label:     "TETRIS",
			RouteName: "portal.tetris",
			// RouteParams: map[string]string{},
		}
		return []sdkhttp.VuePortalItem{portalItem}
	})

}
