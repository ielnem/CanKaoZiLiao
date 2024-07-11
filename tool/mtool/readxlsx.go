package main

import (
	"github.com/xuri/excelize/v2"
)

const xlsxDefaultName = "servers.xlsx"

type Server struct {
	ip       string
	port     string
	user     string
	password string
}

func readXlsx() ([]*Server, error) {
	ret := make([]*Server, 0)
	wb, err := excelize.OpenFile(excelName)
	if err != nil {
		return nil, err
	}
	defer wb.Close()
	rows, err := wb.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for idx, row := range rows {
		if idx == 0 {
			continue
		}
		if len(row) < 5 {
			continue
		}
		tmp := new(Server)
		tmp.ip = row[1]
		tmp.port = row[2]
		tmp.user = row[3]
		tmp.password = row[4]
		ret = append(ret, tmp)
	}

	return ret, nil
}
