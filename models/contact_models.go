package models

import (

	"../structs"
	"fmt"
	"strconv"
)

type ContactModels struct {

}

func (m *ContactModels)CheckAddContact(id_cms_users string , nik string)structs.JsonResponse{

	contactStruct := structs.Contact{}

	response := responseStruct
	//response.Data = contactStruct
	response.Data = contactStruct.Id
	response.ApiMessage = succ


	checkContact := idb.DB.Raw("select * from contact where nik = '"+nik+"'").Scan(&contactStruct).RecordNotFound()

	if checkContact {
		//tidak ada data
		response.ApiStatus = 0
	}else{

		//ada data

		contactColabCheckStruct := structs.ContactCollaborate{}
		checkContactCmsUser := idb.DB.Raw("select * from contact_collaborate where id_contact = "+fmt.Sprint(contactStruct.Id)+" and id_cms_users = "+id_cms_users+"").Scan(&contactColabCheckStruct).RecordNotFound()

		response.Data = contactStruct.Id
		if checkContactCmsUser {
			//tidak ada data
			response.ApiStatus = 1
		}else{
			//ada data
			response.ApiStatus = 2
		}

	}

	return response
}
func (m *ContactModels)ContactCreate(modul string , id_convert string ,nik string , first_name string , last_name string , birth_place string , birth_date string,
	gender string , id_mst_religion string , id_contact_mst_status_marital string , id_mst_job string , id_mst_data_source string ,
	id_cms_users string , number string , number2 string , address string , rt string , rw string , id_mst_address string,
	id_mst_category_address string , mother string , family string , company string , position string , working_time string ,
	income string , outlay string , id_contact_mst_status_place string , id_contact_mst_status_employee string )structs.JsonResponse{

	contactStruct := structs.Contact{}
	contactCollabStruct := structs.ContactCollaborate{}
	contactPhoneStruct := structs.ContactPhone{}
	contactAddressStruct := structs.ContactAddress{}
	contactAdditionalStruct := structs.ContactAdditional{}

	response := responseStruct

	id_mst_religion_c , _ := strconv.ParseInt(id_mst_religion,10,64)
	id_contact_mst_status_marital_c , _ := strconv.ParseInt(id_contact_mst_status_marital,10,64)
	id_mst_job_c , _ := strconv.ParseInt(id_mst_job,10,64)
	id_mst_data_source_c , _ := strconv.ParseInt(id_mst_data_source,10,64)
	id_cms_users_c , _ := strconv.ParseInt(id_cms_users,10,64)
	id_mst_address_c , _ := strconv.ParseInt(id_mst_address,10,64)
	id_mst_category_address_c , _ := strconv.ParseInt(id_mst_category_address,10,64)
	id_contact_mst_status_place_c , _ := strconv.ParseInt(id_contact_mst_status_place,10,64)
	id_contact_mst_status_employee_c , _ := strconv.ParseInt(id_contact_mst_status_employee,10,64)
	income_c , _ := strconv.Atoi(income)
	outlay_c , _ := strconv.Atoi(outlay)
	working_time_c , _ := strconv.Atoi(working_time)

	contactStruct.CreatedAt = GetTimeNow()
	contactStruct.IdCmsUsers = id_cms_users_c
	contactStruct.IdMstJob = id_mst_job_c
	contactStruct.IdMstDataSource = id_mst_data_source_c
	contactStruct.FirstName = first_name
	contactStruct.LastName = last_name
	contactStruct.BirthDate = birth_date
	contactStruct.Gender = gender
	contactStruct.BirthPlace = birth_place
	contactStruct.IdContactMstStatusMarital = id_contact_mst_status_marital_c
	contactStruct.IdMstReligion = id_mst_religion_c
	contactStruct.Nik = nik

	err := idb.DB.Table("contact").Create(&contactStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd

	}else{

		contactColabAdd := func() {

			contactCollabStruct.IdCmsUsers = id_cms_users_c
			contactCollabStruct.IdContact = contactStruct.Id

			_ = idb.DB.Table("contact_collaborate").Create(&contactCollabStruct).Error
		}
		contactPhoneCreate := func() {

			contactPhoneStruct.IdContact = contactStruct.Id
			contactPhoneStruct.IdCmsUsers = id_cms_users_c
			contactPhoneStruct.CreatedAt = GetTimeNow()
			contactPhoneStruct.Number = number
			contactPhoneStruct.Status = "Y"

			_ = idb.DB.Table("contact_phone").Create(&contactPhoneStruct).Error

			if number2 != ""{

				contactPhoneStruct := structs.ContactPhone{}

				contactPhoneStruct.IdContact = contactStruct.Id
				contactPhoneStruct.IdCmsUsers = id_cms_users_c
				contactPhoneStruct.CreatedAt = GetTimeNow()
				contactPhoneStruct.Number = number2
				contactPhoneStruct.Status = "Y"

				_ = idb.DB.Table("contact_phone").Create(&contactPhoneStruct).Error

			}

		}
		contactAddressCreate := func() {

			contactAddressStruct.CreatedAt = GetTimeNow()
			contactAddressStruct.IdCmsUsers = id_cms_users_c
			contactAddressStruct.IdContact = contactStruct.Id
			contactAddressStruct.Address = address
			contactAddressStruct.IdMstCategoryAddress = id_mst_category_address_c
			contactAddressStruct.IdMstAddress = id_mst_address_c
			contactAddressStruct.Rt = rt
			contactAddressStruct.Rw = rw

			_ = idb.DB.Table("contact_address").Create(&contactAddressStruct).Error
		}
		contactAdditionalCreate := func() {

			contactAdditionalStruct.IdContact = contactStruct.Id
			contactAdditionalStruct.IdCmsUsers = id_cms_users_c
			contactAdditionalStruct.CreatedAt = GetTimeNow()
			contactAdditionalStruct.Company = company
			contactAdditionalStruct.Family = family
			contactAdditionalStruct.IdContactMstStatusEmployee = id_contact_mst_status_employee_c
			contactAdditionalStruct.IdContactMstStatusPlace = id_contact_mst_status_place_c
			contactAdditionalStruct.Income = income_c
			contactAdditionalStruct.Outlay = outlay_c
			contactAdditionalStruct.Mother = mother
			contactAdditionalStruct.Position = position
			contactAdditionalStruct.WorkingTime = working_time_c

			_ = idb.DB.Table("contact_additional").Create(&contactAdditionalStruct).Error
		}
		checkModul := func() {

			switch modul {
			case "lead":
				_ = idb.DB.Raw("update lead set id_lead_mst_status = 5 , updated_by = "+id_cms_users+" , " +
					"updated_at = now() where id = "+id_convert+"").Error
				break
			case "target":
				_ = idb.DB.Raw("update target set id_target_mst_status = 6 , updated_by = "+id_cms_users+" , " +
					"updated_at = now() where id = "+id_convert+"").Error
				break


			}
		}
		logCreate := func() {

			mstLogsStruct := structs.MstLogs{}

			mstLogsStruct.IdCmsUsers = id_cms_users_c
			mstLogsStruct.IdData = contactStruct.Id
			mstLogsStruct.IdModul = 4
			mstLogsStruct.Jenis = "create"
			mstLogsStruct.CreatedAt = GetTimeNow()

			_ = idb.DB.Table("mst_logs").Create(&mstLogsStruct).Error

		}

		go contactColabAdd()
		go contactPhoneCreate()
		go contactAddressCreate()
		go contactAdditionalCreate()
		go checkModul()
		go logCreate()

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = contactStruct.Id

	}



	return response
}

func (m *ContactModels)ContactList(id_cms_users string , limit string , offset string)structs.JsonResponse{

	response := responseStruct
	contactListStruct := []structs.ContactList{}
	response.Data = contactListStruct

	err := idb.DB.Raw("select a.id , " +
		"b.id as id_contact ," +
		"b.first_name as contact_first_name ," +
		"b.last_name as contact_last_name ," +
		"b.gender as contact_gender ," +
		"c.id_order_mst_status," +
		"d.status " +
		"from contact_collaborate a " +
		"left join contact b " +
		"on b.id = a.id_contact " +
		"left join "+order+" c " +
		"on c.id_contact = b.id and c.id = (select max(e.id) from "+order+" e where e.id_contact = b.id) " +
		"left join order_mst_status d " +
		"on d.id = c.id_order_mst_status " +
		"where a.id_cms_users = "+id_cms_users+" " +
		"order by b.updated_at desc , b.created_at desc")

	if limit != "" {

		limit_c , _ := strconv.Atoi(limit)
		err = err.Limit(limit_c)
	}

	if offset != ""{

		offset_c , _ := strconv.Atoi(offset)
		err = err.Offset(offset_c)
	}

	errx := err.Scan(&contactListStruct).Error

	if errx != nil {

		response.ApiMessage = errDB
	}else{

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = contactListStruct
	}

	return response
}

func (m *ContactModels)ContactDetail(id_contact string)structs.JsonResponse{

	response := responseStruct
	contactDetailInfoStruct := structs.ContactDetailInfo{}
	contactInfoStruct := structs.ContactInfo{}
	contactPhoneStruct := []structs.ContactPhoneDetail{}
	contactAddressStruct := []structs.ContactAddressDetail{}

	err := idb.DB.Raw("select a.* , c.id_order_mst_status , w.status ," +
		"(select count(id) from "+order+" z where z.id_contact=a.id and z.id_order_mst_status=2) as potential_order , " +
		"(select count(id) from "+order+" x where x.id_contact=a.id and x.id_order_mst_status=5) as won_order ," +
		"d.agama as mst_religion_agama , e.status as contact_mst_status_marital_status , f.job as mst_job_job , " +
		"g.name as cms_users_name , h.datasource as mst_data_source_datasource , " +
		"i.email , i.mother , i.family , j.status as contact_mst_status_place_status , i.company , " +
		"k.status as contact_mst_status_employee_status , i.position , i.working_time , i.income , i.outlay " +
		"from contact a " +
		"left join "+order+" c " +
		"on c.id_contact = a.id and c.id = (select max(e.id) from "+order+" e where e.id_contact = a.id) " +
		"left join order_mst_status w " +
		"on w.id = c.id_order_mst_status " +
		"left join mst_religion d " +
		"on d.id = a.id_mst_religion " +
		"left join contact_mst_status_marital e " +
		"on e.id = a.id_contact_mst_status_marital " +
		"left join mst_job f " +
		"on f.id = a.id_mst_job " +
		"left join cms_users g " +
		"on g.id = a.id_cms_users " +
		"left join mst_data_source h " +
		"on h.id = a.id_mst_data_source " +
		"left join contact_additional i " +
		"on i.id_contact = a.id " +
		"left join contact_mst_status_place j " +
		"on j.id = i.id_contact_mst_status_place " +
		"left join contact_mst_status_employee k " +
		"on k.id = i.id_contact_mst_status_employee " +
		"where a.id = "+id_contact+"").Scan(&contactInfoStruct).Error

	if err != nil {

		response.ApiMessage = errDB
	}else{

		_ = idb.DB.Table("contact_phone").Where("id_contact = "+id_contact+"").Find(&contactPhoneStruct).Error
		_ = idb.DB.Raw("select * , c.category as mst_category_address_category from contact_address a " +
			"left join mst_address b " +
			"on b.id = a.id_mst_address " +
			"left join mst_category_address c " +
			"on c.id = a.id_mst_category_address " +
			"where id_contact = "+id_contact+"").Scan(&contactAddressStruct).Error

		contactDetailInfoStruct.ContactInfo = contactInfoStruct
		contactDetailInfoStruct.ContactAddressDetail = contactAddressStruct
		contactDetailInfoStruct.ContactPhoneDetail = contactPhoneStruct

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = contactDetailInfoStruct
	}

	return response
}

func (m *ContactModels)ContactNote(id_contact string)structs.JsonResponse{

	contactNoteStruct := []structs.ContactNote{}
	response := responseStruct

	err := idb.DB.Raw("select a.id , a.note , b.name as cms_users_name , a.created_at from contact_note a " +
		"left join cms_users b " +
		"on b.id = a.id_cms_users where a.id_contact = "+id_contact+" order by a.created_at desc").Scan(&contactNoteStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	}else{

		response.ApiMessage = succ
		response.Data = contactNoteStruct
	}

	return response

}

func (m *ContactModels)ContactNoteCreate(id_contact string , id_cms_users string , note string)structs.JsonResponse{

	response := responseStruct
	contactNoteCreateStruct := structs.ContactNoteCreate{}

	id_contact_c , _ := strconv.ParseInt(id_contact,10,64)
	id_cms_users_c , _ := strconv.ParseInt(id_cms_users,10,64)

	contactNoteCreateStruct.IdContact = id_contact_c
	contactNoteCreateStruct.IdCmsUsers = id_cms_users_c
	contactNoteCreateStruct.Note = note
	contactNoteCreateStruct.CreatedAt = GetTimeNow()

	err := idb.DB.Table("contact_note").Create(&contactNoteCreateStruct).Error

	if err != nil {

		response.ApiMessage = errDBAdd

	}else{

		response.ApiStatus = 1
		response.ApiMessage = succ
	}

	return response
}

func (m *ContactModels)ContactCollabList(id_contact string)structs.JsonResponse{

	response := responseStruct
	contactCollabStruct := []structs.ContactCollabList{}
	response.Data = contactCollabStruct

	err := idb.DB.Raw("select a.id , b.name as cms_users_name , b.npm as cms_users_npm from contact_collaborate a " +
		"left join cms_users b " +
		"on b.id = a.id_cms_users " +
		"where a.id_contact = "+id_contact+"").Scan(&contactCollabStruct).Error

	if err!=nil {
		response.ApiMessage = errDB
	}else{
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = contactCollabStruct
	}

	return response
}

func (m *ContactModels)ContactUpdate(id_contact string ,nik string , first_name string , last_name string , birth_place string , birth_date string,
	gender string , id_mst_religion string , id_contact_mst_status_marital string , id_mst_job string , id_mst_data_source string ,
	id_cms_users string , number string , number2 string , address string , rt string , rw string , id_mst_address string,
	id_mst_category_address string , mother string , family string , company string , position string , working_time string ,
	income string , outlay string , id_contact_mst_status_place string , id_contact_mst_status_employee string , id_number1 string , id_number2 string)structs.JsonResponse{

	contactStruct := structs.Contact{}
	contactPhoneStruct := structs.ContactPhone{}
	contactAddressStruct := structs.ContactAddress{}
	contactAdditionalStruct := structs.ContactAdditional{}

	response := responseStruct

	id_mst_religion_c , _ := strconv.ParseInt(id_mst_religion,10,64)
	id_contact_mst_status_marital_c , _ := strconv.ParseInt(id_contact_mst_status_marital,10,64)
	id_mst_job_c , _ := strconv.ParseInt(id_mst_job,10,64)
	id_mst_data_source_c , _ := strconv.ParseInt(id_mst_data_source,10,64)
	id_cms_users_c , _ := strconv.ParseInt(id_cms_users,10,64)
	id_mst_address_c , _ := strconv.ParseInt(id_mst_address,10,64)
	id_mst_category_address_c , _ := strconv.ParseInt(id_mst_category_address,10,64)
	id_contact_mst_status_place_c , _ := strconv.ParseInt(id_contact_mst_status_place,10,64)
	id_contact_mst_status_employee_c , _ := strconv.ParseInt(id_contact_mst_status_employee,10,64)
	income_c , _ := strconv.Atoi(income)
	outlay_c , _ := strconv.Atoi(outlay)
	working_time_c , _ := strconv.Atoi(working_time)

	contactStruct.IdCmsUsers = id_cms_users_c
	contactStruct.IdMstJob = id_mst_job_c
	contactStruct.IdMstDataSource = id_mst_data_source_c
	contactStruct.FirstName = first_name
	contactStruct.LastName = last_name
	contactStruct.BirthDate = birth_date
	contactStruct.Gender = gender
	contactStruct.BirthPlace = birth_place
	contactStruct.IdContactMstStatusMarital = id_contact_mst_status_marital_c
	contactStruct.IdMstReligion = id_mst_religion_c
	contactStruct.Nik = nik
	contactStruct.UpdatedAt = GetTimeNow()
	contactStruct.UpdatedBy = id_cms_users_c

	err := idb.DB.Table("contact").Where("id = "+id_contact+"").Update(&contactStruct).Error

	if err != nil {

		response.ApiMessage = errDBUpdate

	}else{

		contactPhoneCreate := func() {

			contactPhoneStruct.Number = number
			contactPhoneStruct.Status = "Y"
			contactPhoneStruct.UpdatedAt = GetTimeNow()
			contactPhoneStruct.UpdatedBy = id_cms_users_c

			_ = idb.DB.Table("contact_phone").Where("id = "+id_number1+"").Update(&contactPhoneStruct).Error

			if number2 != ""{

				contactPhoneStruct := structs.ContactPhone{}

				contactPhoneStruct.Number = number2
				contactPhoneStruct.Status = "Y"
				contactPhoneStruct.UpdatedAt = GetTimeNow()
				contactPhoneStruct.UpdatedBy = id_cms_users_c

				_ = idb.DB.Table("contact_phone").Where("id = "+id_number2+"").Update(&contactPhoneStruct).Error

			}

		}
		contactAddressCreate := func() {

			contactAddressStruct.Address = address
			contactAddressStruct.IdMstCategoryAddress = id_mst_category_address_c
			contactAddressStruct.IdMstAddress = id_mst_address_c
			contactAddressStruct.Rt = rt
			contactAddressStruct.Rw = rw
			contactAddressStruct.UpdatedAt = GetTimeNow()

			_ = idb.DB.Table("contact_address").Where("id_contact = "+id_contact+"").Update(&contactAddressStruct).Error
		}
		contactAdditionalCreate := func() {

			contactAdditionalStruct.Company = company
			contactAdditionalStruct.Family = family
			contactAdditionalStruct.IdContactMstStatusEmployee = id_contact_mst_status_employee_c
			contactAdditionalStruct.IdContactMstStatusPlace = id_contact_mst_status_place_c
			contactAdditionalStruct.Income = income_c
			contactAdditionalStruct.Outlay = outlay_c
			contactAdditionalStruct.Mother = mother
			contactAdditionalStruct.Position = position
			contactAdditionalStruct.WorkingTime = working_time_c
			contactAdditionalStruct.UpdatedAt = GetTimeNow()
			contactAdditionalStruct.UpdatedBy = id_cms_users_c

			_ = idb.DB.Table("contact_additional").Where("id_contact = "+id_contact+"").Update(&contactAdditionalStruct).Error
		}


		go contactPhoneCreate()
		go contactAddressCreate()
		go contactAdditionalCreate()

		response.ApiStatus = 1
		response.ApiMessage = succ

	}



	return response
}







