package task

import (
	"crawler/configs"
	"crawler/global"
	"crawler/internal/spiders"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

func RunTask(cfg configs.Jobs, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task start", time.Now())
	task := job(cfg)
	c := cron.New()
	// 每天的23点执行任务
	var spec string
	if cfg.Model == "full" {
		spec = "0 23 * * *"
	} else {
		spec = "30 9,10,11,14,16,18,20,22 * * *"
	}
	task.Model = cfg.Model
	_, err := c.AddFunc(spec, task.Run)
	if err != nil {
		fmt.Println("添加定时任务出错:", err)
		return
	} else {
		fmt.Println("start task at: ", time.Now(), "model: ", cfg.Model)
	}
	c.Start()
	select {}
}

func job(cfg configs.Jobs) (c spiders.Media) {
	var name string
	switch cfg.Name {
	case "yzzy":
		name = global.Yzzy_Source
	case "ffzy":
		name = global.Ffzy_Source
	default:
		return
	}
	c = spiders.NewClient(name, cfg.Model, cfg.Page)
	return
}
