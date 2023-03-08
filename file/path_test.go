package file

import (
	"testing"
)

// 正确入参
func TestPathVerify(t *testing.T) {
	var err error
	//模版文件路径
	tmplPath := "../model.tmpl"
	//最终文件生成路径
	filePath := "./"
	//包名
	var packageName string

	//文件路径校验以及转换
	tmplPath, filePath, packageName, err = PathVerify(filePath)
	if err != nil {
		t.Errorf("PathVerify errr:%s", err.Error())
		return
	}
	t.Log(tmplPath)
	t.Log(filePath)
	t.Log(packageName)
}

// 模版文件不存在
func TestPathVerify_TmplNotExist(t *testing.T) {
	var err error

	//最终文件生成路径
	filePath := "./"

	//文件路径校验以及转换
	_, filePath, _, err = PathVerify(filePath)
	if err != nil {
		t.Logf("success %s", err.Error())
		return
	}
}

// 模版文件不存在
func TestPathVerify_FilePathNotExist(t *testing.T) {
	var err error
	//模版文件路径
	tmplPath := "../model.tmpl"
	//最终文件生成路径
	filePath := "./test_file/test"

	//文件路径校验以及转换
	tmplPath, filePath, _, err = PathVerify(filePath)
	if err != nil {
		t.Errorf("PathVerify errr:%s", err.Error())
		return
	}
	t.Log(tmplPath)
	t.Log(filePath)
}
