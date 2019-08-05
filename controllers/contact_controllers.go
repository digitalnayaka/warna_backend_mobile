package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../models"
)

var contactModel = new(models.ContactModels)

type ContactController struct {}

func (m *ContactController)CheckAddContact (c* gin.Context){

	response := responseStruct

	nik := c.Query("nik")
	id_cms_users := c.Query("id_cms_users")

	if nik == ""{
		response.ApiMessage = "nik required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else{

		response = contactModel.CheckAddContact(id_cms_users,nik)
	}

	c.JSON(http.StatusOK,response)


}
func (m *ContactController)ContactCreate (c* gin.Context){

	response := responseStruct

	modul := c.PostForm("modul")
	id_convert := c.PostForm("id_convert")
	nik := c.PostForm("nik")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	birth_place := c.PostForm("birth_place")
	birth_date := c.PostForm("birth_date")
	gender := c.PostForm("gender")
	id_mst_religion := c.PostForm("id_mst_religion")
	id_contact_mst_status_marital := c.PostForm("id_contact_mst_status_marital")
	id_mst_job := c.PostForm("id_mst_job")
	id_mst_data_source := c.PostForm("id_mst_data_source")
	id_cms_users := c.PostForm("id_cms_users")
	number := c.PostForm("number")
	number2 := c.PostForm("number2")
	address := c.PostForm("address")
	rt := c.PostForm("rt")
	rw := c.PostForm("rw")
	id_mst_address := c.PostForm("id_mst_address")
	id_mst_category_address := c.PostForm("id_mst_category_address")
	mother := c.PostForm("mother")
	family := c.PostForm("family")
	company := c.PostForm("company")
	position := c.PostForm("position")
	working_time := c.PostForm("working_time")
	income := c.PostForm("income")
	outlay := c.PostForm("outlay")
	id_contact_mst_status_place := c.PostForm("id_contact_mst_status_place")
	id_contact_mst_status_employee := c.PostForm("id_contact_mst_status_employee")


	if modul == ""{
		response.ApiMessage = "modul required"
	}else if id_convert == ""{
		response.ApiMessage = "id_convert required"
	}else if nik == ""{
		response.ApiMessage = "nik required"
	}else if first_name == ""{
		response.ApiMessage = "first_name required"
	}else if last_name == ""{
		response.ApiMessage = "last_name required"
	}else if gender == ""{
		response.ApiMessage = "gender required"
	}else if birth_date == ""{
		response.ApiMessage = "birth_date required"
	}else if birth_place == ""{
		response.ApiMessage = "birth_place required"
	}else if id_mst_religion == ""{
		response.ApiMessage = "id_mst_religion required"
	}else if id_contact_mst_status_marital == ""{
		response.ApiMessage = "id_contact_mst_status_marital required"
	}else if id_mst_job == ""{
		response.ApiMessage = "id_mst_job required"
	}else if id_mst_data_source == ""{
		response.ApiMessage = "id_mst_data_source required"
	}else if number == ""{
		response.ApiMessage = "number required"
	}else if number2 == ""{
		response.ApiMessage = "number2 required"
	}else if address == ""{
		response.ApiMessage = "address required"
	}else if rt == ""{
		response.ApiMessage = "rt required"
	}else if rw == ""{
		response.ApiMessage = "rw required"
	}else if id_mst_address == ""{
		response.ApiMessage = "id_mst_address required"
	}else if id_mst_category_address == ""{
		response.ApiMessage = "id_mst_category_address required"
	}else if mother == ""{
		response.ApiMessage = "mother required"
	}else if family == ""{
		response.ApiMessage = "family required"
	}else if company == ""{
		response.ApiMessage = "company required"
	}else if position == ""{
		response.ApiMessage = "position required"
	}else if working_time == ""{
		response.ApiMessage = "working_time required"
	}else if income == ""{
		response.ApiMessage = "income required"
	}else if outlay == ""{
		response.ApiMessage = "outlay required"
	}else if id_contact_mst_status_place == ""{
		response.ApiMessage = "id_contact_mst_status_place required"
	}else if id_contact_mst_status_employee == ""{
		response.ApiMessage = "id_contact_mst_status_employee required"
	}else{

		response = contactModel.ContactCreate(modul,id_convert,nik,first_name,last_name, birth_place, birth_date,
			gender , id_mst_religion , id_contact_mst_status_marital , id_mst_job  , id_mst_data_source , id_cms_users,
			number , number2 , address , rt  , rw  , id_mst_address , id_mst_category_address , mother , family , company ,
			position, working_time, income , outlay , id_contact_mst_status_place  , id_contact_mst_status_employee )
	}

	c.JSON(http.StatusOK,response)


}
func (m *ContactController)ContactList(c* gin.Context){

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_cms_users == ""{

		response.ApiMessage = "id_cms_users required"
	}else{

		response = contactModel.ContactList(id_cms_users,limit,offset)
	}

	c.JSON(http.StatusOK,response)
}
func (m *ContactController)ContactInfoDetail(c* gin.Context){

	response := responseStruct

	id := c.Query("id")

	if id == "" {
		response.ApiMessage = "id required"
	}else{

		response = contactModel.ContactDetail(id)
	}

	c.JSON(http.StatusOK,response)
}
func (m *ContactController)ContactNote(c* gin.Context){

	response := responseStruct

	id_contact := c.Query("id_contact")

	if id_contact == "" {
		response.ApiMessage = "id_contact required"
	}else{

		response = contactModel.ContactNote(id_contact)
	}

	c.JSON(http.StatusOK,response)
}

func (m *ContactController)ContactNoteCreate(c*gin.Context){

	response := responseStruct

	id_contact := c.PostForm("id_contact")
	id_cms_users := c.PostForm("id_cms_users")
	note := c.PostForm("note")

	if id_contact == "" {
		response.ApiMessage = "id_contact required"
	}else if id_cms_users == ""{
		response.ApiMessage = "id_cms_users required"
	}else if note == ""{
		response.ApiMessage = "jenis required"
	}else{
		response = contactModel.ContactNoteCreate(id_contact,id_cms_users,note)
	}
	c.JSON(http.StatusOK,response)

}
func (m *ContactController)ContactCollabList(c*gin.Context){

	response := responseStruct

	id_contact := c.Query("id_contact")

	if id_contact == "" {
		response.ApiMessage = "id_contact required"
	}else{
		response = contactModel.ContactCollabList(id_contact)
	}
	c.JSON(http.StatusOK,response)

}
func (m *ContactController)ContactUpdate (c* gin.Context){

	response := responseStruct

	id_contact := c.PostForm("id_contact")
	nik := c.PostForm("nik")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	birth_place := c.PostForm("birth_place")
	birth_date := c.PostForm("birth_date")
	gender := c.PostForm("gender")
	id_mst_religion := c.PostForm("id_mst_religion")
	id_contact_mst_status_marital := c.PostForm("id_contact_mst_status_marital")
	id_mst_job := c.PostForm("id_mst_job")
	id_mst_data_source := c.PostForm("id_mst_data_source")
	id_cms_users := c.PostForm("id_cms_users")
	number := c.PostForm("number")
	number2 := c.PostForm("number2")
	address := c.PostForm("address")
	rt := c.PostForm("rt")
	rw := c.PostForm("rw")
	id_mst_address := c.PostForm("id_mst_address")
	id_mst_category_address := c.PostForm("id_mst_category_address")
	mother := c.PostForm("mother")
	family := c.PostForm("family")
	company := c.PostForm("company")
	position := c.PostForm("position")
	working_time := c.PostForm("working_time")
	income := c.PostForm("income")
	outlay := c.PostForm("outlay")
	id_contact_mst_status_place := c.PostForm("id_contact_mst_status_place")
	id_contact_mst_status_employee := c.PostForm("id_contact_mst_status_employee")
	id_number1 := c.PostForm("id_number1")
	id_number2 := c.PostForm("id_number2")


	if id_contact == ""{
		response.ApiMessage = "id_contact required"
	}else{

		response = contactModel.ContactUpdate(id_contact,nik,first_name,last_name, birth_place, birth_date,
			gender , id_mst_religion , id_contact_mst_status_marital , id_mst_job  , id_mst_data_source , id_cms_users,
			number , number2 , address , rt  , rw  , id_mst_address , id_mst_category_address , mother , family , company ,
			position, working_time, income , outlay , id_contact_mst_status_place  , id_contact_mst_status_employee , id_number1,id_number2)
	}

	c.JSON(http.StatusOK,response)


}


