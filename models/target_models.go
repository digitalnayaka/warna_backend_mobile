package models

import (
	"../structs"
	"fmt"
	"mime/multipart"
	"strconv"
)

type TargetModels struct {
}

func (m *TargetModels) TargetListing(id_cms_users string, id_target_mst_status string, favorite string, limit string, offset string) structs.JsonResponse {

	targetListingStruct := []structs.TargetListing{}
	response := responseStruct
	response.ApiStatus = 1
	response.Data = targetListingStruct

	query := "select a.id , a.first_name , a.last_name , a.id_target_mst_status , b.recall , " +
		"b.id_mst_log_desc ,d.id_mst_log_status ,e.status ,  a.created_at as created_at_target , d.description ," +
		" f.id , " +
		"a.category , f.id_mst_visum_status , f.revisit , h.status as visit_status" +
		" from target a " +
		"left join target_log b on b.id_target = a.id and b.id_cms_users = a.id_cms_users and b.id = " +
		"(select max(c.id) from target_log c where c.id_target = b.id_target and c.id_cms_users = b.id_cms_users) " +
		"left join mst_log_desc d on d.id = b.id_mst_log_desc " +
		"left join mst_log_status e on e.id = d.id_mst_log_status " +
		"left join target_visum f on f.id_target = a.id and f.id_cms_users = a.id_cms_users " +
		"and f.id = (select max(g.id) " +
		"from target_visum g where g.id_target = f.id_target and g.id_cms_users =f.id_cms_users ) " +
		"left join mst_visum_status h on h.id = f.id_mst_visum_status " +
		"where a.id_cms_users = " + id_cms_users + ""

	if id_target_mst_status != "" {

		query = query + " and id_target_mst_status = " + id_target_mst_status + ""
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

	err = err.Scan(&targetListingStruct)
	errx := err.Error

	if errx != nil {
		response.ApiMessage = emptyDB
	} else {
		response.ApiMessage = succ
		response.Data = targetListingStruct
	}

	return response
}
func (m *TargetModels) TargetListingVisit(types string, id_cms_users string, limit string, offset string) structs.JsonResponse {

	targetListingStruct := []structs.TargetListing{}
	response := responseStruct
	response.ApiStatus = 1
	response.Data = targetListingStruct

	null_visit := "is not null"
	order_by := "f.created_at desc"
	if types == "new" {
		null_visit = "is null and id_target_mst_status = 1"
		order_by = "a.created_at desc"
	}

	err := idb.DB.Raw("select a.id , a.first_name , a.last_name , a.id_target_mst_status , b.recall , " +
		"b.id_mst_log_desc ,d.id_mst_log_status ,e.status , a.created_at as created_at_target , d.description ," +
		" f.id , f.created_at as created_at_target_visum , f.revisit , a.category , h.status as visit_status from target a " +
		"left join target_log b on b.id_target = a.id and b.id_cms_users = a.id_cms_users and b.id = " +
		"(select max(c.id) from target_log c where c.id_target = b.id_target and c.id_cms_users = b.id_cms_users) " +
		"left join mst_log_desc d on d.id = b.id_mst_log_desc " +
		"left join mst_log_status e on e.id = d.id_mst_log_status " +
		"left join target_visum f on f.id_target = a.id and f.id_cms_users = a.id_cms_users " +
		"and f.id = (select max(g.id) from target_visum g where g.id_target = f.id_target and g.id_cms_users =f.id_cms_users ) " +
		"left join mst_visum_status h " +
		"on h.id = f.id_mst_visum_status " +
		"where a.id_cms_users = " + id_cms_users + " and f.id " + null_visit + " order by " + order_by + "")

	if limit != "" {
		limits, _ := strconv.Atoi(limit)
		err = err.Limit(limits)
	}

	if offset != "" {
		offsets, _ := strconv.Atoi(offset)
		err = err.Offset(offsets)
	}

	err = err.Scan(&targetListingStruct)
	errx := err.Error

	if errx != nil {
		response.ApiMessage = emptyDB
	} else {
		response.ApiMessage = succ
		response.Data = targetListingStruct
	}

	return response
}

func (m *TargetModels) DetailTargetTotal(id string, id_cms_users string) structs.JsonResponse {

	detailTargetTotalStruct := structs.DetailTargetTotal{}
	response := responseStruct
	response.Data = detailTargetTotalStruct

	err := idb.DB.Raw("select (select count(id) from target_log where id_target = " + id + " and id_cms_users = " + id_cms_users + ") as log_total , " +
		"(select count(id) from target_visum where id_target = " + id + " and id_cms_users = " + id_cms_users + ") as visum_total , " +
		"(select count(id) from target_note where id_target = " + id + " and id_cms_users = " + id_cms_users + ") as note_total").
		Scan(&detailTargetTotalStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = detailTargetTotalStruct
	}

	return response
}

func (m *TargetModels) DetailTarget(id string) structs.JsonResponse {

	detailTargetStruct := structs.DetailTarget{}

	response := responseStruct
	response.Data = detailTargetStruct

	err := idb.DB.Raw("select " +
		"a.id , " +
		"a.hp_1 ," +
		"a.hp_2 ," +
		"a.provider_1 ," +
		"a.provider_2 ," +
		"a.nopol ," +
		"a.business_code , " +
		"a.category , " +
		"a.priority , " +
		"a.no_contract , " +
		"a.id_mst_data_source ," +
		"c.datasource as mst_data_source_datasource , " +
		"a.first_name , " +
		"a.last_name , " +
		"a.job , " +
		"a.address , " +
		"a.kelurahan ," +
		"a.kecamatan , " +
		"a.kabupaten , " +
		"a.provinsi , " +
		"a.id_target_mst_status , " +
		"e.status as target_mst_status_status , " +
		"f.name as cms_users_name ," +
		"(select count(id) from target_log where id_target = a.id and id_cms_users = a.id_cms_users) as log_total , " +
		"g.id_mst_log_desc ,h.description , h.id_mst_log_status," +
		"(select count(id) from target_visum where id_target = a.id and id_cms_users = a.id_cms_users) as visum_total , " +
		"(select count(id) from target_note where id_target = a.id and id_cms_users = a.id_cms_users) as note_total " +
		"from target a " +
		"left join mst_data_source c " +
		"on c.id = a.id_mst_data_source " +
		"left join target_mst_status e " +
		"on e.id = a.id_target_mst_status " +
		"left join cms_users f " +
		"on f.id = a.id_cms_users " +
		"left join target_log g " +
		"on g.id_target = a.id and g.id_cms_users = a.id_cms_users and g.id = " +
		"(select max(h.id) from target_log h where h.id_target = g.id_target and h.id_cms_users = g.id_cms_users) " +
		"left join mst_log_desc h " +
		"on h.id = g.id_mst_log_desc " +
		"left join mst_log_status i on i.id = h.id_mst_log_status " +
		"where a.id = " + id + "").Scan(&detailTargetStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = detailTargetStruct
	}

	return response
}

func (m *TargetModels) TargetLog(id string, id_cms_users string) structs.JsonResponse {

	targetLogStruct := []structs.TargetLog{}
	response := responseStruct
	response.Data = targetLogStruct

	err := idb.DB.Raw("select a.id , b.first_name as target_first_name , b.last_name as target_last_name , " +
		"a.duration , a.id_mst_log_desc , c.description as mst_log_desc_description , a.recall , e.name as cms_users_name , " +
		"a.created_at , c.id_mst_log_status , d.status from target_log a left join target b " +
		"on b.id = a.id_target " +
		"left join mst_log_desc c " +
		"on c.id = a.id_mst_log_desc " +
		"left join mst_log_status d " +
		"on d.id = c.id_mst_log_status " +
		"left join cms_users e " +
		"on e.id = a.id_cms_users " +
		"where a.id_target = " + id + " and a.id_cms_users = " + id_cms_users + " order by a.created_at desc").Scan(&targetLogStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiMessage = succ
		response.Data = targetLogStruct
	}

	return response
}

func (m *TargetModels) TargetVisum(id string, id_cms_users string) structs.JsonResponse {

	targetVisumStruct := []structs.TargetVisum{}
	response := responseStruct
	response.Data = targetVisumStruct

	err := idb.DB.Raw("select a.id , b.first_name as target_first_name , b.last_name as target_last_name , " +
		"e.status as mst_visum_status_status , a.revisit , a.id_cms_users , c.name as cms_users_name , a.created_at , " +
		"d.photo from target_visum a " +
		"left join target b " +
		"on b.id = a.id_target " +
		"left join cms_users c " +
		"on c.id = a.id_cms_users " +
		"left join target_photo d " +
		"on d.id_target_visum = a.id " +
		"left join mst_visum_status e " +
		"on e.id = a.id_mst_visum_status " +
		"where a.id_target = " + id + " and a.id_cms_users = " + id_cms_users + " order by a.created_at desc ").Scan(&targetVisumStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiMessage = succ
		response.Data = targetVisumStruct
	}

	return response
}

func (m *TargetModels) TargetNote(id string, id_cms_users string) structs.JsonResponse {

	targetNoteStruct := []structs.TargetNote{}
	response := responseStruct

	err := idb.DB.Raw("select a.id , a.note , b.name as cms_users_name , a.created_at from target_note a " +
		"left join cms_users b " +
		"on b.id = a.id_cms_users where a.id_target = " + id + " and a.id_cms_users = " + id_cms_users + " order by a.created_at desc").Scan(&targetNoteStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {

		response.ApiMessage = succ
		response.Data = targetNoteStruct
	}

	return response
}

func (m *TargetModels) TargetLogCreate(id_target string, duration string, id_mst_log_desc string, recall string,
	id_cms_users string, id_modul string, id_data string, jenis string) structs.JsonResponse {

	targetLogCreateStruct := structs.TargetLogCreate{}
	mstLogsStruct := structs.MstLogs{}
	targetReadStruct := structs.DetailTarget{}
	mstLogDescStruct := structs.MstLogDesc{}

	response := responseStruct

	id_target_c, _ := strconv.ParseInt(id_target, 10, 64)
	duration_c, _ := strconv.Atoi(duration)
	id_mst_log_desc_c, _ := strconv.ParseInt(id_mst_log_desc, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
	id_modul_c, _ := strconv.ParseInt(id_modul, 10, 64)
	id_data_c, _ := strconv.ParseInt(id_data, 10, 64)

	targetLogCreateStruct.IdCmsUsers = id_cms_users_c
	targetLogCreateStruct.Duration = duration_c
	targetLogCreateStruct.Recall = recall
	targetLogCreateStruct.IdTarget = id_target_c
	targetLogCreateStruct.IdMstLogDesc = id_mst_log_desc_c
	targetLogCreateStruct.CreatedAt = GetTimeNow()

	errTargetLogCreate := idb.DB.Table("target_log").Create(&targetLogCreateStruct).Error

	if errTargetLogCreate != nil {

		response.ApiMessage = errDBAdd

	} else {

		mstLogsStruct.IdCmsUsers = id_cms_users_c
		mstLogsStruct.IdData = id_data_c
		mstLogsStruct.IdModul = id_modul_c
		mstLogsStruct.Jenis = jenis
		mstLogsStruct.CreatedAt = GetTimeNow()

		_ = idb.DB.Table("mst_logs").Create(&mstLogsStruct).Error

		_ = idb.DB.Table("target").Where("id = " + id_data + "").First(&targetReadStruct).Error

		id_target_status := targetReadStruct.IdTargetMstStatus
		id_target = id_data
		id_status := 0

		_ = idb.DB.Raw("select * from mst_log_desc where id = " + id_mst_log_desc + "").Scan(&mstLogDescStruct).Error

		id_log_status := mstLogDescStruct.IdMstLogStatus

		if id_target_status == 1 && id_log_status == 1 {
			id_status = 2
		} else if id_target_status == 2 && id_log_status == 1 {
			id_status = 2
		} else if id_target_status == 3 && id_log_status == 1 {
			id_status = 2
		} else if id_target_status == 5 && id_log_status == 1 {
			id_status = 2
		} else if id_target_status == 1 && id_log_status == 2 {
			id_status = 5
		} else if id_target_status == 2 && id_log_status == 2 {
			id_status = 5
		} else if id_target_status == 3 && id_log_status == 2 {
			id_status = 5
		} else if id_target_status == 5 && id_log_status == 2 {
			id_status = 5
		} else {
			id_status = 3
		}

		_ = idb.DB.Exec("update target set updated_at = now() , " +
			"id_target_mst_status = " + fmt.Sprint(id_status) + " , updated_by = " + id_cms_users + " " +
			"where id = " + id_target + "").RowsAffected

		response.ApiStatus = 1
		response.ApiMessage = succ

	}

	return response
}

func (m *TargetModels) TargetNoteCreate(id_target string, id_cms_users string, note string) structs.JsonResponse {

	response := responseStruct
	targetNoteCreateStruct := structs.TargetNoteCreate{}

	id_target_c, _ := strconv.ParseInt(id_target, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)

	targetNoteCreateStruct.IdTarget = id_target_c
	targetNoteCreateStruct.IdCmsUsers = id_cms_users_c
	targetNoteCreateStruct.Note = note
	targetNoteCreateStruct.CreatedAt = GetTimeNow()

	err := idb.DB.Table("target_note").Create(&targetNoteCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd

	} else {

		response.ApiStatus = 1
		response.ApiMessage = succ
	}

	return response
}
func (m *TargetModels) TargetVisumCreate(id_target string, id_cms_users string, revisit string, id_mst_visum_status string, files multipart.File, header *multipart.FileHeader) structs.JsonResponse {

	response := responseStruct
	targetVisumCreateStruct := structs.TargetVisumCreate{}

	id_target_c, _ := strconv.ParseInt(id_target, 10, 64)
	id_cms_users_c, _ := strconv.ParseInt(id_cms_users, 10, 64)
	id_mst_visum_status_c, _ := strconv.ParseInt(id_mst_visum_status, 10, 64)

	targetVisumCreateStruct.IdTarget = id_target_c
	targetVisumCreateStruct.IdCmsUsers = id_cms_users_c
	targetVisumCreateStruct.Revisit = revisit
	targetVisumCreateStruct.IdMstVisumStatus = id_mst_visum_status_c
	targetVisumCreateStruct.CreatedAt = GetTimeNow()

	err := idb.DB.Table("target_visum").Create(&targetVisumCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd

	} else {

		uploadImageVisum := func() {

			url := UploadImage("target", fmt.Sprint(targetVisumCreateStruct.Id), files, header)

			if url != "" {

				targetVisumPhotoCreate := structs.TargetVisumPhotoCreate{}

				targetVisumPhotoCreate.IdCmsUsers = id_cms_users_c
				targetVisumPhotoCreate.IdTarget = id_target_c
				targetVisumPhotoCreate.IdTargetVisum = targetVisumCreateStruct.Id
				targetVisumPhotoCreate.Photo = url
				targetVisumPhotoCreate.CreatedAt = GetTimeNow()

				_ = idb.DB.Table("target_photo").Create(&targetVisumPhotoCreate).Error

			}
		}
		logCreate := func() {

			mstLogsStruct := structs.MstLogs{}

			mstLogsStruct.IdCmsUsers = id_cms_users_c
			mstLogsStruct.IdData = id_target_c
			mstLogsStruct.IdModul = 6
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
func (m *TargetModels) TargetSearch(id_cms_users string, search string) structs.JsonResponse {

	targetListingStruct := []structs.TargetListing{}
	response := responseStruct
	response.ApiStatus = 1
	response.Data = targetListingStruct

	err := idb.DB.Raw("select a.id , a.first_name , a.last_name , a.id_target_mst_status , b.recall , " +
		"b.id_mst_log_desc ,d.id_mst_log_status ,e.status , a.created_at as created_at_target , d.description ," +
		" f.id , f.created_at as created_at_target_visum , f.revisit , a.category , h.status as visit_status from target a " +
		"left join target_log b on b.id_target = a.id and b.id_cms_users = a.id_cms_users and b.id = " +
		"(select max(c.id) from target_log c where c.id_target = b.id_target and c.id_cms_users = b.id_cms_users) " +
		"left join mst_log_desc d on d.id = b.id_mst_log_desc " +
		"left join mst_log_status e on e.id = d.id_mst_log_status " +
		"left join target_visum f on f.id_target = a.id and f.id_cms_users = a.id_cms_users " +
		"and f.id = (select max(g.id) from target_visum g where g.id_target = f.id_target and g.id_cms_users =f.id_cms_users ) " +
		"left join mst_visum_status h " +
		"on h.id = f.id_mst_visum_status " +
		"where a.id_cms_users = " + id_cms_users + " and  concat(a.first_name , ' ' , a.last_name , ' ' , a.hp_1 , ' ' , a.hp_2) ilike '%" + search + "%' order by a.first_name asc ," +
		" a.updated_at desc,d.created_at desc , f.created_at desc  ").Scan(&targetListingStruct).Error

	if err != nil {
		response.ApiStatus = 1
		response.ApiMessage = emptyDB
	} else {
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = targetListingStruct
	}

	return response
}

//func (m *TargetModels)TargetPhoneCheckNewDB(number string)structs.JsonResponse{
//
//	response := responseStruct
//
//	rows := idb.DB.Raw("select * from target_phone where number = '"+number+"'").RecordNotFound()
//
//	if rows{
//
//		response.ApiStatus = 1
//		response.ApiMessage = succ
//	}else{
//
//		response.ApiMessage = succ
//	}
//
//	return response
//}
