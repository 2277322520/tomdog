package utils

import "log"

// Logging 简单地打印日志
// todo 目前不支持格式化输出，日后再优化，暂时调方字符串拼接日志信息即可
func Logging(info string) {
	if GlobalObject.Log {
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println(info)
	}
}
