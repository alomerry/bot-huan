package steam

import (
	"context"
)

const (
	INTERFACE_ISTEAM_USER = "ISteamUser"

	RELATIONSHIP_FRIEND = "friend"
)

// https://partner.steamgames.com/doc/webapi/ISteamUser

type UserInterface _interface

type GetFriendListRequest struct {
	Key     string `json:"key"`
	SteamId uint64 `json:"steamid"`
}

type GetFriendListResponse struct {
	FriendsList FriendsList `json:"friendslist"`
}

type FriendsList struct {
	Friends []RelatedPeople `json:"friends"`
}

type RelatedPeople struct {
	SteamId      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  uint64 `json:"friend_since"`
}

// GetFriendList 为指定应用获取全局成就百分比
func (u *UserInterface) GetFriendList(ctx context.Context, req *GetFriendListRequest) (*GetFriendListResponse, error) {
	// GET https://api.steampowered.com/ISteamUser/GetFriendList/v1/
	resp := &GetFriendListResponse{}
	req.Key = u.sc.key
	endpoint := getEndPoint(INTERFACE_ISTEAM_USER, 1)
	endpoint = appendQuery(endpoint, req)
	err := (*_interface)(u).execute(ctx, endpoint, resp)
	return resp, err
}
