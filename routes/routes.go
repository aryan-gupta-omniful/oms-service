package routes

import (
	"context"
	"oms-service/controllers"

	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
)

func Initialize(ctx context.Context, s *http.Server) error {
	// Health Check Route
	s.GET("/health", controllers.GetHealth)

	// API v1 Routes Group
	v1 := s.Engine.Group("/api/v1")
	{
		// Hubs Routes
		hubs := v1.Group("/orders")
		{
			hubs.POST("/bulkorder", controllers.CreateBulkOrders)
		}
	}

	log.Infof("Routes initialized successfully")
	return nil
}
