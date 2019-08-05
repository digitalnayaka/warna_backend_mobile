package models

import (
	"../structs"
	"fmt"
	"strconv"
)

type CmsUser struct {

}

func (m *CmsUser) Login(npm string , password string , device_id string ,useragent string , version string) structs.JsonResponse  {

	response := responseStruct
	cmsStruct := cmsUserStruct

	response.Data = cmsStruct

	err := idb.DB.Raw("select * , outlet_name as mst_outlet_outlet_name , cms_privileges.name as cms_privileges_name " +
		"from cms_users " +
		"left join cms_privileges " +
		"on cms_privileges.id = cms_users.id_cms_privileges " +
		"left join mst_outlet" +
		" on mst_outlet.id = cms_users.id_mst_outlet " +
		"where npm = '"+npm+"' limit 1").
		Scan(&cmsStruct).RecordNotFound()

	if err {
		response.ApiMessage = "NPM tidak ditemukan"
	}else{

		checkPassword , errPass := DecryptPassword(cmsStruct.Password)

		if errPass!=nil {

			response.ApiMessage = "Password salah"
		}else{

			if checkPassword == password{

				if cmsStruct.Status == "Y" {

					eula := func() {

						eulaStruct := structs.UserAgreement{}

						eulaStruct.CreatedAt = GetTimeNow()
						eulaStruct.Eula = "Y"
						eulaStruct.IdCmsUsers = cmsStruct.Id

						idb.DB.Table("user_agreement").Create(&eulaStruct)
					}
					userlogs := func() {

						userlogsStruct := structs.UserLogs{}

						userlogsStruct.IdCmsUsers = cmsStruct.Id
						userlogsStruct.CreatedAt = GetTimeNow()
						userlogsStruct.Description = "Login"
						userlogsStruct.Useragent = useragent
						userlogsStruct.Version = version

						idb.DB.Table("user_logs").Create(&userlogsStruct)
					}
					device_id := func() {

						idb.DB.Exec("update cms_users set device_id = '"+device_id+"' where id = "+fmt.Sprint(cmsStruct.Id)+"")
					}

					go eula()
					go userlogs()
					go device_id()

					response.ApiStatus = 1
					response.ApiMessage = succ
					response.Data = cmsStruct

				}else{

					response.ApiMessage = "Akun anda dinonaktifkan , tidak dapat diakses"
				}

			}else{

				response.ApiMessage = "Password salah"
			}
		}
	}

	return response
}

func (m *CmsUser) CheckDeviceId(id string , devicde_id string)structs.JsonResponse{

	deviceCheck := cmsUserStruct
	response := responseStruct
	response.Data = deviceCheck
	err := idb.DB.Raw("select * from cms_users where id = "+id+" and device_id = '"+devicde_id+"'").Scan(&deviceCheck).Error

	if err != nil{
		response.ApiMessage = "ID Device tidak sama"
	}else{
		response.ApiStatus = 1
		response.ApiMessage = succ
	}

	return response

}

func (m *CmsUser) ActivityUser(id_cms_users string , limit string , offset string)structs.JsonResponse{

	logsStruct := []structs.ListLogUser{}
	response := responseStruct
	response.ApiStatus = 1
	response.Data = logsStruct

	err := idb.DB.Raw("select * , cms_users.name as cms_users_name from mst_logs left join cms_users on cms_users.id = mst_logs.id_cms_users where mst_logs.id_cms_users = "+id_cms_users+" order by mst_logs.created_at desc")

	if limit != ""{
		limits,_:= strconv.Atoi(limit)
		err = err.Limit(limits)
	}

	if offset != ""{
		offsets,_:= strconv.Atoi(offset)
		err = err.Offset(offsets)
	}

	err = err.Scan(&logsStruct)
	errx := err.Error

	if errx != nil{

		response.ApiMessage = emptyDB
	}else{

		for i , value  := range logsStruct{

			switch value.IdModul {
			case 1:

				break
			case 2 , 7:

				humanStruct := structs.HumanName{}

				idb.DB.Table("lead").Where("id = "+fmt.Sprint(value.IdData)+"").Find(&humanStruct)

				logsStruct[i].DataAktivitas = humanStruct.FirstName+" "+humanStruct.LastName

				break

			case 3,6:

				humanStruct := structs.HumanName{}

				idb.DB.Table("target").Where("id = "+fmt.Sprint(value.IdData)+"").Find(&humanStruct)

				logsStruct[i].DataAktivitas = humanStruct.FirstName+" "+humanStruct.LastName

				break
			case 4:

				humanStruct := structs.HumanName{}

				idb.DB.Table("contact").Where("id = "+fmt.Sprint(value.IdData)+"").Find(&humanStruct)

				logsStruct[i].DataAktivitas = humanStruct.FirstName+" "+humanStruct.LastName

				break

			case 5:

				humanStruct := structs.HumanName{}

				tableOrder := `"order" a`
				//tableOrder := "`order` a"

				idb.DB.Raw("select * from "+tableOrder+" left join contact on contact.id = a.id_contact where a.id = "+fmt.Sprint(value.IdData)+"").Scan(&humanStruct)

				logsStruct[i].DataAktivitas = humanStruct.FirstName+" "+humanStruct.LastName

				break


				
			}
		}

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = logsStruct
	}

	return response
}

func (m *CmsUser) PerformaIndicator(id_cms_users string , inputDate string , inputDateBulanKemarin string)structs.JsonResponse{

	response := responseStruct
	performaIndicatorStruct := structs.PerformIndicator{}

	err := idb.DB.Raw("select " +
		"(SELECT sum(activity_report.brosur::int8) FROM activity_report WHERE activity_report.id_cms_users = "+id_cms_users+" and activity_report.created_at::text ilike '%"+inputDate+"%') as count_brosur , " +
		"(SELECT COUNT(target_log.id) FROM target_log WHERE target_log.created_at::text ilike '%"+inputDate+"%' and target_log.id_cms_users = "+id_cms_users+") as count_tele , " +
		"(select count(lead.id) from lead where lead.id_cms_users = "+id_cms_users+" and lead.created_at::text ilike '%"+inputDate+"%') as count_newDB , " +
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+" and a.created_at::text ilike '%"+inputDate+"%') as count_order , " +
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+" and a.created_at::text ilike '%"+inputDate+"%' and a.id_order_mst_status = 5) as count_booking , " +
		"(SELECT COUNT(target_log.id) FROM target_log WHERE target_log.created_at::text ilike '%"+inputDateBulanKemarin+"%' and target_log.id_cms_users = "+id_cms_users+") as count_teleBlnKemarin , " +
		"(select count(lead.id) from lead where lead.id_cms_users = "+id_cms_users+" and lead.created_at::text ilike '%"+inputDateBulanKemarin+"%') as count_newDBBlnKemarin , " +
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+" and a.created_at::text ilike '%"+inputDateBulanKemarin+"%') as count_orderBulanKemarin , " +
		"(SELECT COUNT(target_log.id) FROM target_log WHERE target_log.id_cms_users = "+id_cms_users+") as count_teleTotal , " +
		"(select count(lead.id) from lead where lead.id_cms_users = "+id_cms_users+") as count_newDBTotal , " +
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+") as count_orderTotal , " +
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+" and a.id_order_mst_status = 5) as count_bookingTotal , " +
		"(SELECT sum(activity_report.brosur::int8) FROM activity_report WHERE activity_report.id_cms_users = "+id_cms_users+" and activity_report.created_at::text ilike '%"+inputDateBulanKemarin+"%') as count_brosurBulanKemarin , " +
		"(SELECT sum(activity_report.brosur::int8) FROM activity_report WHERE activity_report.id_cms_users = "+id_cms_users+") as count_brosurBulanTotal , " +
		"(SELECT count(target_visum.id)+(SELECT count(lead_visum.id) FROM lead_visum WHERE lead_visum.id_cms_users = "+id_cms_users+" and lead_visum.created_at::text ilike '%"+inputDate+"%') as total FROM target_visum WHERE target_visum.id_cms_users = "+id_cms_users+" and target_visum.created_at::text ilike '%"+inputDate+"%') as count_visum , " +
		"(SELECT count(target_visum.id)+(SELECT count(lead_visum.id) FROM lead_visum WHERE lead_visum.id_cms_users = "+id_cms_users+" and lead_visum.created_at::text ilike '%"+inputDateBulanKemarin+"%') as total FROM target_visum WHERE target_visum.id_cms_users = "+id_cms_users+" and target_visum.created_at::text ilike '%"+inputDateBulanKemarin+"%') as count_visumBlnKemarin , " +
		"(SELECT count(target_visum.id)+(SELECT count(lead_visum.id) FROM lead_visum WHERE lead_visum.id_cms_users = "+id_cms_users+") as total FROM target_visum WHERE target_visum.id_cms_users = "+id_cms_users+") as count_visumTotal , "+
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+" and a.created_at::text ilike '%"+inputDateBulanKemarin+"%' and a.id_order_mst_status = 5) as count_bookingBulanKemarin").
		Scan(&performaIndicatorStruct).Error

	if err != nil{

		response.ApiMessage = errDB

	}else{

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = performaIndicatorStruct
	}


	return response
}
func (m *CmsUser) RekapActivity(id_cms_users string , inputDate string , input_id_mst_outlet string)structs.JsonResponse{

	response := responseStruct
	rekapStruct := structs.Rekap{}

	err := idb.DB.Raw("select " +
		"(SELECT sum(activity_report.brosur::int8) FROM activity_report WHERE activity_report.id_cms_users = "+id_cms_users+" and activity_report.created_at::text ilike '%"+inputDate+"%') as count_brosur , " +
		"(SELECT COUNT(target_log.id) FROM target_log WHERE target_log.created_at::text ilike '%"+inputDate+"%' and target_log.id_cms_users = "+id_cms_users+") as count_tele , " +
		"(select count(lead.id) from lead where lead.id_cms_users = "+id_cms_users+" and lead.created_at::text ilike '%"+inputDate+"%') as count_newDB , " +
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+" and a.created_at::text ilike '%"+inputDate+"%') as count_order , " +
		"(select count(a.id) from "+order+" a where a.id_cms_users = "+id_cms_users+" and a.created_at::text ilike '%"+inputDate+"%' and a.id_order_mst_status = 5) as count_booking , " +
		"(SELECT count(target_visum.id)+(SELECT count(lead_visum.id) FROM lead_visum WHERE lead_visum.id_cms_users = "+id_cms_users+" and lead_visum.created_at::text ilike '%"+inputDate+"%') as total FROM target_visum WHERE target_visum.id_cms_users = "+id_cms_users+" and target_visum.created_at::text ilike '%"+inputDate+"%') as count_visum , " +
		"(select count(a.id) from "+order+" a where a.id_mst_outlet = "+input_id_mst_outlet+" and a.created_at::text ilike '%"+inputDate+"%') as count_order_outlet ," +
		"(select count(a.id) from "+order+" a where a.id_mst_outlet = "+input_id_mst_outlet+" and a.created_at::text ilike '%"+inputDate+"%' and a.id_order_mst_status = 5) as count_booking_outlet").
		Scan(&rekapStruct).Error

	if err != nil{

		response.ApiMessage = errDB

	}else{

		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = rekapStruct
	}


	return response
}
