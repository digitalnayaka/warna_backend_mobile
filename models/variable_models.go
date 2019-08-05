package models

import (
	"../structs"
	"time"
)

var errDB = "Gagal mengambil data"
var errDBUpdate = "Gagal mengubah data"
var errDBAdd = "Gagal menambah data"
var emptyDB = "Data Kosong"
var succ = "Sukses"
var order = `"order"`
var errDBDelete = "Gagal menghapus"

var responseStruct structs.JsonResponse
var cmsUserStruct structs.CmsUsersStruct

func GetTimeNow() string {

	t := time.Now()
	return string(t.Format("2006-01-02 15:04:05"))
	//return string(t.Format("2006-01-02 15:04:05"))
}


