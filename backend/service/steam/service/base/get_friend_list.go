package base

import (
	"bot_huan/core/component"
	"bot_huan/proto/steam/base"
	"context"
	"errors"
	"github.com/alomerry/copier"
	"github.com/alomerry/steam-web-go-api"
	"github.com/spf13/cast"
)

func (BaseService) GetFriendList(ctx context.Context, req *base.GetFriendListRequest) (*base.GetFriendListResponse, error) {
	if req.SteamId == "" {
		return nil, errors.New("steamId must not empty")
	}

	resp, err := component.Client.GetIUser().GetFriendList(ctx, &steam.GetFriendListRequest{
		SteamId: cast.ToUint64(req.SteamId),
	})
	if err != nil {
		return nil, err
	}
	result := &base.GetFriendListResponse{}
	copier.Instance(nil).From(resp).CopyTo(result)

	return result, nil
}
