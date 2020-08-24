package service

import (
	"fmt"
	"forsisterexcel/types"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func Excel1(filePath, sheet string) []types.Company {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Print("文件打开错误", err)
		return nil
	}

	rows := f.GetRows(sheet)

	var companies []types.Company
	// 获取工作表中指定的单元格
	for _, row := range rows {
		// 将公司放入List
		var company types.Company
		company.Author = row[0]
		company.CompanyName = row[2]
		companies = append(companies, company)
	}

	var result []types.Company
	// 去除，只要amy的数据
	for _, company := range companies {
		if company.Author == "Amy" {
			result = append(result, company)
		}
	}
	return result
}

func Excel2(filePath, sheet string) []string {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Print("文件打开错误", err)
		return nil
	}

	var companies []string
	rows := f.GetRows(sheet)
	for _, row := range rows {
		// 将公司放入List
		company := row[1]
		companies = append(companies, company)
	}
	return companies
}
