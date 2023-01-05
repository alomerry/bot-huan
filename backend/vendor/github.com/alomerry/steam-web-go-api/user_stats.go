package steam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	INTERFACE_ISTEAM_USER_STATS = "ISteamUserStats"
)

// https://partner.steamgames.com/doc/webapi/ISteamUserStats

type UserStatsInterface _interface

type GetGlobalAchievementPercentagesForAppRequest struct {
	GameId uint64 `json:"gameid"` // 要获取成就百分比的 GameID valid:"require"
}

type GetGlobalStatsForGameRequest struct {
	AppId     uint32 `json:"appid"`               // 要获取全局统计的 AppID
	Count     uint32 `json:"count"`               // 要获取数据的统计的数量
	Name0     string `json:"name[0]"`             // 要获取数据的统计的名称
	StartDate uint32 `json:"startdate,omitempty"` // 每日合计的开始日期（unix 时间戳）
	EndDate   uint32 `json:"enddate,omitempty"`   // 每日合计的结束日期（unix 时间戳）
}

type GetUserStatsForGameRequest struct {
	Key     string `json:"key"`     // Steamworks Web API 用户验证密钥
	SteamId uint64 `json:"steamid"` // 用户的 SteamID
	AppId   uint32 `json:"appid"`   // 游戏的 appid
}

type GetUserStatsForGameResponse struct {
	PlayerStats PlayerStats `json:"playerstats"`
}

type PlayerStats struct {
	SteamId  string  `json:"steamID"`
	GameName string  `json:"gameName"`
	Stats    []Stats `json:"stats"`
}

type Stats struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// GetGlobalAchievementPercentagesForApp 为指定应用获取全局成就百分比
func (UserStatsInterface) GetGlobalAchievementPercentagesForApp() {
	// GET https://api.steampowered.com/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v2/
}

// GetGlobalStatsForGame 为指定应用获取全局统计数据百分比
func (UserStatsInterface) GetGlobalStatsForGame(ctx context.Context, req *GetGlobalStatsForGameRequest) {
	// GET https://api.steampowered.com/ISteamUserStats/GetGlobalStatsForGame/v1/
}

// GetUserStatsForGame 获得指定用户在一个应用中设置的统计列表
func (u *UserStatsInterface) GetUserStatsForGame(ctx context.Context, req *GetUserStatsForGameRequest) (*GetUserStatsForGameResponse, error) {
	// GET https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/
	resp := &GetUserStatsForGameResponse{}
	req.Key = u.sc.key
	endpoint := getEndPoint(INTERFACE_ISTEAM_USER_STATS, 2)
	endpoint = appendQuery(endpoint, req)
	err := (*_interface)(u).execute(ctx, endpoint, resp)
	return resp, err
}

func appendQuery(endpoint string, req interface{}) string {
	if reflect.ValueOf(req).IsZero() {
		return ""
	}

	mapper := make(map[string]interface{})
	data, _ := json.Marshal(req)
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	_ = d.Decode(&mapper)
	index := 0
	for k, v := range mapper {
		if index == 0 {
			endpoint = fmt.Sprintf("%s?%s=%v", endpoint, k, v)
		} else {
			endpoint = fmt.Sprintf("%s&%s=%v", endpoint, k, v)
		}
		index++
	}
	return endpoint
}
