package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var orderModels = models.OrderModels{}

type OrderController struct {
}

func (m *OrderController) OrderListing(c *gin.Context) {

	id_cms_users := c.Query("id_cms_users")
	id_order_mst_status := c.Query("id_order_mst_status")
	id_contact := c.Query("id_contact")
	limit := c.Query("limit")
	offset := c.Query("offset")

	response := orderModels.OrderListing(id_cms_users, id_contact, id_order_mst_status, limit, offset)

	c.JSON(http.StatusOK, response)

}
func (m *OrderController) OrderDetail(c *gin.Context) {

	response := responseStruct

	id := c.Query("id")

	if id == "" {
		response.ApiMessage = "id required"
	} else {
		response = orderModels.OrderDetail(id)
	}

	c.JSON(http.StatusOK, response)

}

func (m *OrderController) OrderCreate(c *gin.Context) {

	response := responseStruct

	id_contact := c.PostForm("id_contact")
	id_cms_users := c.PostForm("id_cms_users")
	id_mst_outlet := c.PostForm("id_mst_outlet")
	id_mst_product := c.PostForm("id_mst_product")
	id_order_mst_status := c.PostForm("id_order_mst_status")
	id_mst_data_source := c.PostForm("id_mst_data_source")
	category := c.PostForm("category")
	id_mst_cabang_fif := c.PostForm("id_mst_cabang_fif")
	status_address := c.PostForm("status_address")
	survey := c.PostForm("survey")
	id_mst_unit := c.PostForm("id_mst_unit")
	nopol := c.PostForm("nopol")
	tax_status := c.PostForm("tax_status")
	owner := c.PostForm("owner")
	otr_taksasi := c.PostForm("otr_taksasi")
	nomor_taksasi := c.PostForm("nomor_taksasi")
	otr_custom := c.PostForm("otr_custom")
	plafond := c.PostForm("plafond")
	down_payment := c.PostForm("down_payment")
	installment := c.PostForm("installment")
	tenor := c.PostForm("tenor")
	need := c.PostForm("need")
	name := c.PostForm("name")
	birth_place := c.PostForm("birth_place")
	birth_date := c.PostForm("birth_date")
	id_mst_job := c.PostForm("id_mst_job")
	company := c.PostForm("company")
	id_contact_mst_status_employee := c.PostForm("id_contact_mst_status_employee")
	position := c.PostForm("position")
	working_time := c.PostForm("working_time")
	income := c.PostForm("income")
	outlay := c.PostForm("outlay")

	if id_contact == "" {
		response.ApiMessage = "id_contact required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if id_mst_outlet == "" {
		response.ApiMessage = "id_mst_outlet required"
	} else if id_mst_product == "" {
		response.ApiMessage = "id_mst_product required"
	} else if id_order_mst_status == "" {
		response.ApiMessage = "id_order_mst_status required"
	} else if id_mst_data_source == "" {
		response.ApiMessage = "id_mst_data_source required"
	} else if id_mst_cabang_fif == "" {
		response.ApiMessage = "id_mst_cabang_fif required"
	} else if id_mst_unit == "" {
		response.ApiMessage = "id_mst_unit required"
	} else if nopol == "" {
		response.ApiMessage = "nopol required"
	} else if tax_status == "" {
		response.ApiMessage = "tax_status required"
	} else if owner == "" {
		response.ApiMessage = "owner required"
	} else if otr_custom == "" {
		response.ApiMessage = "otr_custom required"
	} else if plafond == "" {
		response.ApiMessage = "plafond required"
	} else if down_payment == "" {
		response.ApiMessage = "down_payment required"
	} else if installment == "" {
		response.ApiMessage = "installment required"
	} else if tenor == "" {
		response.ApiMessage = "tenor required"
	} else if need == "" {
		response.ApiMessage = "need required"
	} else if name == "" {
		response.ApiMessage = "name required"
	} else if birth_place == "" {
		response.ApiMessage = "birth_place required"
	} else if birth_date == "" {
		response.ApiMessage = "birth_date required"
	} else {
		response = orderModels.OrderCreate(id_contact, id_cms_users, id_mst_outlet, id_mst_product,
			id_order_mst_status, id_mst_data_source, category, id_mst_cabang_fif, status_address,
			survey,
			id_mst_unit, nopol, tax_status, owner, otr_taksasi,
			nomor_taksasi,
			otr_custom, plafond, down_payment, installment, tenor, need,
			name, birth_place, birth_date, id_mst_job, company, id_contact_mst_status_employee,
			position, working_time, income, outlay)
	}

	c.JSON(http.StatusOK, response)
}

func (m *OrderController) OrderPhotoList(c *gin.Context) {

	response := responseStruct

	id_order := c.Query("id_order")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_order == "" {
		response.ApiMessage = "id_order required"
	} else {

		response = orderModels.OrderPhotoList(id_order, limit, offset)
	}

	c.JSON(http.StatusOK, response)
}

func (m *OrderController) OrderNote(c *gin.Context) {

	response := responseStruct

	id_order := c.Query("id_order")

	if id_order == "" {
		response.ApiMessage = "id_order required"
	} else {

		response = orderModels.OrderNote(id_order)
	}

	c.JSON(http.StatusOK, response)
}
func (m *OrderController) OrderNoteCreate(c *gin.Context) {

	response := responseStruct

	id_order := c.PostForm("id_order")
	id_cms_users := c.PostForm("id_cms_users")
	note := c.PostForm("note")

	if id_order == "" {
		response.ApiMessage = "id_order required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if note == "" {
		response.ApiMessage = "jenis required"
	} else {
		response = orderModels.OrderNoteCreate(id_order, id_cms_users, note)
	}
	c.JSON(http.StatusOK, response)

}

func (m *OrderController) OrderDocumentCreate(c *gin.Context) {

	response := responseStruct

	id_order := c.Request.FormValue("id_order")
	id_cms_users := c.Request.FormValue("id_cms_users")
	description := c.Request.FormValue("description")
	file, header, _ := c.Request.FormFile("photo")

	if id_order == "" {
		response.ApiMessage = "id_order required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if description == "" {
		response.ApiMessage = "description required"
	} else if file == nil || header == nil {
		response.ApiMessage = "photo required"
	} else {
		response = orderModels.OrderDocumentCreate(id_order, id_cms_users, description, file, header)
	}
	c.JSON(http.StatusOK, response)

}
func (m *OrderController) OrderDocumentUpdate(c *gin.Context) {

	response := responseStruct

	id_order := c.Request.FormValue("id_order")
	id := c.Request.FormValue("id")
	description := c.Request.FormValue("description")
	file, header, _ := c.Request.FormFile("photo")

	if id_order == "" {
		response.ApiMessage = "id_order required"
	} else if id == "" {
		response.ApiMessage = "id required"
	} else if description == "" {
		response.ApiMessage = "description required"
	} else if file == nil || header == nil {
		response.ApiMessage = "photo required"
	} else {
		response = orderModels.OrderDocumentUpdate(id, id_order, description, file, header)
	}
	c.JSON(http.StatusOK, response)

}
func (m *OrderController) OrderDocumentDelete(c *gin.Context) {

	response := responseStruct

	id := c.PostForm("id")

	if id == "" {
		response.ApiMessage = "id required"
	} else {
		response = orderModels.OrderDocumentDelete(id)
	}
	c.JSON(http.StatusOK, response)

}
func (m *OrderController) OrderOTRUpdate(c *gin.Context) {

	response := responseStruct

	id := c.PostForm("id")
	otr_taksasi := c.PostForm("otr_taksasi")
	nomor_taksasi := c.PostForm("nomor_taksasi")
	id_cms_users := c.PostForm("updated_by")

	if id == "" {
		response.ApiMessage = "id required"
	} else if otr_taksasi == "" {
		response.ApiMessage = "otr_taksasi required"
	} else if nomor_taksasi == "" {
		response.ApiMessage = "nomor_taksasi required"
	} else if id_cms_users == "" {
		response.ApiMessage = "updated_by required"
	} else {

		response = orderModels.OrderOTRUpdate(id, otr_taksasi, nomor_taksasi, id_cms_users)
	}

	c.JSON(http.StatusOK, response)
}

func (m *OrderController) OrderSearch(c *gin.Context) {

	id_cms_users := c.Query("id_cms_users")
	search := c.Query("first_name")

	response := orderModels.OrderSearch(id_cms_users, search)

	c.JSON(http.StatusOK, response)

}
