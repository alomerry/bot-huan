package component

import (
    "bot_huan/core/log"
    "encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	NeteaseAPI Netease
)

type Netease struct {
}

func (Netease) Search(key string, page uint8) (int32, string, string) {
	offset := strconv.FormatInt(int64(page-1), 10)
	resp, err := http.Get("https://netease.alomerry.com/search?keywords=" + key + "&limit=1&offset=" + offset)
	if err != nil {
		log.GetLogger().Err(err)
		return -1, "", ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.GetLogger().Err(err)
		return -1, "", ""
	}

	response := struct {
		Code   uint8 `json:"code"`
		Result struct {
			Songs []struct {
				Name string `json:"name"`
				Id   int32  `json:"id"`
				Al   struct {
					PicUrl string `json:"picUrl"`
				} `json:"al"`
			}
		} `json:"result"`
		Mention []json.RawMessage `json:"mention"`
	}{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.GetLogger().Err(err)
		return -1, "", ""
	}
	if response.Code != http.StatusOK || len(response.Result.Songs) < 1 {
		return -1, "", ""
	}
	return response.Result.Songs[0].Id, response.Result.Songs[0].Al.PicUrl, response.Result.Songs[0].Name
}

func (Netease) GetSongUrl(id int32) string {
	if id <= 0 {
		return ""
	}
	resp, err := http.Get("https://netease.alomerry.com/song/url?id=" + strconv.FormatInt(int64(id), 10))
	if err != nil {
		log.GetLogger().Err(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.GetLogger().Err(err)
		return ""
	}
	response := struct {
		Code uint8 `json:"code"`
		Data []struct {
			Id  int32  `json:"id"`
			Url string `json:"url"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.GetLogger().Err(err)
		return ""
	}
	if response.Code != http.StatusOK || len(response.Data) < 1 {
		return ""
	}
	return response.Data[0].Url
}
