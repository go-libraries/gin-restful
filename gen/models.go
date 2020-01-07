package gen

import (
	"fmt"
	"github.com/go-libraries/genModels"
)

func BuildModels(build *Build) {
	dsn := build.Dsn

	if dsn != "" {
		Mysql := genModels.GetMysqlToGo()
		fmt.Println(dsn)
		Mysql.Driver.SetDsn(dsn)

		//Mysql.SetStyle("bee")
		Mysql.SetStyle("gorm")
		Mysql.SetModelPath(build.ProjectPath + "/"+build.ProjectName+ "/models")
		//Mysql.SetIgnoreTables("cate")
		Mysql.SetPackageName("models")
		Mysql.Run()
	}
}
