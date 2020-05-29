package ExportStructFile

import (
	gd "DataBaseExportStruct/GetData"
	"fmt"
	"os"
	"os/exec"
    "path/filepath"
)

func DataStructToStr(databasename string) string {
	gd.DataBaseName = databasename
	gd.DataBaseStructList()
	return formatDataStructFile()
}
func cfdir(){
 	file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
     CreateFile(path)
}

func CreateFile(dir string) (string, error) {
	src := fmt.Sprintf("%v/", dir) //gd.DataBaseName
	fmt.Println(src)
	if FileIsExist(src) {
		fmt.Println("文件夹已存在!")
		return src, nil
	}

	err := os.MkdirAll(src, 0777)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("你不够权限创建文件")
		}
		return "", err
	}
	return src, nil
}
func FilePutContent(filename, filecontent string) (int, error) {
	path := fmt.Sprintf("%v%v.go", filename, gd.DataBaseName)
	fmt.Println(path)
	if FileIsExist(path) {
		fmt.Println("该文件已存在!")
		return -1, nil
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Println("无法创建文件!")
		return 0, err
	}
	defer f.Close()
	return f.WriteString(filecontent)
}

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func formatDataStructFile() string {
	gd.StrContent += fmt.Sprintf("package %v\nimport(\n\t%v\n)\n",
		"main", gd.IsTime) //gd.DataBaseName
	for _, val := range gd.DataTableList {
		if val == "" {
			break
		}
		gd.StrContent += fmt.Sprintf("/**\n*%v\n*/\n",
			gd.DataTableRemrkeList[val])
		gd.StrContent += fmt.Sprintf("type %v struct{\n", val)
		gd.StrContent += gd.TableDataStruct[val]
		gd.StrContent += fmt.Sprintf(gd.TablePKList[val])
		gd.StrContent += fmt.Sprintf("}\n")
		gd.StrContent += gd.TableGetSet[val]
	}
	return gd.StrContent
}
