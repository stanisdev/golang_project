package middlewares

import (
	"app/models"
	"app/services"
	"github.com/gin-gonic/gin"
)

func FindNotificationById(c *gin.Context) {
	id := c.MustGet("id").(uint)
	ntf := models.GetDmInstance().FindNotificationById(id)

	if (ntf.ID < 1) { // Notification not found
		services.WrongUrlParams(c)
		c.Abort()
		return
	}
	c.Set("notification", ntf)
	c.Set("oldCompanyID", ntf.CompanyID)
	c.Next()
}