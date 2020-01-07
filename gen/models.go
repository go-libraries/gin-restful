package gen

import (
	"github.com/go-libraries/genModels"
)

func BuildModels(build *Build) {
	dsn := build.Dsn

	if dsn != "" {
		Mysql := genModels.GetMysqlToGo()
		Mysql.Driver.SetDsn(dsn)

		//Mysql.SetStyle("bee")
		Mysql.SetStyle("gorm")
		Mysql.SetModelPath(build.ProjectPath)
		//Mysql.SetIgnoreTables("cate")
		Mysql.SetPackageName("models")
		Mysql.Run()
	}
}
