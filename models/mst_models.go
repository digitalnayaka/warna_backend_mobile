package models

import (
	"../structs"
)

type MstModels struct {
}

func (m *MstModels) MstLogStatus() structs.JsonResponse {

	mstLogStatusStruct := []structs.MstLogStatus{}
	response := responseStruct
	response.Data = mstLogStatusStruct

	err := idb.DB.Table("mst_log_status").Order("status asc").Find(&mstLogStatusStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = mstLogStatusStruct
	}

	return response
}

func (m *MstModels) MstReligion() structs.JsonResponse {

	mstStruct := []structs.MstReligion{}
	response := responseStruct
	response.Data = mstStruct

	err := idb.DB.Table("mst_religion").Order("agama asc").Find(&mstStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = mstStruct
	}

	return response
}
func (m *MstModels) ContactMstStatusMarital() structs.JsonResponse {

	mstStruct := []structs.ContactMstStatusMarital{}
	response := responseStruct
	response.Data = mstStruct

	err := idb.DB.Table("contact_mst_status_marital").Order("status asc").Find(&mstStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = mstStruct
	}

	return response
}
func (m *MstModels) MstPlace() structs.JsonResponse {

	mstStruct := []structs.MstPlace{}
	response := responseStruct
	response.Data = mstStruct

	err := idb.DB.Table("contact_mst_status_place").Order("status asc").Find(&mstStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = mstStruct
	}

	return response
}
func (m *MstModels) ContactMstStatusEmployee() structs.JsonResponse {

	mstStruct := []structs.ContactMstStatusEmployee{}
	response := responseStruct
	response.Data = mstStruct

	err := idb.DB.Table("contact_mst_status_employee").Order("status asc").Find(&mstStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1
		response.ApiMessage = succ
		response.Data = mstStruct
	}

	return response
}

func (m *MstModels) MstVisumStatus() structs.JsonResponse {

	mstVisumStatusStruct := []structs.MstVisumStatus{}
	response := responseStruct
	response.Data = mstVisumStatusStruct

	err := idb.DB.Table("mst_visum_status").Order("status asc").Find(&mstVisumStatusStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1

		response.ApiMessage = succ
		response.Data = mstVisumStatusStruct
	}

	return response
}
func (m *MstModels) MstLogDesc(id_mst_log_status string) structs.JsonResponse {

	mstLogDescStruct := []structs.MstLogDesc{}
	response := responseStruct
	response.Data = mstLogDescStruct

	err := idb.DB.Table("mst_log_desc")

	if id_mst_log_status != "" {

		err = err.Where("id_mst_log_status = " + id_mst_log_status + "")
	}

	errx := err.Order("description asc").Find(&mstLogDescStruct).Error

	if errx != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1

		response.ApiMessage = succ
		response.Data = mstLogDescStruct
	}

	return response
}
func (m *MstModels) MstJob() structs.JsonResponse {

	mstJobStruct := []structs.MstJob{}
	response := responseStruct
	response.Data = mstJobStruct

	err := idb.DB.Table("mst_job").Order("job asc").Find(&mstJobStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1

		response.ApiMessage = succ
		response.Data = mstJobStruct
	}

	return response
}
func (m *MstModels) MstDatasource() structs.JsonResponse {

	mstDatasourceStruct := []structs.MstDataSource{}
	response := responseStruct
	response.Data = mstDatasourceStruct

	err := idb.DB.Table("mst_data_source").Order("datasource asc").Find(&mstDatasourceStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1

		response.ApiMessage = succ
		response.Data = mstDatasourceStruct
	}

	return response
}
func (m *MstModels) MstCategoryAddress() structs.JsonResponse {

	mstCategoryAddressStruct := []structs.MstCategoryAddress{}
	response := responseStruct
	response.Data = mstCategoryAddressStruct

	err := idb.DB.Table("mst_category_address").Order("category asc").Find(&mstCategoryAddressStruct).Error

	if err != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1

		response.ApiMessage = succ
		response.Data = mstCategoryAddressStruct
	}

	return response
}
func (m *MstModels) MstUnitUfi(id_mst_branch string, merk string, year string) structs.JsonResponse {

	mstUnitUfiStruct := []structs.MstUnitUfi{}
	response := responseStruct
	response.Data = mstUnitUfiStruct

	err := idb.DB.Table("mst_unit").Select("mst_unit.* , mst_branch.branch_name as mst_branch_branch_name").
		Joins("left join mst_branch on mst_branch.id = mst_unit.id_mst_branch")

	if id_mst_branch != "" {

		err = err.Where("mst_unit.id_mst_branch = " + id_mst_branch + "")
	}

	if merk != "" {

		err = err.Where("mst_unit.merk like '%" + merk + "%'")
	}

	if year != "" {

		err = err.Where("mst_unit.year = " + year + "")
	}

	err = err.Order("mst_unit.type asc").Scan(&mstUnitUfiStruct)
	errx := err.Error

	if errx != nil {
		response.ApiMessage = errDB
	} else {
		response.ApiStatus = 1

		response.ApiMessage = succ
		response.Data = mstUnitUfiStruct
	}

	return response
}
func (m *MstModels) MstCabangFif(branch_name string, pos_name string) structs.JsonResponse {

	response := responseStruct
	cabangFif := []structs.MstCabangFif{}

	err := idb.DB.Table("mst_cabang_fif").Select("distinct on (branch_name) *")

	if branch_name != "" {

		err = err.Where("branch_name ilike '%" + branch_name + "%'")
	}

	err = err.Order("branch_name asc").Find(&cabangFif)
	_ = err.Error

	response.ApiStatus = 1
	response.ApiMessage = succ
	response.Data = cabangFif

	return response
}
func (m *MstModels) MstNeed() structs.JsonResponse {

	response := responseStruct
	needStruct := []structs.MstNeed{}

	_ = idb.DB.Table("mst_need").Order("need asc").Find(&needStruct).Error

	response.ApiStatus = 1
	response.Data = needStruct
	response.ApiMessage = succ
	return response
}
