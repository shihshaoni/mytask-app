package delivery

import (
	"mytask-app/model"
	"mytask-app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MyTaskHandler struct {
	Repo *repository.MyTaskRepository
}

func NewMyTaskHandler(r *gin.Engine, repo *repository.MyTaskRepository) {
	handler := &MyTaskHandler{Repo: repo}

	r.POST("/my-tasks", handler.CreateMyTask)
	r.GET("/my-tasks/:id", handler.GetMyTask)
	r.PUT("/my-tasks/:id", handler.UpdateMyTask)
	r.DELETE("/my-tasks/:id", handler.DeleteMyTask)
}

func (h *MyTaskHandler) CreateMyTask(c *gin.Context) {
	var MyTask model.MyTask
	if err := c.ShouldBindJSON(&MyTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdMyTask, err := h.Repo.Create(&MyTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdMyTask)
}

func (h *MyTaskHandler) GetMyTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	myTask, err := h.Repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "MyTask not found!"})
		return
	}
	c.JSON(http.StatusOK, myTask)
}

func (h *MyTaskHandler) UpdateMyTask(c *gin.Context) {
	var MyTask model.MyTask
	id, _ := strconv.Atoi(c.Param("id"))
	MyTask.ID = uint(id)
	if err := c.ShouldBindJSON(&MyTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedMyTask, err := h.Repo.Update(&MyTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedMyTask)
}

func (h *MyTaskHandler) DeleteMyTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var MyTask model.MyTask
	MyTask.ID = uint(id)
	if err := h.Repo.Delete(&MyTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "MyTask deleted"})
}
