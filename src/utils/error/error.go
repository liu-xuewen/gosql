package error

import "utils/logger"

func CheckErr(err error) {
	if err != nil {
		logger.Error("err:", err)
		return
	}
}
