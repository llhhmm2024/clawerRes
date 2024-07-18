package bootstrap

import (
	"crawler/configs"
	"crawler/global"
	"crawler/internal/spiders"
	"crawler/pkg/task"
	"sync"
)

var wg sync.WaitGroup

func SetupTask() {
	jobs := configs.Cfg.Jobs
	for _, job := range jobs {
		if job.State == "1" {
			wg.Add(1)
			go runjob(job)
		}
	}
	wg.Wait()
	<-global.MysqlCahn
}

func runjob(cfg configs.Jobs) {
	defer wg.Done()
	var name string
	switch cfg.Name {
	case "yzzy":
		name = global.Yzzy_Source
	case "ffzy":
		name = global.Ffzy_Source
	default:
		return
	}
	c := spiders.NewClient(name, cfg.Model, cfg.Page)
	c.Run()
}

// 启动定时任务
func RunCronTask() {
	jobs := configs.Cfg.Jobs
	for _, t := range jobs {
		if t.State == "1" {
			if t.Model == "all" {
				for _, v := range []string{"full", "incr"} {
					wg.Add(1)
					tmp := t
					tmp.Model = v
					go task.RunTask(tmp, &wg)
				}
			} else {
				wg.Add(1)
				go task.RunTask(t, &wg)
			}
		}
	}
	wg.Wait()
	<-global.MysqlCahn
}

// 指定定时任务
func RunSpecialTask() {
	a := spiders.Media{}
	a.SpecialId()
}

// 更新m3u8
func UpdateM3u8() {
}
