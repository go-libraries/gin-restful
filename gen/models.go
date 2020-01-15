package gen

func BuildModels(build *Build) {
	dsn := build.Dsn

	if dsn != "" {
		//Mysql := genModels.GetMysqlToGo()
		//Mysql.Driver.SetDsn(dsn)
		//
		////Mysql.SetStyle("bee")
		//Mysql.SetStyle("gorm")
		//path := build.ProjectPath + "/" + build.ProjectName + "/models"
		//_, e := os.Stat(path)
		//if e != nil {
		//	os.MkdirAll(path, 0666)
		//}
		//Mysql.SetModelPath(build.ProjectPath + "/" + build.ProjectName + "/models")
		////Mysql.SetIgnoreTables("cate")
		//Mysql.SetPackageName("models")
		//Mysql.Run()
	}
}
