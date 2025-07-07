package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountController struct {

}

func (ac *accountController) HandleCreateAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Account created"})
}
