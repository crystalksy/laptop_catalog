package route

import (
	"laptop_catalog/constants"
	"laptop_catalog/controllers"
	"laptop_catalog/middleware"

	m "laptop_catalog/middleware"

	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/admins", controllers.GetAdminsController)
	e.GET("/admins/:id", controllers.GetAdminByID)
	m.LogMiddleware(e)

	e.POST("/admins", controllers.CreateAdminController)
	e.DELETE("/admins/:id", controllers.DeleteAdminByID)
	e.PUT("/admins/:id", controllers.UpdateAdminByID)
	e.POST("/login", controllers.LoginAdminController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(midd.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/admins", controllers.GetAdminsController)

	eJwt := e.Group("/jwt")
	eJwt.Use(midd.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/admins", controllers.GetAdminsController)

	e.GET("/batteries", controllers.GetBatteries)
	e.GET("/batteries/:id", controllers.GetBatteryByID)
	e.POST("/batteries", controllers.CreateBattery)
	e.DELETE("/batteries/:id", controllers.DeleteBatteryByID)
	e.PUT("/batteries/:id", controllers.UpdateBatteryByID)

	e.GET("/brands", controllers.GetBrands)
	e.GET("/brands/:id", controllers.GetBrandByID)
	e.POST("/brands", controllers.CreateBrand)
	e.DELETE("/brands/:id", controllers.DeleteBrandByID)
	e.PUT("/brands/:id", controllers.UpdateBrandByID)

	e.GET("/types", controllers.GetTypes)
	e.GET("/types/:id", controllers.GetTypeByID)
	e.POST("/types", controllers.CreateType)
	e.DELETE("/types/:id", controllers.DeleteTypeByID)
	e.PUT("/types/:id", controllers.UpdateTypeByID)

	e.GET("/connectors", controllers.GetConnectors)
	e.GET("/connectors/:id", controllers.GetConnectorByID)
	e.POST("/connectors", controllers.CreateConnector)
	e.DELETE("/connectors/:id", controllers.DeleteConnectorByID)
	e.PUT("/connectors/:id", controllers.UpdateConnectorByID)

	e.GET("/memories", controllers.GetMemories)
	e.GET("/memories/:id", controllers.GetMemoryByID)
	e.POST("/memories", controllers.CreateMemory)
	e.DELETE("/memories/:id", controllers.DeleteMemoryByID)
	e.PUT("/memories/:id", controllers.UpdateMemoryByID)

	e.GET("/designs", controllers.GetDesigns)
	e.GET("/designs/:id", controllers.GetDesignByID)
	e.POST("/designs", controllers.CreateDesign)
	e.DELETE("/designs/:id", controllers.DeleteDesignByID)
	e.PUT("/designs/:id", controllers.UpdateDesignByID)

	e.GET("/descs", controllers.GetDescs)
	e.GET("/descs/:id", controllers.GetDescByID)
	e.POST("/descs", controllers.CreateDesc)
	e.DELETE("/descs/:id", controllers.DeleteDescByID)
	e.PUT("/descs/:id", controllers.UpdateDescByID)

	e.GET("/cameras", controllers.GetCameras)
	e.GET("/cameras/:id", controllers.GetCameraByID)
	e.POST("/cameras", controllers.CreateCamera)
	e.DELETE("/cameras/:id", controllers.DeleteCameraByID)
	e.PUT("/cameras/:id", controllers.UpdateCameraByID)

	e.GET("/displays", controllers.GetDisplays)
	e.GET("/displays/:id", controllers.GetDisplayByID)
	e.POST("/displays", controllers.CreateDisplay)
	e.DELETE("/displays/:id", controllers.DeleteDisplayByID)
	e.PUT("/displays/:id", controllers.UpdateDisplayByID)

	e.GET("/storages", controllers.GetStorages)
	e.GET("/storages/:id", controllers.GetStorageByID)
	e.POST("/storages", controllers.CreateStorage)
	e.DELETE("/storages/:id", controllers.DeleteStorageByID)
	e.PUT("/storages/:id", controllers.UpdateStorageByID)

	e.GET("/features", controllers.GetFeatures)
	e.GET("/features/:id", controllers.GetFeatureByID)
	e.POST("/features", controllers.CreateFeature)
	e.DELETE("/features/:id", controllers.DeleteFeatureByID)
	e.PUT("/features/:id", controllers.UpdateFeatureByID)

	e.GET("/processors", controllers.GetProcessors)
	e.GET("/processors/:id", controllers.GetProcessorByID)
	e.POST("/processors", controllers.CreateProcessor)
	e.DELETE("/processors/:id", controllers.DeleteProcessorByID)
	e.PUT("/processors/:id", controllers.UpdateProcessorByID)

	return e
}
