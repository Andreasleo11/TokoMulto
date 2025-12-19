package api

import (
	"github.com/Andreasleo11/TokoMulto/internal/config"
	"github.com/Andreasleo11/TokoMulto/internal/module/user"
	"github.com/Andreasleo11/TokoMulto/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.NewMySQL(cfg)

	r := gin.Default()

	// Mount module routes
	user.RegisterRoutes(r, db)
	category.test(r, string)

	r.Run(":" + cfg.AppPort)
}
