package gen

import (
	"github.com/go-libraries/genModels"
	"os"
)

func BuildModels(build *Build) {
	dsn := build.Dsn

	if dsn != "" {
		path := build.ProjectPath + "/" + build.ProjectName + "/models"
		Mysql := genModels.GetMysqlToGo()
		Mysql.Driver.SetDsn(dsn)
		Mysql.SetStyle("gorm")

		genModels.GormInit = ``
		genModels.GormTpl = ``
		//path := build.ProjectPath + "/" + build.ProjectName + "/models"
		_, e := os.Stat(path)
		if e != nil {
			os.MkdirAll(path, 0666)
		}
		Mysql.SetModelPath(path)
		Mysql.SetPackageName("models")
		Mysql.Run()
	}
}
