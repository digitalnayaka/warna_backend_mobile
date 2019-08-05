package models

import (

	"../config"
	// "github.com/gin-gonic/gin"

)

var (

	db = config.DBInit()
	idb = config.InDB{DB: db}
	// idb = inDB

)