package models

import (
	"../structs"
	"fmt"
	"mime/multipart"
	"strconv"
)

type LeadModels struct {
}

func (m *LeadModels) LeadListing(id_cms_users string, id_lead_mst_status string, favorite string, limit string, offset string) structs.JsonResponse {

	leadListingStruct := []structs.LeadListing{}
	response := responseStruct
	response.ApiStatus = 1
	response.Data = leadListingStruct

	query := "select a.id , a.first_name , a.last_name , a.id_lead_mst_status , b.recall , " +
		"b.id_mst_log_desc ,d.id_mst_log_status ,e.status , a.favorite , a.created_at as created_at_lead , d.description ," +
		" f.id from lead a " +
		"left join lead_log b on b.id_lead = a.id and b.id_cms_users = a.id_cms_users and b.id = " +
		"(select max(c.id) from lead_log c where c.id_lead = b.id_lead and c.id_cms_users = b.id_cms_users) " +
		"left join mst_log_desc d on d.id = b.id_mst_log_desc " +
		"left join mst_log_status e on e.id = d.id_mst_log_status " +
		"left join lead_visum f on f.id_lead = a.id and f.id_cms_users = a.id_cms_users " +
		"and f.id = (select max(g.id) from lead_visum g where g.id_lead = f.id_lead and g.id_cms_users =f.id_cms_users ) " +
		"where a.id_cms_users = " + id_cms_users + ""

	if id_lead_mst_status != "" {

		query = query + " and id_lead_mst_status = " + id_lead_mst_status + ""
	}

	if favorite != "" {

		query = query + " and favorite = " + favorite + ""
	}

	err := idb.DB.Raw(query + " order by a.updated_at desc, a.created_at desc")

	if limit != "" {
		limits, _ := strconv.Atoi(limit)
		err = err.Limit(limits)
	}

	if offset != "" {
		offsets, _ := strconv.Atoi(offset)
		err = err.Offset(offsets)
	}

	err = err.Scan(&leadListingStruct)
	errx := err.Error

	if errx != nil {
		response.ApiMessage = emptyDB
	} else {
		response.ApiMessage = succ
		response.Data = leadListingStruct
	}

	return response
}
func (m *LeadModels) LeadListingVisit(types string, id_cms_users string, limit string, offset string) structs.JsonResponse {

	leadListingStruct := []structs.LeadListing{}
	response := responseStruct
	response.ApiStatus = 1
	response.Data = leadListingStruct

	null_visit := "is not null"
	if types == "new" {
		null_visit = "is null and id_lead_mst_status = 1"
	}

	err := idb.DB.Raw("select a.id , a.first_name , a.last_name , a.id_lead_mst_status , b.recall , " +
		"b.id_mst_log_desc ,d.id_mst_log_status ,e.status , a.favorite , a.created_at as created_at_lead , d.description ," +
		" f.id , f.created_at as created_at_lead_visum , f.revisit , h.status as visit_status from lead a " +
		"left join lead_log b on b.id_lead = a.id and b.id_cms_users = a.id_cms_users and b.id = " +
		"(select max(c.id) from lead_log c where c.id_lead = b.id_lead and c.id_cms_users = b.id_cms_users) " +
		"left join mst_log_desc d on d.id = b.id_mst_log_desc " +
		"left join mst_log_status e on e.id = d.id_mst_log_status " +
		"left join lead_visum f on f.id_lead = a.id and f.id_cms_users = a.id_cms_users " +
		"and f.id = (select max(g.id) from lead_visum g where g.id_lead = f.id_lead and g.id_cms_users =f.id_cms_users ) " +
		"left join mst_visum_status h " +
		"on h.id = f.id_mst_visum_status " +
		"where a.id_cms_users = " + id_cms_users + " and f.id " + null_visit + " and id_lead_mst_status <> 6 order by a.updated_at desc, a.created_at desc")

	if limit != "" {
		limits, _ := strconv.Atoi(limit)
		err = err.Limit(limits)
	}

	if offset != "" {
		offsets, _ := strconv.Atoi(offset)
		err = err.Offset(offsets)
	}

	err = err.Scan(&leadListingStruct)
	errx := err.Error

	if errx != nil {
		response.ApiMessage = emptyDB
	} else {
		response.ApiMessage = succ
		response.Data = leadListingStruct
	}

	return response
}

func (m *LeadModels) LeadUpdateFav(id string, favorite string, updated_by string) structs.JsonResponse {

	response := responseStruct

	rows := idb.DB.Exec("update lead set favorite = " + favorite + " , updated_by = " + updated_by + " , updated_at = now() where id = " + id + "").RowsAffected

	if rows >= 1 {

		response.ApiStatus = 1
		response.ApiMessage = succ
	} else {
		response.ApiMessage = errDBUpdate
	}

	return response
}
func (m *LeadModels) LeadUpdateHapus(id string, updated_by string) structs.JsonResponse {

	response := responseStruct

	rows := idb.DB.Exec("update lead set id_lead_mst_status = 6 , updated_by = " + updated_by + " , updated_at = now() where id = " + id + "").RowsAffected

	if rows >= 1 {

		response.ApiStatus = 1
		response.ApiMessage = succ
	} else {
		response.ApiMessage = errDBUpdate
	}

	return response
}

func (m *LeadModels) DetailLeadTotal(id string, id_cms_users string) structs.JsonResponse {

	detailLeadTotalStruct := structs.DetailLeadTotal{}
	response := responseStruct
	response.Data = detailLeadTotalStruct

	err := idb.DB.Raw("select (select count(id) from lead_log where id_lead = " + id + " and id_cms_users = " + id_cms_users + ") as log_total , " +
		"(select count(id) from lead_visum where id_lead = " + id + " and id_cms_users = " + id_cms_users + ") as visum_total , " +
		"(select count(id) from lead_note where id_lead = " + id + " and id_cms_users = " + id_cms_users + ") as note_total").
		Scan(&detailLeadTotalStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = detailLeadTotalStruct
	}

	return response
}

func (m *LeadModels) DetailLead(id string) structs.JsonResponse {

	detailLeadStruct := structs.DetailLead{}
	leadPhoneStruct := []structs.LeadPhone{}
	addressLeadStruct := []structs.LeadAddress{}
	unitLeadStruct := []structs.LeadUnit{}

	response := responseStruct
	response.Data = detailLeadStruct

	err := idb.DB.Raw("select a.id ,a.id_mst_outlet , b.outlet_name as mst_outlet_outlet_name ," +
		"a.id_mst_data_source ,c.datasource as mst_data_source_datasource , a.first_name , " +
		"a.last_name , a.id_mst_job , d.job as mst_job_job , a.id_lead_mst_status , " +
		"e.status as lead_mst_status_status , f.name as cms_users_name ," +
		"(select count(id) from lead_log where id_lead = a.id and id_cms_users = a.id_cms_users) as log_total , " +
		"g.id_mst_log_desc ,h.description , h.id_mst_log_status," +
		"(select count(id) from lead_visum where id_lead = a.id and id_cms_users = a.id_cms_users) as visum_total , " +
		"(select count(id) from lead_note where id_lead = a.id and id_cms_users = a.id_cms_users) as note_total " +
		"from lead a " +
		"left join mst_outlet b on b.id = a.id_mst_outlet " +
		"left join mst_data_source c " +
		"on c.id = a.id_mst_data_source " +
		"left join mst_job d " +
		"on d.id = a.id_mst_job " +
		"left join lead_mst_status e " +
		"on e.id = a.id_lead_mst_status " +
		"left join cms_users f " +
		"on f.id = a.id_cms_users " +
		"left join lead_log g " +
		"on g.id_lead = a.id and g.id_cms_users = a.id_cms_users and g.id = " +
		"(select max(h.id) from lead_log h where h.id_lead = g.id_lead and h.id_cms_users = g.id_cms_users) " +
		"left join mst_log_desc h " +
		"on h.id = g.id_mst_log_desc " +
		"left join mst_log_status i on i.id = h.id_mst_log_status " +
		"where a.id = " + id + "").Scan(&detailLeadStruct).Error

	idb.DB.Raw("select * from lead_phone where id_lead = " + id + "").Scan(&leadPhoneStruct)
	idb.DB.Raw("select * from lead_address a " +
		"left join mst_address b " +
		"on b.id = a.id_mst_address " +
		"where a.id_lead = " + id + "").Scan(&addressLeadStruct)
	idb.DB.Raw("select * , b.nama as mst_product_nama , b.status as mst_product_status , " +
		"c.id_lead_product as id_lead_product_detail from lead_product a " +
		"left join mst_product b " +
		"on b.id = a.id_mst_product " +
		"left join lead_product_detail c " +
		"on c.id_lead_product = a.id " +
		"left join mst_unit d " +
		"on d.id = c.id_mst_unit " +
		"where a.id_lead = " + id + "").Scan(&unitLeadStruct)

	if err != nil {
		response.ApiMessage = errDB
	} else {
		detailLeadStruct.LeadAddress = addressLeadStruct
		detailLeadStruct.LeadPhone = leadPhoneStruct
		detailLeadStruct.LeadUnit = unitLeadStruct

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = detailLeadStruct
	}

	return response
}

func (m *LeadModels) LeadLog(id string, id_cms_users string) structs.JsonResponse {

	leadLogStruct := []structs.LeadLog{}
	response := responseStruct
	response.Data = leadLogStruct

	err := idb.DB.Raw("select a.id , b.first_name as lead_first_name , b.last_name as lead_last_name , " +
		"a.duration , a.id_mst_log_desc , c.description as mst_log_desc_description , a.recall , e.name as cms_users_name , " +
		"a.created_at , c.id_mst_log_status , d.status from lead_log a left join lead b " +
		"on b.id = a.id_lead " +
		"left join mst_log_desc c " +
		"on c.id = a.id_mst_log_desc " +
		"left join mst_log_status d " +
		"on d.id = c.id_mst_log_status " +
		"left join cms_users e " +
		"on e.id = a.id_cms_users " +
		"where a.id_lead = " + id + " and a.id_cms_users = " + id_cms_users + " order by a.created_at desc ").Scan(&leadLogStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiMessage = succ
		response.Data = leadLogStruct
	}

	return response
}

func (m *LeadModels) LeadVisum(id string, id_cms_users string) structs.JsonResponse {

	leadVisumStruct := []structs.LeadVisum{}
	response := responseStruct
	response.Data = leadVisumStruct

	err := idb.DB.Raw("select a.id , b.first_name as lead_first_name , b.last_name as lead_last_name , " +
		"e.status as mst_visum_status_status , a.revisit , a.id_cms_users , c.name as cms_users_name , a.created_at , " +
		"d.photo from lead_visum a " +
		"left join lead b " +
		"on b.id = a.id_lead " +
		"left join cms_users c " +
		"on c.id = a.id_cms_users " +
		"left join lead_photo d " +
		"on d.id_lead_visum = a.id " +
		"left join mst_visum_status e " +
		"on e.id = a.id_mst_visum_status " +
		"where a.id_lead = " + id + " and a.id_cms_users = " + id_cms_users + "").Scan(&leadVisumStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiMessage = succ
		response.Data = leadVisumStruct
	}

	return response
}

func (m *LeadModels) LeadNote(id string, id_cms_users string) structs.JsonResponse {

	leadNoteStruct := []structs.LeadNote{}
	response := responseStruct

	err := idb.DB.Raw("select a.id , a.note , b.name as cms_users_name , a.created_at from lead_note a " +
		"left join cms_users b " +
		"on b.id = a.id_cms_users where a.id_lead = " + id + " and a.id_cms_users = " + id_cms_users + " order by a.created_at desc").Scan(&leadNoteStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiMessage = succ
		response.Data = leadNoteStruct
	}

	return response
}

func (m *LeadModels) LeadLogCreate(id_lead string, duration string, id_mst_log_desc string, recall string,
	id_cms_users string, id_modul string, id_data string, jenis string) structs.JsonResponse {

	leadLogCreateStruct := structs.LeadLogCreate{}
	mstLogsStruct := structs.MstLogs{}
	leadReadStruct := structs.DetailLead{}
	mstLogDescStruct := structs.MstLogDesc{}

	response := responseStruct

	id_lead_c, _ := strconv.ParseInt(id_lead, 10, 64)
	duration_c, _ := strconv.Atoi(duration)
	id_mst_log_desc_c, _ := strconv.ParseInt(id_mst_log_desc, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
	id_modul_c, _ := strconv.ParseInt(id_modul, 10, 64)
	id_data_c, _ := strconv.ParseInt(id_data, 10, 64)

	leadLogCreateStruct.IdCmsUsers = id_cms_users_c
	leadLogCreateStruct.Duration = duration_c
	leadLogCreateStruct.Recall = recall
	leadLogCreateStruct.IdLead = id_lead_c
	leadLogCreateStruct.IdMstLogDesc = id_mst_log_desc_c
	leadLogCreateStruct.CreatedAt = GetTimeNow()

	errLeadLogCreate := idb.DB.Table("lead_log").Create(&leadLogCreateStruct).Error

	if errLeadLogCreate != nil {

		response.ApiMessage = errDBAdd

	} else {

		mstLogsStruct.IdCmsUsers = id_cms_users_c
		mstLogsStruct.IdData = id_data_c
		mstLogsStruct.IdModul = id_modul_c
		mstLogsStruct.Jenis = jenis
		mstLogsStruct.CreatedAt = GetTimeNow()

		_ = idb.DB.Table("mst_logs").Create(&mstLogsStruct).Error

		_ = idb.DB.Table("lead").Where("id = " + id_data + "").First(&leadReadStruct).Error

		id_lead_status := leadReadStruct.IdLeadMstStatus
		id_lead = id_data
		id_status := 0

		_ = idb.DB.Raw("select * from mst_log_desc where id = " + id_mst_log_desc + "").Scan(&mstLogDescStruct).Error

		id_log_status := mstLogDescStruct.IdMstLogStatus

		if id_lead_status == 1 && id_log_status == 1 {
			id_status = 2
		} else if id_lead_status == 1 && id_log_status == 2 {
			id_status = 4
		} else if id_lead_status == 2 && id_log_status == 1 {
			id_status = 2
		} else if id_lead_status == 2 && id_log_status == 2 {
			id_status = 4
		} else if id_lead_status == 3 && id_log_status == 1 {
			id_status = 2
		} else if id_lead_status == 3 && id_log_status == 2 {
			id_status = 4
		} else if id_lead_status == 4 && id_log_status == 1 {
			id_status = 2
		} else if id_lead_status == 4 && id_log_status == 2 {
			id_status = 4
		} else {
			id_status = 3
		}

		_ = idb.DB.Exec("update lead set updated_at = now() , " +
			"id_lead_mst_status = " + fmt.Sprint(id_status) + " , updated_by = " + id_cms_users + " " +
			"where id = " + id_lead + "").RowsAffected

		response.ApiStatus = 1
		response.ApiMessage = succ

	}

	return response
}

func (m *LeadModels) LeadNoteCreate(id_lead string, id_cms_users string, note string) structs.JsonResponse {

	response := responseStruct
	leadNoteCreateStruct := structs.LeadNoteCreate{}

	id_lead_c, _ := strconv.ParseInt(id_lead, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)

	leadNoteCreateStruct.IdLead = id_lead_c
	leadNoteCreateStruct.IdCmsUsers = id_cms_users_c
	leadNoteCreateStruct.Note = note
	leadNoteCreateStruct.CreatedAt = GetTimeNow()

	err := idb.DB.Table("lead_note").Create(&leadNoteCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd

	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
	}

	return response
}
func (m *LeadModels) LeadVisumCreate(id_lead string, id_cms_users string, revisit string, id_mst_visum_status string, files multipart.File, header *multipart.FileHeader) structs.JsonResponse {

	response := responseStruct
	leadVisumCreateStruct := structs.LeadVisumCreate{}

	id_lead_c, _ := strconv.ParseInt(id_lead, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
	id_mst_visum_status_c, _ := strconv.ParseInt(id_mst_visum_status, 10, 64)

	leadVisumCreateStruct.IdLead = id_lead_c
	leadVisumCreateStruct.IdCmsUsers = id_cms_users_c
	leadVisumCreateStruct.Revisit = revisit
	leadVisumCreateStruct.IdMstVisumStatus = id_mst_visum_status_c
	leadVisumCreateStruct.CreatedAt = GetTimeNow()

	err := idb.DB.Table("lead_visum").Create(&leadVisumCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd

	} else {

		uploadImageVisum := func() {

			url := UploadImage("lead", fmt.Sprint(leadVisumCreateStruct.Id), files, header)

			if url != "" {

				leadVisumPhotoCreate := structs.LeadVisumPhotoCreate{}

				leadVisumPhotoCreate.IdCmsUsers = id_cms_users_c
				leadVisumPhotoCreate.IdLead = id_lead_c
				leadVisumPhotoCreate.IdLeadVisum = leadVisumCreateStruct.Id
				leadVisumPhotoCreate.Photo = url
				leadVisumPhotoCreate.CreatedAt = GetTimeNow()

				_ = idb.DB.Table("lead_photo").Create(&leadVisumPhotoCreate).Error

			}
		}
		logCreate := func() {

			mstLogsStruct := structs.MstLogs{}

			mstLogsStruct.IdCmsUsers = id_cms_users_c
			mstLogsStruct.IdData = id_lead_c
			mstLogsStruct.IdModul = 7
			mstLogsStruct.Jenis = "visit"
			mstLogsStruct.CreatedAt = GetTimeNow()

			_ = idb.DB.Table("mst_logs").Create(&mstLogsStruct).Error

		}

		go uploadImageVisum()
		go logCreate()

		response.ApiStatus = 1
		response.ApiMessage = succ
	}

	return response
}
func (m *LeadModels) LeadPhoneCheckNewDB(number string) structs.JsonResponse {

	response := responseStruct
	leadInfoStruct := structs.LeadCreate{}

	rows := idb.DB.Raw("select b.id , b.first_name , b.last_name from lead_phone a " +
		"left join lead b " +
		"on b.id = a.id_lead " +
		"where a.number = '" + number + "'").Scan(&leadInfoStruct).RecordNotFound()

	if rows {

		response.ApiStatus = 0
		response.ApiMessage = succ
	} else {
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = leadInfoStruct
	}

	return response
}
func (m *LeadModels) LeadCreate(id_mst_outlet string, id_mst_data_source string, id_cms_users string,
	first_name string, last_name string, id_mst_job string, id_lead_mst_status string, number string,
	id_mst_address string, id_mst_category_address string, id_mst_product string, id_mst_unit string,
	nopol string, tax_status string, owner string, address string, address_bool bool, product_bool bool) structs.JsonResponse {

	leadCreateStruct := structs.LeadCreate{}
	leadPhoneCreateStruct := structs.LeadPhone{}
	leadAddressCreateStruct := structs.LeadAddressCreate{}
	leadProductCreateStruct := structs.LeadProductCreate{}
	leadProductDetailCreate := structs.LeadProductDetailCreate{}

	response := responseStruct
	response.Data = leadCreateStruct

	id_mst_outlet_c, _ := strconv.ParseInt(id_mst_outlet, 10, 64)
	id_mst_data_source_c, _ := strconv.ParseInt(id_mst_data_source, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
	id_mst_job_c, _ := strconv.ParseInt(id_mst_job, 10, 64)
	id_lead_mst_status_c, _ := strconv.ParseInt(id_lead_mst_status, 10, 64)
	id_mst_address_c, _ := strconv.ParseInt(id_mst_address, 10, 64)
	id_mst_category_address_c, _ := strconv.ParseInt(id_mst_category_address, 10, 64)
	id_mst_product_c, _ := strconv.ParseInt(id_mst_product, 10, 64)
	id_mst_unit_c, _ := strconv.ParseInt(id_mst_unit, 10, 64)

	leadCreateStruct.IdCmsUsers = id_cms_users_c
	leadCreateStruct.IdLeadMstStatus = id_lead_mst_status_c
	leadCreateStruct.LastName = last_name
	leadCreateStruct.FirstName = first_name
	leadCreateStruct.IdMstDataSource = id_mst_data_source_c
	leadCreateStruct.IdMstJob = id_mst_job_c
	leadCreateStruct.IdMstOutlet = id_mst_outlet_c
	leadCreateStruct.CreatedAt = GetTimeNow()

	err := idb.DB.Table("lead").Create(&leadCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd
	} else {

		createOther := func() {

			leadPhoneCreateStruct.IdCmsUsers = id_cms_users_c
			leadPhoneCreateStruct.IdLead = leadCreateStruct.Id
			leadPhoneCreateStruct.Status = "Y"
			leadPhoneCreateStruct.Number = number
			leadPhoneCreateStruct.CreatedAt = GetTimeNow()

			_ = idb.DB.Table("lead_phone").Create(&leadPhoneCreateStruct).Error

			if address_bool {

				leadAddressCreateStruct.IdLead = leadCreateStruct.Id
				leadAddressCreateStruct.Address = address
				leadAddressCreateStruct.IdCmsUsers = id_cms_users_c
				leadAddressCreateStruct.IdMstAddress = id_mst_address_c
				leadAddressCreateStruct.IdMstCategoryAddress = id_mst_category_address_c
				leadAddressCreateStruct.CreatedAt = GetTimeNow()

				_ = idb.DB.Table("lead_address").Create(&leadAddressCreateStruct).Error
			}

			if product_bool {

				leadProductCreateStruct.IdCmsUsers = id_cms_users_c
				leadProductCreateStruct.IdLead = leadCreateStruct.Id
				leadProductCreateStruct.IdMstProduct = id_mst_product_c
				leadProductCreateStruct.CreatedAt = GetTimeNow()

				_ = idb.DB.Table("lead_product").Create(&leadProductCreateStruct).Error

				leadProductDetailCreate.IdCmsUsers = id_cms_users_c
				leadProductDetailCreate.Nopol = nopol
				leadProductDetailCreate.IdLeadProduct = leadProductCreateStruct.Id
				leadProductDetailCreate.IdMstUnit = id_mst_unit_c
				leadProductDetailCreate.Owner = owner
				leadProductDetailCreate.TaxStatus = tax_status
				leadProductDetailCreate.CreatedAt = GetTimeNow()

				_ = idb.DB.Table("lead_product_detail").Create(&leadProductDetailCreate)
			}

			logCreate := func() {

				mstLogsStruct := structs.MstLogs{}

				mstLogsStruct.IdCmsUsers = id_cms_users_c
				mstLogsStruct.IdData = leadCreateStruct.Id
				mstLogsStruct.IdModul = 2
				mstLogsStruct.Jenis = "create"
				mstLogsStruct.CreatedAt = GetTimeNow()

				_ = idb.DB.Table("mst_logs").Create(&mstLogsStruct).Error

			}

			logCreate()

		}

		go createOther()

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = leadCreateStruct

	}

	return response
}

func (m *LeadModels) LeadUpdate(id_lead string, id_mst_outlet string, id_mst_data_source string, id_cms_users string,
	first_name string, last_name string, id_mst_job string, id_lead_mst_status string, number string,
	id_mst_address string, id_mst_category_address string, id_mst_product string, id_mst_unit string,
	nopol string, tax_status string, owner string, address string, address_bool bool, product_bool bool) structs.JsonResponse {

	leadCreateStruct := structs.LeadCreate{}
	leadPhoneCreateStruct := structs.LeadPhone{}
	leadAddressCreateStruct := structs.LeadAddressCreate{}
	leadProductCreateStruct := structs.LeadProductCreate{}
	leadProductDetailCreate := structs.LeadProductDetailCreate{}

	response := responseStruct
	response.Data = leadCreateStruct

	id_lead_c, _ := strconv.ParseInt(id_lead, 10, 64)
	id_mst_outlet_c, _ := strconv.ParseInt(id_mst_outlet, 10, 64)
	id_mst_data_source_c, _ := strconv.ParseInt(id_mst_data_source, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
	id_mst_job_c, _ := strconv.ParseInt(id_mst_job, 10, 64)
	id_lead_mst_status_c, _ := strconv.ParseInt(id_lead_mst_status, 10, 64)
	id_mst_address_c, _ := strconv.ParseInt(id_mst_address, 10, 64)
	id_mst_category_address_c, _ := strconv.ParseInt(id_mst_category_address, 10, 64)
	id_mst_product_c, _ := strconv.ParseInt(id_mst_product, 10, 64)
	id_mst_unit_c, _ := strconv.ParseInt(id_mst_unit, 10, 64)

	leadCreateStruct.IdCmsUsers = id_cms_users_c
	leadCreateStruct.IdLeadMstStatus = id_lead_mst_status_c
	leadCreateStruct.LastName = last_name
	leadCreateStruct.FirstName = first_name
	leadCreateStruct.IdMstDataSource = id_mst_data_source_c
	leadCreateStruct.IdMstJob = id_mst_job_c
	leadCreateStruct.IdMstOutlet = id_mst_outlet_c
	leadCreateStruct.UpdatedAt = GetTimeNow()
	leadCreateStruct.UpdatedBy = id_cms_users_c

	err := idb.DB.Table("lead").Where("id = " + id_lead + "").Update(&leadCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBUpdate
	} else {

		updateOther := func() {

			leadPhoneCreateStruct.IdCmsUsers = id_cms_users_c
			leadPhoneCreateStruct.Status = "Y"
			leadPhoneCreateStruct.Number = number

			_ = idb.DB.Table("lead_phone").Where("id_lead = " + id_lead + "").Update(&leadPhoneCreateStruct).Error

			if address_bool {

				leadAddressCheck := leadAddressCreateStruct

				rows := idb.DB.Raw("select * from lead_address where id_lead = " + id_lead + "").Scan(&leadAddressCheck).RecordNotFound()

				if rows {

					leadAddressCreateStruct.IdLead = id_lead_c
					leadAddressCreateStruct.Address = address
					leadAddressCreateStruct.IdCmsUsers = id_cms_users_c
					leadAddressCreateStruct.IdMstAddress = id_mst_address_c
					leadAddressCreateStruct.IdMstCategoryAddress = id_mst_category_address_c
					leadAddressCreateStruct.CreatedAt = GetTimeNow()

					_ = idb.DB.Table("lead_address").Create(&leadAddressCreateStruct).Error
				} else {

					leadAddressCreateStruct.IdLead = id_lead_c
					leadAddressCreateStruct.Address = address
					leadAddressCreateStruct.IdCmsUsers = id_cms_users_c
					leadAddressCreateStruct.IdMstAddress = id_mst_address_c
					leadAddressCreateStruct.IdMstCategoryAddress = id_mst_category_address_c
					leadAddressCreateStruct.UpdatedAt = GetTimeNow()
					leadAddressCreateStruct.UpdatedBy = id_cms_users_c

					_ = idb.DB.Table("lead_address").Where("id = " + fmt.Sprint(leadAddressCheck.Id) + "").Update(&leadAddressCreateStruct).Error

				}
			}

			if product_bool {

				leadProductCheck := leadProductCreateStruct

				rows := idb.DB.Raw("select * from lead_product where id_lead = " + id_lead + "").Scan(&leadProductCheck).RecordNotFound()

				if rows {

					leadProductCreateStruct.IdCmsUsers = id_cms_users_c
					leadProductCreateStruct.IdLead = id_lead_c
					leadProductCreateStruct.IdMstProduct = id_mst_product_c
					leadProductCreateStruct.CreatedAt = GetTimeNow()

					_ = idb.DB.Table("lead_product").Create(&leadProductCreateStruct).Error

					leadProductDetailCreate.IdCmsUsers = id_cms_users_c
					leadProductDetailCreate.Nopol = nopol
					leadProductDetailCreate.IdLeadProduct = leadProductCreateStruct.Id
					leadProductDetailCreate.IdMstUnit = id_mst_unit_c
					leadProductDetailCreate.Owner = owner
					leadProductDetailCreate.TaxStatus = tax_status
					leadProductDetailCreate.CreatedAt = GetTimeNow()

					_ = idb.DB.Table("lead_product_detail").Create(&leadProductDetailCreate)

				} else {

					leadProductCreateStruct.IdCmsUsers = id_cms_users_c
					leadProductCreateStruct.IdLead = leadCreateStruct.Id
					leadProductCreateStruct.IdMstProduct = id_mst_product_c
					leadProductCreateStruct.UpdatedAt = GetTimeNow()
					leadProductCreateStruct.UpdatedBy = id_cms_users_c

					_ = idb.DB.Table("lead_product").Where("id = " + fmt.Sprint(leadProductCheck.Id) + "").Update(&leadProductCreateStruct).Error

					leadProductDetailCreate.IdCmsUsers = id_cms_users_c
					leadProductDetailCreate.Nopol = nopol
					//leadProductDetailCreate.IdLeadProduct = leadProductCreateStruct.Id
					leadProductDetailCreate.IdMstUnit = id_mst_unit_c
					leadProductDetailCreate.Owner = owner
					leadProductDetailCreate.TaxStatus = tax_status
					leadProductDetailCreate.UpdatedAt = GetTimeNow()
					leadProductDetailCreate.UpdatedBy = id_cms_users_c

					_ = idb.DB.Table("lead_product_detail").Where("id_lead_product = " + fmt.Sprint(leadProductCheck.Id) + "").Update(&leadProductDetailCreate)

				}

			}

		}

		go updateOther()

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = leadCreateStruct

	}

	return response
}

func (m *LeadModels) LeadCollabCreate(id_lead string, id_cms_users string) structs.JsonResponse {

	response := responseStruct
	leadColabStruct := structs.LeadColab{}
	leadColabCreateStruct := structs.LeadColab{}

	checkColab := idb.DB.Raw("select * from lead_collaborate where id_lead = " + id_lead + " " +
		"and id_cms_users = " + id_cms_users + "").Scan(&leadColabStruct).RecordNotFound()

	if checkColab {

		id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
		id_lead_c, _ := strconv.ParseInt(id_lead, 10, 64)

		leadColabCreateStruct.IdCmsUsers = id_cms_users_c
		leadColabCreateStruct.IdLead = id_lead_c

		err := idb.DB.Table("lead_collaborate").Create(&leadColabCreateStruct).Error

		if err != nil {

			response.ApiStatus = 0
			response.ApiMessage = errDBAdd

		} else {

			response.ApiStatus = 2
			response.ApiMessage = succ
		}

	} else {

		response.ApiStatus = 1
		response.ApiMessage = "Lead Colab ditemukan"
	}

	return response
}

func (m *LeadModels) LeadCollabList(id_cms_users string, limit string, offset string) structs.JsonResponse {

	response := responseStruct
	leadCollabListStruct := []structs.LeadCollabListing{}
	response.Data = leadCollabListStruct

	err := idb.DB.Raw("select a.id , a.id_lead , b.first_name as lead_first_name , b.last_name as lead_last_name , " +
		"c.name as cms_users_name , d.recall ,d.id_mst_log_desc ,f.id_mst_log_status , f.description , g.status " +
		"from lead_collaborate a " +
		"left join lead b " +
		"on b.id = a.id_lead " +
		"left join cms_users c " +
		"on c.id = a.id_cms_users " +
		"left join lead_log d " +
		"on d.id_lead = a.id_lead and d.id = " +
		"(select max(e.id) from lead_log e where e.id_lead = d.id_lead) " +
		"left join mst_log_desc f " +
		"on f.id = d.id_mst_log_desc " +
		"left join mst_log_status g " +
		"on g.id = f.id_mst_log_status " +
		"where a.id_cms_users = " + id_cms_users + " and b.id_lead_mst_status <> 6 order by b.updated_at desc , b.created_at desc").Scan(&leadCollabListStruct).Error

	if err != nil {

		response.ApiMessage = errDB

	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = leadCollabListStruct
	}

	return response
}

func (m *LeadModels) LeadSearch(id_cms_users string, search string) structs.JsonResponse {

	response := responseStruct
	leadSearch := []structs.LeadSearch{}

	err := idb.DB.Raw("select a.id , a.id_cms_users, a.first_name , a.last_name , d.recall ,d.id_mst_log_desc ," +
		"f.id_mst_log_status , f.description , g.status ,a.favorite , a.created_at as created_at_lead from lead a " +
		"left join lead_log d " +
		"on d.id_lead = a.id and d.id = " +
		"(select max(e.id) from lead_log e where e.id_lead = d.id_lead) " +
		"left join mst_log_desc f " +
		"on f.id = d.id_mst_log_desc " +
		"left join mst_log_status g " +
		"on g.id = f.id_mst_log_status " +
		"left join lead_phone h " +
		"on h.id_lead = a.id and h.id = (select max(i.id) from lead_phone i where i.id_lead = h.id_lead) " +
		"where concat(a.first_name,' ',a.last_name,' ',h.number) ilike '%" + search + "%' and a.id_cms_users = " + id_cms_users + " " +
		"order by a.updated_at desc , d.created_at desc , a.first_name asc").Scan(&leadSearch).Error

	if err != nil {
		response.ApiStatus = 1

		response.ApiMessage = emptyDB
	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = leadSearch
	}

	return response
}
