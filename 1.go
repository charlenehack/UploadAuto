func main() {
    excelFileName := "test.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Printf("open failed: %s\n", err)
    }
    for _, sheet := range xlFile.Sheets {
        fmt.Printf("Sheet Name: %s\n", sheet.Name)
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
                text := cell.String()
                fmt.Printf("%s\n", text)
            }
        }
    }
}


func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 创建表头
	row := sheet.AddRow()
	row.SetHeightCM(1)  // 设置行高
	cellA := row.AddCell()
	cellA.Value = "ID"
	cellB := row.AddCell()
	cellB.Value = "姓名"

	//依次添加内容
	for i := 0; i <= 10; i++ {
		row := sheet.AddRow()
		cellA := row.AddCell()
		cellA.Value = strconv.Itoa(i)
		cellB := row.AddCell()
		cellB.Value = "tom"
	}

	err = file.Save("1.xlsx")
	if err != nil {
        	fmt.Printf(err.Error())
    	}
}


func main() {
	file, _ := xlsx.OpenFile("1.xlsx")
	sheet := file.Sheets[0]   // file.Sheets返回一个[]*sheet，有多个sheet可循环获取
	for i := 10; i <= 20; i++ {
		row := sheet.AddRow()
		cellA := row.AddCell()
		cellA.Value = strconv.Itoa(i)
		cellB := row.AddCell()
		cellB.Value = "tim"
	}
	file.Save("1.xlsx")
}
