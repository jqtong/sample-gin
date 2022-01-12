package conf

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

type logTarget struct {
	FileName string `mapstructure:"file"`
	Level    string `mapstructure:"level"`
}

type serviceFormatter struct {
}

func (f serviceFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(entry.Time.Format("2006-01-02 15:04:05"))

	// append HTTP Status Code
	if httpCode, ok := entry.Data["HTTPStatusCode"]; ok {
		b.WriteString(fmt.Sprintf(" [HTTP %d]", httpCode.(int)))
	}

	// append JSON Resp Code
	if jsonCode, ok := entry.Data["JsonCode"]; ok {
		b.WriteString(fmt.Sprintf(" [JSON %d]", jsonCode.(int)))
	}

	// append latency
	if latency, ok := entry.Data["latency"]; ok {
		b.WriteString(fmt.Sprintf(" [latency %dms]", latency.(int64)))
	}

	// append method
	if method, ok := entry.Data["method"]; ok {
		b.WriteString(" " + method.(string))
	}

	// append endpoint
	if endpoint, ok := entry.Data["endpoint"]; ok {
		b.WriteString(" " + endpoint.(string))
	}

	// append url
	if url, ok := entry.Data["request"]; ok {
		b.WriteString(" " + url.(string))
	}

	// append msg
	if msg, ok := entry.Data["msg"]; ok {
		b.WriteString(" " + msg.(string))
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func initApplicationLog() *logrus.Logger {

	appLog := initLog("app")

	return appLog
}

func initLog(channel string) *logrus.Logger {

	var target logTarget
	err := viper.UnmarshalKey("logs."+channel, &target)
	if err != nil {
		log.Fatalf("init log target %s error: %s", channel, err)
	}

	level, err := logrus.ParseLevel(target.Level)
	if err != nil {
		log.Fatalf("log target %s level invalid: %s", channel, err)
	}

	logFile, err := os.OpenFile(target.FileName,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	if err != nil {
		log.Fatalf("Open access log for target %s file error: %s", channel, err)
	}

	targetLog := logrus.New()
	targetLog.SetOutput(logFile)
	targetLog.SetLevel(level)
	targetLog.SetFormatter(&logrus.TextFormatter{})

	return targetLog
}
