package controller

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pwhb/lear/pkg/model"
)

func RenderHTML(ctx *gin.Context) {
	id := rand.Intn(111) + 1
	poem, err := model.GetOnePoem(id)
	if err != nil {
		log.Println(err)
	}
	ctx.HTML(http.StatusOK, "index.html",
		poem)
}

func GetAllPoems(ctx *gin.Context) {
	poemList, err := model.GetAllPoems()
	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, poemList)
}

func GetOnePoem(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, err)
	}
	res, err := model.GetOnePoem(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func GetRandomPoem(ctx *gin.Context) {
	id := rand.Intn(111) + 1
	res, err := model.GetOnePoem(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, res)
}
