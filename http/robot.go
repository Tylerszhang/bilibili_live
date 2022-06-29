package http

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/k-si/bili_live/entity"
	"log"
)

// 调用青云客机器人api
func RequestQingyunkeRobot(msg string) (string, error) {
	var err error
	var url = "http://api.qingyunke.com/api.php?key=free&appid=0&msg=" + msg
	var resp *resty.Response

	if resp, err = cli.R().
		SetHeader("Content-Type", "utf-8").
		Get(url); err != nil {
		log.Println("请求qingyunke机器人接口失败：", err)
		return "", err
	}

	r := &entity.QinugyunkeRobotReplay{}
	err = json.Unmarshal(resp.Body(), r)

	return r.Content, err
}
