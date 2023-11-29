package model

import "bj38web/service/getCaptcha/vdb"

// 保存验证码到 数据库/redis
func SaveImgCode(uuid string, code string) error {
	vdb.SetEx(uuid, code, 60)
	return nil
}
