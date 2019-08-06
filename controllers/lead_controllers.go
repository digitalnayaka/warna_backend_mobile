package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var leadModel = new(models.LeadModels)

type LeadController struct{}

func (m *LeadController) LeadListing(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	id_lead_mst_status := c.Query("id_lead_mst_status")
	favorite := c.Query("favorite")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.LeadListing(id_cms_users, id_lead_mst_status, favorite, limit, offset)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadListingVisit(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	types := c.Query("types")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.LeadListingVisit(types, id_cms_users, limit, offset)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadUpdateFav(c *gin.Context) {

	response := responseStruct

	id := c.PostForm("id")
	favorite := c.PostForm("favorite")
	updated_by := c.PostForm("updated_by")

	if id == "" {
		response.ApiMessage = "id required"
	} else if favorite == "" {
		response.ApiMessage = "favorite required"
	} else if updated_by == "" {
		response.ApiMessage = "updated_by required"
	} else {
		response = leadModel.LeadUpdateFav(id, favorite, updated_by)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadUpdateHapus(c *gin.Context) {

	response := responseStruct

	id := c.PostForm("id")
	updated_by := c.PostForm("updated_by")

	if id == "" {
		response.ApiMessage = "id required"
	} else if updated_by == "" {
		response.ApiMessage = "updated_by required"
	} else {
		response = leadModel.LeadUpdateHapus(id, updated_by)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadDetailTotal(c *gin.Context) {

	response := responseStruct

	id := c.Query("id")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.DetailLeadTotal(id, id_cms_users)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadDetail(c *gin.Context) {

	response := responseStruct

	id := c.Query("id")

	if id == "" {
		response.ApiMessage = "id required"
	} else {
		response = leadModel.DetailLead(id)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadLog(c *gin.Context) {

	response := responseStruct

	id := c.Query("id_lead")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.LeadLog(id, id_cms_users)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadVisum(c *gin.Context) {

	response := responseStruct

	id := c.Query("id_lead")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.LeadVisum(id, id_cms_users)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadNote(c *gin.Context) {

	response := responseStruct

	id := c.Query("id_lead")
	id_cms_users := c.Query("id_cms_users")

	if id == "" {
		response.ApiMessage = "id required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.LeadNote(id, id_cms_users)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadLogCreate(c *gin.Context) {

	response := responseStruct

	id_lead := c.PostForm("id_lead")
	duration := c.PostForm("duration")
	id_mst_log_desc := c.PostForm("id_mst_log_desc")
	recall := c.PostForm("recall")
	id_cms_users := c.PostForm("id_cms_users")
	id_modul := c.PostForm("id_modul")
	id_data := c.PostForm("id_data")
	jenis := c.PostForm("jenis")

	if id_lead == "" {
		response.ApiMessage = "id_lead required"
	} else if duration == "" {
		response.ApiMessage = "duration required"
	} else if id_mst_log_desc == "" {
		response.ApiMessage = "id_mst_log_desc required"
	} else if recall == "" {
		response.ApiMessage = "recall required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if id_modul == "" {
		response.ApiMessage = "id_modul required"
	} else if id_data == "" {
		response.ApiMessage = "id_data required"
	} else if jenis == "" {
		response.ApiMessage = "jenis required"
	} else {
		response = leadModel.LeadLogCreate(id_lead, duration, id_mst_log_desc, recall, id_cms_users, id_modul, id_data, jenis)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadNoteCreate(c *gin.Context) {

	response := responseStruct

	id_lead := c.PostForm("id_lead")
	id_cms_users := c.PostForm("id_cms_users")
	note := c.PostForm("note")

	if id_lead == "" {
		response.ApiMessage = "id_lead required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if note == "" {
		response.ApiMessage = "jenis required"
	} else {
		response = leadModel.LeadNoteCreate(id_lead, id_cms_users, note)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadVisumCreate(c *gin.Context) {

	response := responseStruct

	id_lead := c.Request.FormValue("id_lead")
	id_cms_users := c.Request.FormValue("id_cms_users")
	revisit := c.Request.FormValue("revisit")
	id_mst_visum_status := c.Request.FormValue("id_mst_visum_status")
	file, header, _ := c.Request.FormFile("photo")

	if id_lead == "" {
		response.ApiMessage = "id_lead required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if revisit == "" {
		response.ApiMessage = "revisit required"
	} else if id_mst_visum_status == "" {
		response.ApiMessage = "id_mst_visum_status required"
	} else if file == nil || header == nil {
		response.ApiMessage = "photo required"
	} else {
		response = leadModel.LeadVisumCreate(id_lead, id_cms_users, revisit, id_mst_visum_status, file, header)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadCheckPhoneNewDB(c *gin.Context) {

	response := responseStruct

	number := c.Query("number")

	response = leadModel.LeadPhoneCheckNewDB(number)

	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadCreate(c *gin.Context) {

	response := responseStruct

	id_mst_outlet := c.PostForm("id_mst_outlet")
	id_mst_data_source := c.PostForm("id_mst_data_source")
	id_cms_users := c.PostForm("id_cms_users")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	id_mst_job := c.PostForm("id_mst_job")
	id_lead_mst_status := c.PostForm("id_lead_mst_status")
	number := c.PostForm("number")
	id_mst_address := c.PostForm("id_mst_address")
	id_mst_category_address := c.PostForm("id_mst_category_address")
	id_mst_product := c.PostForm("id_mst_product")
	id_mst_unit := c.PostForm("id_mst_unit")
	nopol := c.PostForm("nopol")
	tax_status := c.PostForm("tax_status")
	owner := c.PostForm("owner")
	address := c.PostForm("address")
	address_bool := c.PostForm("address_bool")
	product_bool := c.PostForm("product_bool")

	if id_mst_outlet == "" {
		response.ApiMessage = "id_mst_outlet required"
	} else if id_mst_data_source == "" {
		response.ApiMessage = "id_mst_data_source required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if first_name == "" {
		response.ApiMessage = "first_name required"
	} else if id_mst_job == "" {
		response.ApiMessage = "id_mst_job required"
	} else if id_lead_mst_status == "" {
		response.ApiMessage = "id_lead_mst_status required"
	} else if address_bool == "" {
		response.ApiMessage = "address_bool required"
	} else if product_bool == "" {
		response.ApiMessage = "product_bool required"
	} else {

		address_bool_c, _ := strconv.ParseBool(address_bool)
		product_bool_c, _ := strconv.ParseBool(product_bool)

		response = leadModel.LeadCreate(id_mst_outlet, id_mst_data_source, id_cms_users,
			first_name, last_name, id_mst_job, id_lead_mst_status, number, id_mst_address, id_mst_category_address,
			id_mst_product, id_mst_unit, nopol, tax_status, owner, address, address_bool_c, product_bool_c)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadUpdate(c *gin.Context) {

	response := responseStruct

	id_lead := c.PostForm("id_lead")
	id_mst_outlet := c.PostForm("id_mst_outlet")
	id_mst_data_source := c.PostForm("id_mst_data_source")
	id_cms_users := c.PostForm("id_cms_users")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	id_mst_job := c.PostForm("id_mst_job")
	id_lead_mst_status := c.PostForm("id_lead_mst_status")
	number := c.PostForm("number")
	id_mst_address := c.PostForm("id_mst_address")
	id_mst_category_address := c.PostForm("id_mst_category_address")
	id_mst_product := c.PostForm("id_mst_product")
	id_mst_unit := c.PostForm("id_mst_unit")
	nopol := c.PostForm("nopol")
	tax_status := c.PostForm("tax_status")
	owner := c.PostForm("owner")
	address := c.PostForm("address")
	address_bool := c.PostForm("address_bool")
	product_bool := c.PostForm("product_bool")

	if id_lead == "" {
		response.ApiMessage = "id_lead required"
	} else if id_mst_outlet == "" {
		response.ApiMessage = "id_mst_outlet required"
	} else if id_mst_data_source == "" {
		response.ApiMessage = "id_mst_data_source required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else if first_name == "" {
		response.ApiMessage = "first_name required"
	} else if id_mst_job == "" {
		response.ApiMessage = "id_mst_job required"
	} else if id_lead_mst_status == "" {
		response.ApiMessage = "id_lead_mst_status required"
	} else if address_bool == "" {
		response.ApiMessage = "address_bool required"
	} else if product_bool == "" {
		response.ApiMessage = "product_bool required"
	} else {

		address_bool_c, _ := strconv.ParseBool(address_bool)
		product_bool_c, _ := strconv.ParseBool(product_bool)

		response = leadModel.LeadUpdate(id_lead, id_mst_outlet, id_mst_data_source, id_cms_users,
			first_name, last_name, id_mst_job, id_lead_mst_status, number, id_mst_address, id_mst_category_address,
			id_mst_product, id_mst_unit, nopol, tax_status, owner, address, address_bool_c, product_bool_c)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadColabCreate(c *gin.Context) {

	response := responseStruct

	id_lead := c.PostForm("id_lead")
	id_cms_users := c.PostForm("id_cms_users")

	if id_lead == "" {

		response.ApiMessage = "id_lead required"
	} else if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {

		response = leadModel.LeadCollabCreate(id_lead, id_cms_users)
	}

	c.JSON(http.StatusOK, response)
}

func (m *LeadController) LeadListingCollab(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.LeadCollabList(id_cms_users, limit, offset)
	}
	c.JSON(http.StatusOK, response)

}
func (m *LeadController) LeadSearch(c *gin.Context) {

	response := responseStruct

	id_cms_users := c.Query("id_cms_users")
	search := c.Query("first_name")

	if id_cms_users == "" {
		response.ApiMessage = "id_cms_users required"
	} else {
		response = leadModel.LeadSearch(id_cms_users, search)
	}
	c.JSON(http.StatusOK, response)

}
