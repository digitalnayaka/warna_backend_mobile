package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var mstModel = new(models.MstModels)

type MstController struct{}

func (m *MstController) MstLogStatus(c *gin.Context) {

	response := mstModel.MstLogStatus()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstVisumStatus(c *gin.Context) {

	response := mstModel.MstVisumStatus()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstLogDesc(c *gin.Context) {

	response := responseStruct

	id_mst_log_status := c.Query("id_mst_log_status")

	response = mstModel.MstLogDesc(id_mst_log_status)
	c.JSON(http.StatusOK, response)

}

func (m *MstController) MstJob(c *gin.Context) {

	response := mstModel.MstJob()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstReligion(c *gin.Context) {

	response := mstModel.MstReligion()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) ContactMstStatusMarital(c *gin.Context) {

	response := mstModel.ContactMstStatusMarital()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstPlace(c *gin.Context) {

	response := mstModel.MstPlace()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) ContactMstStatusEmployee(c *gin.Context) {

	response := mstModel.ContactMstStatusEmployee()
	c.JSON(http.StatusOK, response)

}

func (m *MstController) MstDataSource(c *gin.Context) {

	response := mstModel.MstDatasource()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstCategoryAddress(c *gin.Context) {

	response := mstModel.MstCategoryAddress()
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstUnitUfi(c *gin.Context) {

	id_mst_branch := c.Query("id_mst_branch")
	merk := c.Query("merk")
	year := c.Query("year")

	response := mstModel.MstUnitUfi(id_mst_branch, merk, year)
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstCabangFif(c *gin.Context) {

	branch_name := c.Query("branch_name")
	pos_name := c.Query("pos_name")

	response := mstModel.MstCabangFif(branch_name, pos_name)
	c.JSON(http.StatusOK, response)

}
func (m *MstController) MstNeed(c *gin.Context) {

	response := mstModel.MstNeed()
	c.JSON(http.StatusOK, response)

}

func (m *MstController) MstUnitMerk(c *gin.Context) {

	id_mst_branch := c.Query("id_mst_branch")

	response := mstModel.MstUnitMerk(id_mst_branch)
	c.JSON(http.StatusOK, response)
}

func (m *MstController) MstUnitYear(c *gin.Context) {

	id_mst_branch := c.Query("id_mst_branch")
	merk := c.Query("merk")

	response := mstModel.MstUnitYear(id_mst_branch, merk)
	c.JSON(http.StatusOK, response)
}
