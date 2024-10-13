package task

import (
	"github.com/jakecoffman/cron"
)

// CronInstance 创建一个cron实例
var CronInstance = cron.New()

func InitCron() {
	// 每5秒执行一次
	CronInstance.AddFunc(simpleTaskSpec, SimpleTask, "simple_task")

	CronInstance.Start()
}
