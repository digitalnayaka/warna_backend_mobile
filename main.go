package main

import (
	"./controllers"
	"./middleware"
	"./models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//gin.SetMode(gin.ReleaseMode)
	cmsUser := new(controllers.CmsUserController)
	lead := new(controllers.LeadController)
	target := new(controllers.TargetController)
	mst := new(controllers.MstController)
	contact := new(controllers.ContactController)
	order := new(controllers.OrderController)
	router := gin.Default()

	pass, err := models.EncryptPassword("12345")
	fmt.Println(pass, err)

	public := router.Group("")
	{

		public.POST("/login", cmsUser.Login)
		public.Static("/img", "./pbc/img/all/")

	}

	v1 := router.Group("")
	v1.Use(middleware.Auth)
	{
		v1.GET("/check_android_id", cmsUser.CheckDeviceId)
		v1.GET("/user_log_listing", cmsUser.UserLogListing)
		v1.GET("/performa_indicator", cmsUser.PerformaIndicatorUser)
		v1.GET("/rekap_activity", cmsUser.Rekap)
		v1.POST("/change_password", cmsUser.ChangePassword)
		v1.GET("/activity_listing", cmsUser.ActivityList)

		v1.GET("/lead_listing", lead.LeadListing)
		v1.GET("/lead_listing_visit", lead.LeadListingVisit)
		v1.POST("/lead_update_fav", lead.LeadUpdateFav)
		v1.GET("/lead_detail_total", lead.LeadDetailTotal)
		v1.GET("/lead_detail", lead.LeadDetail)
		v1.GET("/lead_log_listing", lead.LeadLog)
		v1.GET("/lead_visum_listing", lead.LeadVisum)
		v1.GET("/lead_note_listing", lead.LeadNote)
		v1.PUT("/lead_hapus", lead.LeadUpdateHapus)
		v1.POST("/lead_log_create_logic", lead.LeadLogCreate)
		v1.POST("/lead_note_create", lead.LeadNoteCreate)
		v1.POST("/lead_visum_create", lead.LeadVisumCreate)
		v1.GET("/check_phone_new_db", lead.LeadCheckPhoneNewDB)
		v1.POST("/lead_create", lead.LeadCreate)
		v1.POST("/lead_update", lead.LeadUpdate)
		v1.POST("/lead_collab_create", lead.LeadColabCreate)
		v1.GET("/lead_listing_collab", lead.LeadListingCollab)
		v1.GET("/lead_search", lead.LeadSearch)

		v1.GET("/target_listing", target.TargetListing)
		v1.GET("/target_listing_visit", target.TargetListingVisit)
		v1.GET("/target_detail_total", target.TargetDetailTotal)
		v1.GET("/target_detail", target.TargetDetail)
		v1.GET("/target_log_listing", target.TargetLog)
		v1.GET("/target_visum_listing", target.TargetVisum)
		v1.GET("/target_note_listing", target.TargetNote)
		v1.POST("/target_log_create_logic", target.TargetLogCreate)
		v1.POST("/target_note_create", target.TargetNoteCreate)
		v1.POST("/target_visum_create", target.TargetVisumCreate)
		v1.GET("/target_search", target.TargetSearch)

		v1.GET("/mst_log_status", mst.MstLogStatus)
		v1.GET("/mst_log_desc", mst.MstLogDesc)
		v1.GET("/mst_visum_status", mst.MstVisumStatus)
		v1.GET("/mst_job", mst.MstJob)
		v1.GET("/mst_data_source", mst.MstDataSource)
		v1.GET("/mst_category_address", mst.MstCategoryAddress)
		v1.GET("/mst_unit_ufi_list", mst.MstUnitUfi)
		v1.GET("/mst_religion", mst.MstReligion)
		v1.GET("/contact_mst_status_marital", mst.ContactMstStatusMarital)
		v1.GET("/mst_place", mst.MstPlace)
		v1.GET("/contact_mst_status_employee", mst.ContactMstStatusEmployee)
		v1.GET("/mst_need_listing", mst.MstNeed)
		v1.GET("/mst_cabang_fif", mst.MstCabangFif)

		v1.GET("/check_add_contact", contact.CheckAddContact)
		v1.POST("/contact_create", contact.ContactCreate)
		v1.POST("/contact_update", contact.ContactUpdate)
		v1.GET("/contact_listing", contact.ContactList)
		v1.GET("/contact_info", contact.ContactInfoDetail)
		v1.GET("/contact_note_listing", contact.ContactNote)
		v1.POST("/contact_note_create", contact.ContactNoteCreate)
		v1.GET("/contact_collaborate_listing", contact.ContactCollabList)
		v1.GET("/contact_search", contact.ContactSearch)

		v1.GET("/order_listing_cfa", order.OrderListing)
		v1.POST("/order_create", order.OrderCreate)
		v1.GET("/order_detail", order.OrderDetail)
		v1.GET("/order_photo_listing", order.OrderPhotoList)
		v1.GET("/order_note_listing", order.OrderNote)
		v1.POST("/order_note_create", order.OrderNoteCreate)
		v1.POST("/order_upload_photo", order.OrderDocumentCreate)
		v1.POST("/order_photo_update", order.OrderDocumentUpdate)
		v1.POST("/order_photo_delete", order.OrderDocumentDelete)
		v1.POST("/order_otr_update", order.OrderOTRUpdate)
		v1.GET("/order_search", order.OrderSearch)

	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":4002")
}
