package file

import (
	"os"
	"path/filepath"
	"strings"
)

func PathVerify(filePath string) (string, string, string, error) {
	var err error
	//查询模版文件是否存在
	gopath := os.Getenv("GOPATH")
	tmplPath := gopath + "/bin"
	_, err = os.Stat(tmplPath)
	if err != nil {
		return "", "", "", err
	}

	//查询最终文件路径是否存在，不存在则创建
	filePath, err = filepath.Abs(filePath)
	if err != nil {
		return "", "", "", err
	}

	//获取最终文件的包名
	list := strings.Split(filePath, "/")
	packageName := list[len(list)-1]

	//查询模版文件是否存在
	_, err = os.Stat(filePath)

	if err == nil {
		return tmplPath, filePath, packageName, err
	}
	//不存在则进行创建
	if os.IsNotExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return "", "", "", err
		}
	}
	return tmplPath, filePath, packageName, nil
}
