package model

import (
	"bj38web/service/getCaptcha/imitate/vdb"
	"time"
)

// 保存验证码到 数据库/redis
func SaveImgCode(uuid string, code string) error {
	vdb.SetEx(uuid, code, int64(time.Second*60))
	return nil
}

func CheckImgCode(uuid string, code string) bool {
	v := vdb.Get(uuid)
	return v == code
}
