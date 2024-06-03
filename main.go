package main

import (
	"encoding/json"
	"net/http"

	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

type scoreRecord struct {
	Hiscore int `json:"hiscore"`
	Score   int `json:"score"`
}

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

	Route := api.Http().HttpRouter().PluginRouter().Post("/score/save", func(w http.ResponseWriter, r *http.Request) {
		var data scoreRecord
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		my_key := "score_records"

		var oldData scoreRecord
		err := api.Config().Plugin(my_key).Get(&oldData)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			oldData = scoreRecord{}
		}

		if data.Score >= oldData.Hiscore {
			data.Hiscore = data.Score
		} else {
			data.Hiscore = oldData.Hiscore
		}

		if err := api.Config().Plugin(my_key).Save(data); err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		api.Http().VueResponse().Json(w, nil, http.StatusOK)
	})

	Route.Name("score.save")
}
