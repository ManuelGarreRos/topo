package handlers

import (
	"TOPO/common"
	"TOPO/internal/models"
	"TOPO/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func NewUserHandler(us *services.UserService, log *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: us,
		log:         log,
	}
}

type UserHandler struct {
	userService *services.UserService
	log         *zap.Logger
}

func (uh UserHandler) ByID(ctx *gin.Context) {
	id := ctx.Param("id")

	u, err := uh.userService.ByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{"Error": "Error al obtener usuario por ID"}) //TODO add jsonError parsing library?
	}
	ctx.JSON(200, u)
}

func (uh UserHandler) PaginatedList(ctx *gin.Context) {
	q, err := uh.parseQuery(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"Error": "Error parseando query"})
	}

	ul, err := uh.userService.PaginatedList(q)
	if err != nil {
		ctx.JSON(500, gin.H{"Error": "Pues alguno que venga..."})
	}
	listUser := models.UserList{
		ul,
		int64(len(ul)),
	}
	ctx.JSON(200, listUser)
}

func (uh UserHandler) Create(ctx *gin.Context) {
	/*q, err := uh.parseQuery(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"Error":"Error parseando query"})
	}*/

	u := &models.User{}
	if err := ctx.ShouldBindJSON(u); err != nil {
		ctx.JSON(500, gin.H{"Error": "Error parseando JSON"})
	}

	if err := uh.userService.Create(u); err != nil {
		ctx.JSON(500, gin.H{"Error": "Yo que se"})
	}

	ctx.JSON(http.StatusOK, "Json recibido")
}

func (uh UserHandler) Delete(ctx *gin.Context) {
	/*
	 IMPORTANT => Dejamos de usar id en la URL y pasamos a usarlo en el body
	*/
	u := &models.User{}
	if err := ctx.ShouldBindJSON(u); err != nil {
		ctx.JSON(500, gin.H{"Error": "Error parseando JSON"})
	}

	if err := uh.userService.Delete(u); err != nil {
		ctx.JSON(500, gin.H{"Error": "Error deleting user"})
	}

	ctx.JSON(http.StatusOK, "User deleted")
}

func (uh UserHandler) Update(ctx *gin.Context) {
	u := &models.User{}
	if err := ctx.ShouldBindJSON(u); err != nil {
		ctx.JSON(500, gin.H{"Error": "Error parseando JSON"})
	}

	if err := uh.userService.Update(u); err != nil {
		ctx.JSON(500, gin.H{"Error": "Error updating user"})
	}

	ctx.JSON(http.StatusOK, "User updated")
}

func (uh UserHandler) parseQuery(ctx *gin.Context) (*models.UserQuery, error) {
	q := &models.UserQuery{}

	if err := q.ParseBase(ctx); err != nil {
		return nil, err
	}

	if err := common.ParseQuery(ctx, q); err != nil {
		return nil, err
	}

	if len(q.Sorts) == 0 {
		q.Sorts = append(q.Sorts, common.Sort{
			Field: models.UserName,
		})
	}

	return q, nil
}
