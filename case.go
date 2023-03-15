package main

import (
	"errors"
	"strings"
	"unicode"
)

type CaseType string

// # 下划线 underscore
// # 大驼峰 upper_camel_case
// # 小驼峰 lower_camel_case
const (
	CaseTypeUnderscore     CaseType = "underscore"
	CaseTypeUpperCamelCase CaseType = "upper_camel_case"
	CaseTypeLowerCamelCase CaseType = "lower_camel_case"
)

var (
	CaseErrorNotSupport = errors.New("暂不支持该类型转换")
	CaseErrorNotFound   = errors.New("未找到对应类型")
)

func ToCase(oldType, newType CaseType, s string) (string, error) {
	var caseNaming CaseNaming
	switch oldType {
	case CaseTypeUnderscore:
		caseNaming = NewUnderScoreCase(s)
	case CaseTypeUpperCamelCase:
		return "", CaseErrorNotSupport
	case CaseTypeLowerCamelCase:
		return "", CaseErrorNotSupport
	default:
		return "", CaseErrorNotFound
	}

	var result string
	switch newType {
	case CaseTypeUnderscore:
		result = caseNaming.ToUnderscore()
	case CaseTypeUpperCamelCase:
		result = caseNaming.ToUpperCamelCase()
	case CaseTypeLowerCamelCase:
		result = caseNaming.ToLowerCamelCase()
	default:
		return "", CaseErrorNotFound
	}
	return result, nil

}

type CaseNaming interface {
	// ToUnderscore 下划线命名法
	ToUnderscore() string
	// ToUpperCamelCase 大驼峰命名法
	ToUpperCamelCase() string
	// ToLowerCamelCase 小驼峰命名法
	ToLowerCamelCase() string
}

// underScoreCase 下划线命名的字符串
type underScoreCase struct {
	s string
}

func NewUnderScoreCase(s string) *underScoreCase {
	return &underScoreCase{
		s: s,
	}
}

// ToUnderscore 转下划线命名法
func (c *underScoreCase) ToUnderscore() string {
	return c.s
}

// ToUpperCamelCase 转大驼峰命名法
func (c *underScoreCase) ToUpperCamelCase() string {
	s := c.s
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	return s
}

// ToLowerCamelCase 转小驼峰命名法
func (c *underScoreCase) ToLowerCamelCase() string {
	s := c.ToUpperCamelCase()
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
