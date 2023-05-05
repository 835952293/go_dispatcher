package main

import (
	"go_trancation_center/app"
	"go_trancation_center/domain/svc"
)

func main() {
	termbuyingSvc := svc.NewTermBuyingSvc()
	goodsSvc := svc.NewGoodsSvc()
	userSvc := svc.NewUserSvc()

	// 起服务
	app := app.NewManagementApp(termbuyingSvc, goodsSvc, userSvc)
	app.InitRouter().Run()
}
