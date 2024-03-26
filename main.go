package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
	// define the portal route
	api.Http().VueRouter().PortalItemsFunc(func(clnt sdkconnmgr.ClientDevice) []sdkhttp.VuePortalItem {
		portalItem := sdkhttp.VuePortalItem{
			Label:     "welcome",
			RouteName: "portal.welcome",
			//RouteParams: map[string]string{"name": "CALOY"},
		}
		return []sdkhttp.VuePortalItem{portalItem}
	})

	// define the portal route

	portalRoute := sdkhttp.VuePortalRoute{
		RouteName: "portal.welcome",
		RoutePath: "/welcome",
		Component: "portal/welcome.vue",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			// params := api.Http().MuxVars(r)
			name := "CALOY"

			data := map[string]interface{}{
				"name": name,
			}

			api.Http().VueResponse().Json(w, data, 200)
		},
		Middlewares: []func(http.Handler) http.Handler{},
	}
	// register portal route
	api.Http().VueRouter().RegisterPortalRoutes(portalRoute)

	api.Http().HttpRouter().PluginRouter().Post("/payments/receieved", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Amount int `json:"amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Payment Received: %d", data.Amount)

	}).Name("payment.received")
}
