type results struct {
	Kind string
	Name string
	Stars string
	Title string
	Date string
	Color string
	Content string
}

var head = map[string]string{
	"A1": "用户名",
	"B1": "星级",
	"C1": "标题",
	"D1": "日期",
	"E1": "属性",
	"F1": "内容",
}

func writeToExcel(ASIN string, r *results)  {
	fileName := filepath.Join("results", ASIN + "_" + r.Kind + ".xlsx")
	// 文件不存在则创建，并写入表头
	yes, _ := os.Stat(fileName)
	if yes == nil {
		xlsx := excelize.NewFile()
		index := xlsx.NewSheet("Sheet1")
		for k, v := range head {
			xlsx.SetCellValue("Sheet1", k, v)
		}
		xlsx.SetActiveSheet(index)
		err := xlsx.SaveAs(fileName)
		if err != nil {
			fmt.Printf("创建文件失败，ASIN -> %v，err -> %v\n", ASIN, err)
		}
	}

	// 获取Excel行数
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Printf("读取文件错误，ASIA -> %v，err -> %v\n", ASIN, err)
	}
	rows, err := xlsx.GetRows("Sheet1")

	// 准备数据
	rowNum := strconv.Itoa(len(rows) + 1)
	data := map[string]string{
		"A" + rowNum: r.Name,
		"B" + rowNum: r.Stars,
		"C" + rowNum: r.Title,
		"D" + rowNum: r.Date,
		"E" + rowNum: r.Color,
		"F" + rowNum: r.Content,
	}
	// 写入数据
	for k, v := range data {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	err = xlsx.Save()
	if err != nil {
		fmt.Printf("写入文件失败，ASIN -> %v，err -> %v\n", ASIN, err)
	}
}
