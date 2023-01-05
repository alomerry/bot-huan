package component

import (
	core_utils "bot_huan/core/utils"

	steam "github.com/alomerry/steam-web-go-api"
)

var Client *steam.SteamClient

func init() {
	if Client == nil {
		Client = steam.NewClient(core_utils.GetSteamWebAPI())
	}
	//Client.GetIUserStats().GetUserStatsForGame()
}
