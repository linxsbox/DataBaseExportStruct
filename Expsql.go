package main

import (
	"DataBaseExportStruct/ExportStructFile"
	//gd "DataBaseExportStruct/GetData"
	"flag"
	"fmt"
	"github.com/sqlhelper"
)

var (
	dbuser = flag.String("u", "root", "数据库用户名")
	dbpawd = flag.String("p", "root", "数据库密码")
	dbhose = flag.String("h", "127.0.0.1", "数据库连接")
	dbport = flag.String("port", "3306", "数据库端口")
	dbname = flag.String("name", "main", "数据库名")
	fpath  = flag.String("path", "当前文件", "导出路径")
)

func main() {
	flag.Parse()

	sqlhelper.DBName = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		*dbuser, *dbpawd, *dbhose, *dbport, *dbname)
	fmt.Println(sqlhelper.DBName, "\t", *dbname)

	sqlhelper.InitConnection()

	fileContent := ExportStructFile.DataStructToStr(*dbname)
	file, err := ExportStructFile.CreateFile(*fpath)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = ExportStructFile.FilePutContent(file, fileContent)
	if err != nil {
		fmt.Println(err.Error())
	}
}
