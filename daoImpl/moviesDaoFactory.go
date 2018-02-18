package dao

import (
	. "github.com/sil-vio/golang-restapi/config"
	. "github.com/sil-vio/golang-restapi/dao"
)

func MoviesDAOFactory() MoviesDAOInterface {
	var config = Config{}
	var dao = MoviesDAO{}
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
	return dao
}
