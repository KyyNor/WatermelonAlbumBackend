package main

import (
	"WatermelonAlbumBackend/internal/task"
	"WatermelonAlbumBackend/internal/web"
)

func main() {
	go task.InitCron()
	web.Start()
}
