func main() {
    f := excelize.NewFile()
    // 创建一个工作表
    index := f.NewSheet("Sheet1")
    // 设置单元格的值
    f.SetCellValue("Sheet1", "A1", "Hello world.")
    f.SetCellValue("Sheet1", "B1", 100)
    // 设置工作簿的默认工作表
    f.SetActiveSheet(index)
    // 根据指定路径保存文件
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        fmt.Println(err)
    }
}

func main() {
    f, err := excelize.OpenFile("Book1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 获取工作表中指定单元格的值
    cell, err := f.GetCellValue("Sheet1", "B2")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 获取 Sheet1 上所有单元格
    rows, err := f.GetRows("Sheet1")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}

func main() {
   f, err := excelize.OpenFile("Book1.xlsx")
   f.SetCellValue("Sheet1", "A2", "wow")
	f.SetCellValue("Sheet1", "B2", "1407")
   if err := f.Save(); err != nil {
       fmt.Println(err)
    }
}

import (
    "fmt"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"

    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    f, err := excelize.OpenFile("Book1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 插入图片
    if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
        fmt.Println(err)
    }
    // 在工作表中插入图片，并设置图片的缩放比例
    if err := f.AddPicture("Sheet1", "D2", "image.jpg", `{
        "x_scale": 0.5,
        "y_scale": 0.5
    }`); err != nil {
        fmt.Println(err)
    }
    // 在工作表中插入图片，并设置图片的打印属性
    if err := f.AddPicture("Sheet1", "H2", "image.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
        fmt.Println(err)
    }
    // 保存文件
    if err = f.Save(); err != nil {
        fmt.Println(err)
    }
}
   
