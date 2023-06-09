package initialize

func Init() {
	errConf := InitConfig()
	if errConf != nil {
		panic(errConf)
	}

	errSql := InitMysqlDB()
	if errSql != nil {
		panic(errSql)
	}

	errMinio := InitMinioIO()
	if errMinio != nil {
		panic(errMinio)
	}
}
