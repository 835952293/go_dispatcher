package app

import (
	"go_trancation_center/domain/svc"

	"github.com/gin-gonic/gin"
)

type managementApp struct {
	termbuyingSvc svc.TermBuyingSvc
	goodsSvc      svc.GoodsSvc
	userSvc       svc.UserSvc
}

func NewManagementApp(
	termbuyingSvc svc.TermBuyingSvc,
	goodsSvc svc.GoodsSvc,
	userSvc svc.UserSvc) *managementApp {
	a := &managementApp{termbuyingSvc: termbuyingSvc, goodsSvc: goodsSvc, userSvc: userSvc}
	return a
}

func (a *managementApp) InitRouter() *gin.Engine {
	app := gin.New()

	// 暂时先不研究导入导出
	// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	v1 := app.Group("/v1/api")
	{
		// 拼单路由
		termbuyingRouter := v1.Group("/termbuying")
		{
			termbuyingRouter.GET("termbuying")
		}

		// 商品路由
		goodsRouter := v1.Group("/goods")
		{
			goodsRouter.GET("")
		}

		// 用户路由
		userRouter := v1.Group("user")
		{
			userRouter.GET("")
		}
	}

	return app
}
