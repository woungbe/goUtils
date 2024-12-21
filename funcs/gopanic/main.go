package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	// pawlog 대신 logrus 사용
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	defer func() {
		if err := recover(); err != nil {
			errMsg := fmt.Sprintf("%v", err)
			logrus.WithFields(logrus.Fields{
				"package": "userFuncs",
				"func":    "DBInsertMultiKYCUserForCP",
			}).Error("Crit Error !!!!! Panic userFuncs", errMsg)
		}
	}()

	// 일부러 패닉을 발생시킴
	panic("something went wrong")
}
