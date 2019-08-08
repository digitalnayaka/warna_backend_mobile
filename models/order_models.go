package models

import (
	"../structs"
	"fmt"
	"mime/multipart"
	"strconv"
)

type OrderModels struct {
}

func (m *OrderModels) OrderListing(id_cms_users string, id_contact string, id_order_mst_status string, limit string, offset string) structs.JsonResponse {

	response := responseStruct
	orderListingStruct := []structs.OrderListing{}
	response.Data = orderListingStruct

	err := idb.DB.Table(order + " a").Select("a.created_at ," +
		" a.id ," +
		"a.id_contact , " +
		"b.first_name as contact_first_name , " +
		"b.last_name as contact_last_name , " +
		"c.status as order_mst_status_status , " +
		"d.plafond ," +
		"f.model").
		Joins("left join contact b " +
			"on b.id = a.id_contact ").
		Joins("left join order_mst_status c " +
			"on c.id = a.id_order_mst_status ").
		Joins("left join order_loan d " +
			"on d.id_order = a.id ").
		Joins("left join order_product_ufi e " +
			"on e.id_order = a.id ").
		Joins("left join mst_unit f " +
			"on f.id = e.id_mst_unit ")

	if id_cms_users != "" {
		err = err.Where("a.id_cms_users = " + id_cms_users + "")
	}

	if id_contact != "" {
		err = err.Where("a.id_contact = " + id_contact + "")
	}

	if id_order_mst_status != "" {
		err = err.Where("a.id_order_mst_status = " + id_order_mst_status + "")
	}

	err = err.Order("a.created_at desc , a.updated_at desc")

	if limit != "" {
		err = err.Limit(limit)
	}

	if offset != "" {
		err = err.Offset(offset)
	}

	errx := err.Find(&orderListingStruct).Error

	if errx != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = orderListingStruct
	}

	return response

}

func (m *OrderModels) OrderCreate(
	id_contact string, id_cms_users string, id_mst_outlet string, id_mst_product string,
	id_order_mst_status string, id_mst_data_source string, category string, id_mst_cabang_fif string, status_address string,
	survey string,
	id_mst_unit string, nopol string, tax_status string, owner string, otr_taksasi string,
	nomor_taksasi string,
	otr_custom string, plafond string, down_payment string, installment string, tenor string, need string,
	name string, birth_place string, birth_date string, id_mst_job string, company string, id_contact_mst_status_employee string,
	position string, working_time string, income string, outlay string) structs.JsonResponse {

	response := responseStruct
	orderCreateStruct := structs.Order{}
	orderUfiStruct := structs.OrderUfi{}
	orderLoanStruct := structs.OrderLoan{}
	orderSuretyStrcut := structs.OrderSurety{}

	id_contact_c, _ := strconv.ParseInt(id_contact, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
	id_mst_outlet_c, _ := strconv.ParseInt(id_mst_outlet, 10, 64)
	id_mst_product_c, _ := strconv.ParseInt(id_mst_product, 10, 64)
	id_order_mst_status_c, _ := strconv.ParseInt(id_order_mst_status, 10, 64)
	id_mst_data_source_c, _ := strconv.ParseInt(id_mst_data_source, 10, 64)
	id_mst_cabang_fif_c, _ := strconv.ParseInt(id_mst_cabang_fif, 10, 64)

	orderCreateStruct.CreatedAt = GetTimeNow()
	orderCreateStruct.IdContact = id_contact_c
	orderCreateStruct.IdCmsUsers = id_cms_users_c
	orderCreateStruct.IdMstDataSource = id_mst_data_source_c
	orderCreateStruct.IdMstProduct = id_mst_product_c
	orderCreateStruct.IdMstOutlet = id_mst_outlet_c
	orderCreateStruct.Category = category
	orderCreateStruct.IdMstCabangFif = id_mst_cabang_fif_c
	orderCreateStruct.IdOrderMstStatus = id_order_mst_status_c
	orderCreateStruct.StatusAddress = status_address
	orderCreateStruct.Survey = survey
	orderCreateStruct.NoOrder = "DummyNoOrder"

	fmt.Println(orderCreateStruct.Survey)

	err := idb.DB.Table("order").Create(&orderCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd
	} else {

		_ = idb.DB.Exec("update " + order + " set no_order = (SELECT concat(to_char(now(),'YYYYMM'), 'OR' , " +
			"LPAD( cast((select count(id)+1 from " + order + " where created_at >= now()::date) as text),6,'0'))) where id = " + fmt.Sprint(orderCreateStruct.Id) + "").RowsAffected

		orderUfiCreate := func() {

			id_mst_unit_c, _ := strconv.ParseInt(id_mst_unit, 10, 64)
			otr_taksasi_c, _ := strconv.Atoi(otr_taksasi)

			orderUfiStruct.IdCmsUsers = id_cms_users_c
			orderUfiStruct.CreatedAt = GetTimeNow()
			orderUfiStruct.TaxStatus = tax_status
			orderUfiStruct.Owner = owner
			orderUfiStruct.IdMstUnit = id_mst_unit_c
			orderUfiStruct.Nopol = nopol
			orderUfiStruct.IdOrder = orderCreateStruct.Id
			orderUfiStruct.NomorTaksasi = nomor_taksasi
			orderUfiStruct.OtrTaksasi = otr_taksasi_c

			_ = idb.DB.Table("order_product_ufi").Create(&orderUfiStruct).Error
		}
		orderLoanCreate := func() {

			otr_custom_c, _ := strconv.Atoi(otr_custom)
			plafond_c, _ := strconv.Atoi(plafond)
			down_payment_c, _ := strconv.Atoi(down_payment)
			installment_c, _ := strconv.Atoi(installment)
			tenor, _ := strconv.Atoi(tenor)

			orderLoanStruct.IdOrder = orderCreateStruct.Id
			orderLoanStruct.CreatedAt = GetTimeNow()
			orderLoanStruct.IdCmsUsers = id_cms_users_c
			orderLoanStruct.DownPayment = down_payment_c
			orderLoanStruct.Installment = installment_c
			orderLoanStruct.Need = need
			orderLoanStruct.OtrCustom = otr_custom_c
			orderLoanStruct.Plafond = plafond_c
			orderLoanStruct.Tenor = tenor

			fmt.Println(orderLoanStruct)

			_ = idb.DB.Table("order_loan").Create(&orderLoanStruct).Error
		}
		orderSurety := func() {

			id_mst_job_c, _ := strconv.ParseInt(id_mst_job, 10, 64)
			id_contact_mst_status_employee_c, _ := strconv.ParseInt(id_contact_mst_status_employee, 10, 64)
			working_time_c, _ := strconv.Atoi(working_time)
			income_c, _ := strconv.Atoi(income)
			outlay_c, _ := strconv.Atoi(outlay)

			orderSuretyStrcut.IdCmsUsers = id_cms_users_c
			orderSuretyStrcut.CreatedAt = GetTimeNow()
			orderSuretyStrcut.IdOrder = orderCreateStruct.Id
			orderSuretyStrcut.WorkingTime = working_time_c
			orderSuretyStrcut.Position = position
			orderSuretyStrcut.Outlay = outlay_c
			orderSuretyStrcut.Income = income_c
			orderSuretyStrcut.IdContactMstStatusEmployee = id_contact_mst_status_employee_c
			orderSuretyStrcut.Company = company
			orderSuretyStrcut.BirthPlace = birth_place
			orderSuretyStrcut.BirthDate = birth_date
			orderSuretyStrcut.IdMstJob = id_mst_job_c
			orderSuretyStrcut.Name = name

			_ = idb.DB.Table("order_surety").Create(&orderSuretyStrcut).Error

		}
		logCreate := func() {

			mstLogsStruct := structs.MstLogs{}

			mstLogsStruct.IdCmsUsers = id_cms_users_c
			mstLogsStruct.IdData = orderCreateStruct.Id
			mstLogsStruct.IdModul = 5
			mstLogsStruct.Jenis = "create"
			mstLogsStruct.CreatedAt = GetTimeNow()

			_ = idb.DB.Table("mst_logs").Create(&mstLogsStruct).Error

		}

		orderUfiCreate()
		orderLoanCreate()
		orderSurety()
		go logCreate()

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = orderCreateStruct.Id
	}

	return response
}

func (m *OrderModels) OrderDetail(id_order string) structs.JsonResponse {

	response := responseStruct
	orderDetailStruct := structs.OrderDetail{}
	orderInfoStruct := structs.OrderInfoDetail{}
	orderUfiStruct := structs.OrderProductUfiDetail{}
	orderLoanStruct := structs.OrderLoan{}
	orderSuretyStruct := structs.OrderSuretyDetail{}
	orderContactDetailStruct := structs.ContactDetailInfo{}

	err := idb.DB.Raw("select a.id ,a.id_contact , b.datasource as mst_data_source_datasource , " +
		"c.outlet_name as mst_outlet_outlet_name , e.nama as mst_product_nama , " +
		"a.no_order , f.status as order_mst_status_status , h.name as cms_users_name , a.created_at , a.category , " +
		"a.status_address , a.survey , a.id_mst_cabang_fif , d.branch_name as cabang_fif , d.pos_name as pos_fif , " +
		"a.id_order_mst_reason , g.reason from " + order + " a " +
		"left join mst_data_source b " +
		"on b.id = a.id_mst_data_source " +
		"left join mst_outlet c " +
		"on c.id = a.id_mst_outlet " +
		"left join mst_cabang_fif d " +
		"on d.id = a.id_mst_cabang_fif " +
		"left join mst_product e " +
		"on e.id = a.id_mst_product " +
		"left join order_mst_status f " +
		"on f.id = a.id_order_mst_status " +
		"left join order_mst_reason g " +
		"on g.id = a.id_order_mst_reason " +
		"left join cms_users h " +
		"on h.id = a.id_cms_users " +
		"where a.id = " + id_order + "").Scan(&orderInfoStruct).Error

	if err != nil {

		response.ApiMessage = errDB
	} else {

		_ = idb.DB.Raw("select a.id , b.year as mst_unit_year , b.kode_unit as mst_unit_kode_unit , " +
			"b.merk as mst_unit_merk , b.type as mst_unit_type , b.model as mst_unit_model , " +
			"b.otr as mst_unit_otr , a.nopol , a.tax_status , a.owner , a.otr_taksasi , a.nomor_taksasi " +
			"from order_product_ufi a " +
			"left join mst_unit b " +
			"on b.id = a.id_mst_unit " +
			"where a.id_order = " + id_order + "").Scan(&orderUfiStruct).Error

		_ = idb.DB.Table("order_loan").Where("id_order = " + id_order + "").First(&orderLoanStruct).Error
		//_ = idb.DB.Table("order_surety a").Select(" a.* , b.job as mst_job_job , " +
		//	"c.status as contact_mst_status_employee_status").
		//	Joins("left join mst_job b on b.id = a.id_mst_job").
		//	Joins("left join contact_mst_status_employee c on c.id = a.id_contact_mst_status_employee").
		//	Where("a.id_order = " + id_order + "").First(&orderSuretyStruct).Error

		_ = idb.DB.Raw("select a.* , b.job as mst_job_job , c.status as contact_mst_status_employee_status " +
			"from order_surety a " +
			"left join mst_job b " +
			"on b.id = a.id_mst_job " +
			"left join contact_mst_status_employee c " +
			"on c.id = a.id_contact_mst_status_employee " +
			"where a.id_order = " + id_order + "").Scan(&orderSuretyStruct).Error

		contactModel := ContactModels{}
		responseContact := contactModel.ContactDetail(fmt.Sprint(orderInfoStruct.IdContact))

		if responseContact.ApiStatus != 0 {

			fieldContact, check := responseContact.Data.(structs.ContactDetailInfo)

			if check {

				orderContactDetailStruct = fieldContact
			}

		}

		orderDetailStruct.ContactDetailInfo = orderContactDetailStruct
		orderDetailStruct.OrderSurety = orderSuretyStruct
		orderDetailStruct.OrderLoan = orderLoanStruct
		orderDetailStruct.OrderProductUfiDetail = orderUfiStruct
		orderDetailStruct.OrderInfoDetail = orderInfoStruct

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = orderDetailStruct

	}

	return response
}

func (m *OrderModels) OrderPhotoList(id_order string, limit string, offset string) structs.JsonResponse {

	response := responseStruct
	orderPhotoStruct := []structs.OrderPhoto{}
	response.Data = orderPhotoStruct

	err := idb.DB.Table("order_photo").Where("id_order = " + id_order + "")

	if limit != "" {

		err = err.Limit(limit)
	}

	if offset != "" {

		err = err.Offset(offset)
	}

	err = err.Find(&orderPhotoStruct)
	errx := err.Error

	if errx != nil {

		response.ApiMessage = errDB
	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = orderPhotoStruct
	}

	return response
}

func (m *OrderModels) OrderNote(id_order string) structs.JsonResponse {

	contactNoteStruct := []structs.ContactNote{}
	response := responseStruct

	err := idb.DB.Raw("select a.id , a.note , b.name as cms_users_name , a.created_at from order_note a " +
		"left join cms_users b " +
		"on b.id = a.id_cms_users where a.id_order = " + id_order + " order by a.created_at desc").Scan(&contactNoteStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiMessage = succ
		response.Data = contactNoteStruct
	}

	return response

}

func (m *OrderModels) OrderNoteCreate(id_order string, id_cms_users string, note string) structs.JsonResponse {

	response := responseStruct
	contactNoteCreateStruct := structs.OrderNoteCreate{}

	id_order_c, _ := strconv.ParseInt(id_order, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)

	contactNoteCreateStruct.IdOrder = id_order_c
	contactNoteCreateStruct.IdCmsUsers = id_cms_users_c
	contactNoteCreateStruct.Note = note
	contactNoteCreateStruct.CreatedAt = GetTimeNow()

	err := idb.DB.Table("order_note").Create(&contactNoteCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd

	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
	}

	return response
}

func (m *OrderModels) OrderDocumentCreate(id_order string, id_cms_users string, description string, files multipart.File, header *multipart.FileHeader) structs.JsonResponse {

	response := responseStruct
	orderDocStruct := structs.OrderPhoto{}

	id_order_c, _ := strconv.ParseInt(id_order, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)

	url := UploadImage("document", id_order, files, header)

	if url != "" {

		orderDocStruct.IdOrder = id_order_c
		orderDocStruct.CreatedAt = GetTimeNow()
		orderDocStruct.IdCmsUsers = id_cms_users_c
		orderDocStruct.Photo = url
		orderDocStruct.Description = description

		err := idb.DB.Table("order_photo").Create(&orderDocStruct).Error

		if err != nil {

			response.ApiMessage = errDBAdd
		} else {

			response.ApiStatus = 1
			response.ApiMessage = succ
		}

	} else {

		response.ApiMessage = "Gagal upload document"
	}

	return response
}
func (m *OrderModels) OrderDocumentUpdate(id string, id_order string, description string, files multipart.File, header *multipart.FileHeader) structs.JsonResponse {

	response := responseStruct
	orderDocStruct := structs.OrderPhoto{}

	id_order_c, _ := strconv.ParseInt(id_order, 10, 64)

	url := UploadImage("document", id_order, files, header)

	if url != "" {

		orderDocStruct.IdOrder = id_order_c
		orderDocStruct.UpdatedAt = GetTimeNow()
		orderDocStruct.Photo = url
		orderDocStruct.Description = description

		err := idb.DB.Table("order_photo").Where("id = " + id + "").Update(&orderDocStruct).Error

		if err != nil {

			response.ApiMessage = errDBUpdate
		} else {

			response.ApiStatus = 1
			response.ApiMessage = succ
		}

	} else {

		response.ApiMessage = "Gagal upload document"
	}

	return response
}
func (m *OrderModels) OrderDocumentDelete(id string) structs.JsonResponse {

	response := responseStruct

	err := idb.DB.Exec("delete from order_photo where id = " + id + "").RowsAffected

	if err >= 1 {
		response.ApiStatus = 1
		response.ApiMessage = succ
	} else {
		response.ApiMessage = errDBDelete
	}

	return response
}
func (m *OrderModels) OrderOTRUpdate(id string, otr_taksasi string, nomor_taksasi string, id_cms_users string) structs.JsonResponse {

	response := responseStruct
	orderUfiStruct := structs.OrderUfi{}
	orderUfiCheck := structs.OrderUfi{}

	check := idb.DB.Raw("select * from order_product_ufi where id = " + id + " " +
		"and otr_taksasi is null").Scan(&orderUfiCheck).RecordNotFound()

	if check {

		otr_taksasi_c, _ := strconv.Atoi(otr_taksasi)
		id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)

		orderUfiStruct.OtrTaksasi = otr_taksasi_c
		orderUfiStruct.NomorTaksasi = nomor_taksasi
		orderUfiStruct.UpdatedBy = id_cms_users_c
		orderUfiStruct.UpdatedAt = GetTimeNow()

		err := idb.DB.Table("order_product_ufi").Where("id = " + id + "").Update(&orderUfiStruct).Error

		if err != nil {

			response.ApiMessage = errDBUpdate
		} else {

			response.ApiStatus = 1
			response.ApiMessage = succ
		}
	} else {

		response.ApiMessage = "OTR tidak dapat dirubah lagi"
	}

	return response

}

func (m *OrderModels) OrderSearch(id_cms_users string, search string) structs.JsonResponse {

	response := responseStruct
	orderListingStruct := []structs.OrderListing{}
	response.Data = orderListingStruct

	err := idb.DB.Table(order + " a").Select("a.created_at ," +
		" a.id ," +
		"a.id_contact , " +
		"b.first_name as contact_first_name , " +
		"b.last_name as contact_last_name , " +
		"c.status as order_mst_status_status , " +
		"d.plafond ," +
		"f.model").
		Joins("left join contact b " +
			"on b.id = a.id_contact ").
		Joins("left join order_mst_status c " +
			"on c.id = a.id_order_mst_status ").
		Joins("left join order_loan d " +
			"on d.id_order = a.id ").
		Joins("left join order_product_ufi e " +
			"on e.id_order = a.id ").
		Joins("left join mst_unit f " +
			"on f.id = e.id_mst_unit ").Where("a.id_cms_users = " + id_cms_users + " and concat(b.first_name , ' ',b.last_name) ilike " +
		"'%" + search + "%'").Order("b.first_name asc")

	errx := err.Find(&orderListingStruct).Error

	if errx != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = orderListingStruct
	}

	return response

}
