package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"server/model"
)

//NupsCtr is a structure to manipulate database
type NupsCtr struct {
	DB *sql.DB
}

//Kouji is a structure of kouji
type Kouji struct {
	ID    int
	Title string
}

//ListKi is a function to list up all ki
func (n *NupsCtr) ListKi(c *gin.Context) {
	kiList, err := model.KiAll(n.DB)
	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(kiList) == 0 {
		c.JSON(http.StatusOK, make([]int, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": kiList,
		"error":  nil,
	})
	return
}

//ListKouji is a function to list up all kouji of ki
func (n *NupsCtr) ListKouji(c *gin.Context) {
	return
}

//AddKi is a function to add new ki
func (n *NupsCtr) AddKi(c *gin.Context) {
	ki, _ := strconv.Atoi(c.Param("ki"))
	fmt.Print(ki)
	err := model.KiInsert(n.DB, ki)

	if err != nil {
		resp := errors.New(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

//GetKouji is a function to get kouji
func (n *NupsCtr) GetKouji(c *gin.Context) {
	return

}

//UploadKouji is a function to upload kouji
func (n *NupsCtr) UploadKouji(c *gin.Context) {
	return

}

//UpdateKouji is a function to update kouji
func (n *NupsCtr) UpdateKouji(c *gin.Context) {
	return

}

//DeleteKouji is a function to delete kouji
func (n *NupsCtr) DeleteKouji(c *gin.Context) {
	return

}
