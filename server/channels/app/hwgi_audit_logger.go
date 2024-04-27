package app

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	// log파일 존재여부 확인
	hwgi_checkAuditLogFile()
	// logrus 설정
	nowDate := time.Now().Format("20060102")
	var filename = "./logs/auditlogs-" + nowDate + ".log"

	// 로그 포맷 설정
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// lumberjack 설정
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100, // MB
		MaxBackups: 7,
		MaxAge:     7, // days
		Compress:   true,
	}

	// logrus 출력 설정
	logger.SetOutput(lumberjackLogger)

}

func hwgi_Info(msg string) {
	logger.Info(msg)
}

func hwgi_Error(msg string) {
	logger.Error(msg)
}

func hwgi_checkAuditLogFile() error {
	nowDate := time.Now().Format("20060102")
	var filename = "./logs/auditlogs-" + nowDate + ".log"

	// 파일이 존재하는지 확인
	_, err := os.Stat(filename)

	// 파일이 존재하지 않으면 생성
	if os.IsNotExist(err) {
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
