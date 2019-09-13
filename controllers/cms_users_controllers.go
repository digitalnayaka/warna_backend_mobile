package controllers

import (
	"../models"
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

var cmsUserModel = new(models.CmsUser)

type CmsUserController struct{}

func (m *CmsUserController) Login(c *gin.Context) {

	response := structs.JsonResponseToken{}

	npm := c.PostForm("npm")
	password := c.PostForm("password")
	device_id := c.PostForm("device_id")
	user_agent := c.PostForm("user_agent")
	version := c.PostForm("version")

	if npm == "" {

		response.ApiMessage = "npm required"
	} else if password == "" {

		response.ApiMessage = "password required"
	} else if device_id == "" {
		response.ApiMessage = "device_id required"
	} else if user_agent == "" {
		response.ApiMessage = "user_agent required"
	} else if version == "" {
		response.ApiMessage = "version required"
	} else {

		response = cmsUserModel.Login(npm, password, device_id, user_agent, version)
	}

	c.JSON(http.StatusOK, response)

}
func (m *CmsUserController) Logout(c *gin.Context) {

	useragent := c.PostForm("useragent")
	description := c.PostForm("description")
	id_cms_users := c.PostForm("id_cms_users")
	version := c.PostForm("version")

	response := cmsUserModel.Logout(useragent, description, id_cms_users, version)

	c.JSON(http.StatusOK, response)
}
func (m *CmsUserController) CheckDeviceId(c *gin.Context) {

	response := responseStruct

	id := c.Query("id")
	device_id := c.Query("device_id")

	if id == "" {
		response.ApiMessage = "id required"
	} else if device_id == "" {
		response.ApiMessage = "device_id required"
	} else {
		response = cmsUserModel.CheckDeviceId(id, device_id)
	}
	c.JSON(http.StatusOK, response)

}
func (m *CmsUserController) UserLogListing(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = cmsUserModel.ActivityUser(id_cms_users, limit, offset)
	}
	c.JSON(http.StatusOK, response)

}
func (m *CmsUserController) PerformaIndicatorUser(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id")
	inputDate := c.Query("inputDate")
	inputDateBulanKemarin := c.Query("inputDateBulanKemarin")
	mst_branch := c.Query("mst_branch")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if inputDate == "" {
		response.ApiMessage = "inputDate required"
	} else if inputDateBulanKemarin == "" {
		response.ApiMessage = "inputDateBulanKemarin required"
	} else {
		response = cmsUserModel.PerformaIndicator(id_cms_users, inputDate, inputDateBulanKemarin, mst_branch)
	}
	c.JSON(http.StatusOK, response)

}
func (m *CmsUserController) Rekap(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id")
	inputDate := c.Query("inputdate")
	input_id_mst_outlet := c.Query("input_id_mst_outlet")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if inputDate == "" {
		response.ApiMessage = "inputDate required"
	} else if input_id_mst_outlet == "" {
		response.ApiMessage = "input_id_mst_outlet required"
	} else {
		response = cmsUserModel.RekapActivity(id_cms_users, inputDate, input_id_mst_outlet)
	}
	c.JSON(http.StatusOK, response)

}

func (m *CmsUserController) ChangePassword(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.PostForm("id_cms_users")
	password_lama := c.PostForm("password_lama")
	password_baru := c.PostForm("password_baru")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if password_lama == "" {
		response.ApiMessage = "password_lama required"
	} else if password_baru == "" {
		response.ApiMessage = "password_baru required"
	} else {
		response = cmsUserModel.ChangePassword(id_cms_users, password_lama, password_baru)
	}

	c.JSON(http.StatusOK, response)
}

func (m *CmsUserController) ActivityList(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	activity_schedule_start_date := c.Query("activity_schedule_start_date")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if activity_schedule_start_date == "" {
		response.ApiMessage = "activity_schedule_start_date"
	} else {
		response = cmsUserModel.ActivityListing(id_cms_users, activity_schedule_start_date)
	}

	c.JSON(http.StatusOK, response)
}
