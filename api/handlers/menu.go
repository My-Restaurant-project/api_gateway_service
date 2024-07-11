package handler

import (
	rese "api_gateway/genproto/reservation_service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

//@
func (r *reservationHandlerImpl) AddMenu(ctx *gin.Context) {
	var addMenuReq rese.AddMenuRequest
	err := ctx.BindJSON(&addMenuReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	addMenuRes, err := r.reservationService.AddMenu(context.TODO(), &addMenuReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Menu added successfully", "menu": addMenuRes})
}

func (r *reservationHandlerImpl) GetMenus(ctx *gin.Context) {
	getMenuReq := rese.GetMenusRequest{}

	getMenuRes, err := r.reservationService.GetMenus(context.TODO(), &getMenuReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Menus retrieved successfully", "menus": getMenuRes.Menus})
}

func (r *reservationHandlerImpl) GetMenu(ctx *gin.Context) {
	id := ctx.Param("id")

	getMenuReq := rese.GetMenuRequest{Id: id}

	getMenuRes, err := r.reservationService.GetMenu(context.TODO(), &getMenuReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Menu retrieved successfully", "menu": getMenuRes.Menu})
}

func (r *reservationHandlerImpl) UpdateMenu(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateMenuReq rese.UpdateMenuRequest
	updateMenuReq.Id = id

	err := ctx.BindJSON(&updateMenuReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	updateMenuReq.Id = id

	updateMenuRes, err := r.reservationService.UpdateMenu(context.TODO(), &updateMenuReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully", "menu": updateMenuRes.Menu})
}

func (r *reservationHandlerImpl) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")

	deleteMenuReq := rese.DeleteMenuRequest{Id: id}

	_, err := r.reservationService.DeleteMenu(context.TODO(), &deleteMenuReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
