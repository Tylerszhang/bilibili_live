package http

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/k-si/bili_live/config"
	"github.com/k-si/bili_live/entity"
	"log"
	"strconv"
	"time"
)

func GetDanmuInfo() (*entity.ResponseBulletInfo, error) {
	var err error
	var resp *resty.Response
	var url = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=" + strconv.Itoa(config.Live.RoomId) + "&type=0"

	r := &entity.ResponseBulletInfo{}
	if resp, err = cli.R().
		SetHeader("user-agent", userAgent).
		Get(url); err != nil {
		log.Println("请求getDanmuInfo失败：", err)
		return nil, err
	}
	if err = json.Unmarshal(resp.Body(), r); err != nil {
		log.Println("Unmarshal失败：", err, "body:", string(resp.Body()))
		return nil, err
	}

	return r, nil
}

func Send(msg string) error {
	var err error
	var url = "https://api.live.bilibili.com/msg/send"
	var resp *resty.Response

	m := make(map[string]string)
	m["bubble"] = "5"
	m["msg"] = msg
	m["color"] = "4546550"
	m["mode"] = "4"
	m["fontsize"] = "25"
	m["rnd"] = strconv.FormatInt(time.Now().Unix(), 10)
	m["roomid"] = strconv.Itoa(config.Live.RoomId)
	m["csrf"] = CookieList["bili_jct"]
	m["csrf_token"] = CookieList["bili_jct"]

	if resp, err = cli.R().
		SetHeader("user-agent", userAgent).
		SetHeader("cookie", CookieStr).
		SetFormData(m).
		Post(url); err != nil {
		log.Println("请求send失败：", err)
		return err
	}
	log.Println("send 弹幕响应：", string(resp.Body()))

	return nil
}
