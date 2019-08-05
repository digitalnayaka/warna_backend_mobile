package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var targetModel = new(models.TargetModels)

type TargetController struct {}

func (m *TargetController)TargetListing (c* gin.Context){

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	id_target_mst_status := c.Query("id_target_mst_status")
	favorite := c.Query("favorite")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	}else{
		response = targetModel.TargetListing(id_cms_users,id_target_mst_status,favorite,limit,offset)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetListingVisit (c* gin.Context){

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	types := c.Query("types")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	}else{
		response = targetModel.TargetListingVisit(types,id_cms_users,limit,offset)
	}
	c.JSON(http.StatusOK,response)

}

func (m *TargetController)TargetDetailTotal (c* gin.Context){

	response := responseStruct

	id := c.Query("id")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else{
		response = targetModel.DetailTargetTotal(id,id_cms_users)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetDetail (c* gin.Context){

	response := responseStruct

	id := c.Query("id")

	if id == "" {
		response.ApiMessage = "id required"
	}else{
		response = targetModel.DetailTarget(id)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetLog (c* gin.Context){

	response := responseStruct

	id := c.Query("id_target")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else{
		response = targetModel.TargetLog(id,id_cms_users)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetVisum (c* gin.Context){

	response := responseStruct

	id := c.Query("id_target")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else{
		response = targetModel.TargetVisum(id,id_cms_users)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetNote (c* gin.Context){

	response := responseStruct

	id := c.Query("id_target")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else{
		response = targetModel.TargetNote(id,id_cms_users)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetLogCreate(c*gin.Context){

	response := responseStruct

	id_target := c.PostForm("id_target")
	duration := c.PostForm("duration")
	id_mst_log_desc := c.PostForm("id_mst_log_desc")
	recall := c.PostForm("recall")
	id_cms_users := c.PostForm("id_cms_users")
	id_modul := c.PostForm("id_modul")
	id_data := c.PostForm("id_data")
	jenis := c.PostForm("jenis")

	if id_target == "" {
		response.ApiMessage = "id_target required"
	}else if duration == ""{
		response.ApiMessage = "duration required"
	}else if id_mst_log_desc == ""{
		response.ApiMessage = "id_mst_log_desc required"
	}else if recall == ""{
		response.ApiMessage = "recall required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else if id_modul == ""{
		response.ApiMessage = "id_modul required"
	}else if id_data == ""{
		response.ApiMessage = "id_data required"
	}else if jenis == ""{
		response.ApiMessage = "jenis required"
	}else{
		response = targetModel.TargetLogCreate(id_target,duration,id_mst_log_desc,recall,id_cms_users,id_modul,id_data,jenis)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetNoteCreate(c*gin.Context){

	response := responseStruct

	id_target := c.PostForm("id_target")
	id_cms_users := c.PostForm("id_cms_users")
	note := c.PostForm("note")

	if id_target == "" {
		response.ApiMessage = "id_target required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else if note == ""{
		response.ApiMessage = "jenis required"
	}else{
		response = targetModel.TargetNoteCreate(id_target,id_cms_users,note)
	}
	c.JSON(http.StatusOK,response)

}
func (m *TargetController)TargetVisumCreate(c*gin.Context){

	response := responseStruct

	id_target := c.Request.FormValue("id_target")
	id_cms_users := c.Request.FormValue("id_cms_users")
	revisit := c.Request.FormValue("revisit")
	id_mst_visum_status := c.Request.FormValue("id_mst_visum_status")
	file, header, _ := c.Request.FormFile("photo")

	if id_target == "" {
		response.ApiMessage = "id_target required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else if revisit == ""{
		response.ApiMessage = "revisit required"
	}else if id_mst_visum_status == ""{
		response.ApiMessage = "id_mst_visum_status required"
	}else if file == nil || header == nil{
		response.ApiMessage = "photo required"
	}else{
		response = targetModel.TargetVisumCreate(id_target,id_cms_users,revisit,id_mst_visum_status,file,header)
	}
	c.JSON(http.StatusOK,response)

}
//func (m *TargetController)TargetCheckPhoneNewDB(c*gin.Context){
//
//	response := responseStruct
//
//	number := c.Query("number")
//
//	response = targetModel.TargetPhoneCheckNewDB(number)
//
//	c.JSON(http.StatusOK,response)
//
//}



