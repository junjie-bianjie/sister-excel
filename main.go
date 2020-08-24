package main

import (
	"forsisterexcel/config"
	"forsisterexcel/service"
	"forsisterexcel/utils"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//companies := service.Excel1(config.File1)
	//var companyNames []string
	//for _, company := range companies {
	//	companyNames = append(companyNames, company.CompanyName)
	//}

	// 之前的amy版本
	//diffs := utils.Different(companyNames, service.Excel2(config.File2))

	// 现在两个公司都在第二列的版本
	diffs := utils.Different(service.Excel2(config.File1), service.Excel2(config.File2))

	var fileStr string
	fileStr = "总计：" + strconv.Itoa(len(diffs))
	for _, v := range diffs {
		fileStr += "\n" + v
	}
	ioutil.WriteFile("输出结果.txt", []byte(fileStr), os.ModePerm)
}
