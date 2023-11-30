package vsm

import (
	"bj38web/service/getCaptcha/imitate/vdb"
	"fmt"
	"math/rand"
	"time"
)

func GenVerifyCode(phone string) (string, error) {
	// 生成 6 位随机数
	code := fmt.Sprintf("%06d", rand.Int31n(1000000))
	vdb.SetEx(phone, code, int64(time.Minute*10))
	return code, nil
}
