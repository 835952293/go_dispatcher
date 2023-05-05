// package ifaces

// import "go_trancation_center/config"

// const EXT = ".xlsx"

// // GetExcelFullUrl get the full access path of the Excel file
// func GetExcelFullUrl(name string) string {
// 	return config.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
// }

// // GetExcelPath get the relative save path of the Excel file
// func GetExcelPath() string {
// 	return config.AppSetting.ExportSavePath
// }

// // GetExcelFullPath Get the full save path of the Excel file
// func GetExcelFullPath() string {
// 	return config.AppSetting.RuntimeRootPath + GetExcelPath()
// }
