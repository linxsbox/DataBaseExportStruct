package GetData

import (
	"fmt"
	"github.com/sqlhelper"
	//"strings"
	"time"
)

var (
	DataTableList       = make([]string, 20)
	DataTableRemrkeList = make(map[string]string)
	TableDataStruct     = make(map[string]string)
	TablePKList         = make(map[string]string)
	TableGetSet         = make(map[string]string)
	TableGet            = make([]string, 20)
	TableSet            = make([]string, 20)
	StrContent          = ""
	t                   = time.Now()
)

func DataBaseStructList() {
	fmt.Printf("读取数据库信息，指定数据库名称为: %v\n开始解析数据库结构\n", DataBaseName)
	databaseTableList, _ := sqlhelper.Query(
		GetDataBaseStruct(), DataBaseName)
	fmt.Printf("解析数据库结构完成 %v\n输出数据库表结构列表\n",
		time.Now().Sub(t))
	for key, v := range databaseTableList {
		if key >= len(DataTableList) {
			DataTableList = append(DataTableList, v[0])
		} else {
			DataTableList[key] = v[0]
		}
		DataTableRemrkeList[v[0]] = v[1]
		for _, val := range v {
			//fmt.Println("表名: ", val)
			tableStructList(key, val)
		}
	}
	fmt.Println("数据结构导出完成 ", time.Now().Sub(t))
	tablePKList()
}

func tableStructList(key int, val string) {
	//fmt.Printf("%v\t\t%v\t\t%v\n", "列名", "字段类型", "字段说明")
	tableTypeList, _ := sqlhelper.Query(
		GetDataBaseTableStruct(),
		DataBaseName, val)
	if val != "" {
		TableDataStruct[val], TableGetSet[val] = "", ""
	}
	for idx, v := range tableTypeList {
		tableGetSet(v[0], v[1], val)
		for k, value := range v {
			tableStructFormat(value, val, k)
		}
		if idx >= len(TableGet) {
			TableGet = append(TableGet, v[0])
		} else {
			TableGet[idx] = v[0]
		}
		//fmt.Println()
	}
	if TableDataStruct[val] == "" {
		//fmt.Println("删除了", val)
		delete(TableDataStruct, val)
	}
}

func tablePKList() {
	fmt.Println("开始处理数据关系")
	tablePKList, _ := sqlhelper.Query(
		GetDataBaseTablePK(), DataBaseName)
	fmt.Println("结果 ", tablePKList)
	fmt.Println("数据表列表 ", DataTableList)
	for _, v := range tablePKList {
		for _, vk := range DataTableList {
			if vk == "" {
				break
			}
			if v[0] == vk {
				TablePKList[vk] += fmt.Sprintf("\t*%v\n", v[1])
			}
		}
	}
}

func tableStructFormat(value, val string, k int) {
	tos := dataBaseChangeType(value)
	//fmt.Printf("%v\t\t", tos)
	if k == 2 {
		TableDataStruct[val] += fmt.Sprintf("//%v\n", tos)
	} else {
		TableDataStruct[val] += fmt.Sprintf("\t%v ", tos)
	}
}

func tableGetSet(sn, st, val string) {
	tostc := dataBaseChangeType(st)
	TableGetSet[val] += fmt.Sprintf("/* get set %v */\n", val)
	TableGetSet[val] += fmt.Sprintf("func (this *%v) Get%v() %v {\n",
		val, sn, tostc)
	TableGetSet[val] += fmt.Sprintf("\treturn this.%v\n}\n", sn)

	ctos := tableGetSetChangeToLower(sn)
	TableGetSet[val] += fmt.Sprintf("func (this *%v) Set%v(%v %v) {\n",
		val, sn, ctos, tostc)
	TableGetSet[val] += fmt.Sprintf("\tthis.%v = %v\n}\n", sn, ctos)
}
