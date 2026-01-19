package handlers

import (
	"log"
	"net/http"
	m "task/models"

	"github.com/gin-gonic/gin"
)

func (h *HandlerModule) CreateNumberHandler(c *gin.Context) {
	var num m.Numbers
	err := c.BindJSON(&num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect data",
		})
		return
	}
	err = h.repo.SaveNumber(num.Number)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	log.Println("Число успешно сохранено", num.Number)
	nums, err := h.repo.GetNumbers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	sortNums := m.ValidateNumbers(nums)
	c.JSON(http.StatusOK, gin.H{
		"success": sortNums,
	})
}
