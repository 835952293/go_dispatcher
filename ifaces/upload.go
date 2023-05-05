package ifaces

// import (
// 	"fmt"
// 	"log"
// 	"mime/multipart"
// 	"os"
// 	"path"
// 	"strings"

// 	"github.com/EDDYCJY/go-gin-example/pkg/file"
// 	"github.com/EDDYCJY/go-gin-example/pkg/logging"
// 	"go_trancation_center/config"
// 	"github.com/EDDYCJY/go-gin-example/pkg/util"
// )

// // GetImageFullUrl get the full access path
// func GetImageFullUrl(name string) string {
// 	return config.Appconfig.PrefixUrl + "/" + GetImagePath() + name
// }

// // GetImageName get image name
// func GetImageName(name string) string {
// 	ext := path.Ext(name)
// 	fileName := strings.TrimSuffix(name, ext)
// 	fileName = util.EncodeMD5(fileName)

// 	return fileName + ext
// }

// // GetImagePath get save path
// func GetImagePath() string {
// 	return config.Appconfig.ImageSavePath
// }

// // GetImageFullPath get full save path
// func GetImageFullPath() string {
// 	return config.Appconfig.RuntimeRootPath + GetImagePath()
// }

// // CheckImageExt check image file ext
// func CheckImageExt(fileName string) bool {
// 	ext := file.GetExt(fileName)
// 	for _, allowExt := range config.Appconfig.ImageAllowExts {
// 		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
// 			return true
// 		}
// 	}

// 	return false
// }

// // CheckImageSize check image size
// func CheckImageSize(f multipart.File) bool {
// 	size, err := file.GetSize(f)
// 	if err != nil {
// 		log.Println(err)
// 		logging.Warn(err)
// 		return false
// 	}

// 	return size <= config.Appconfig.ImageMaxSize
// }

// // CheckImage check if the file exists
// func CheckImage(src string) error {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		return fmt.Errorf("os.Getwd err: %v", err)
// 	}

// 	err = file.IsNotExistMkDir(dir + "/" + src)
// 	if err != nil {
// 		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
// 	}

// 	perm := file.CheckPermission(src)
// 	if perm == true {
// 		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
// 	}

// 	return nil
// }
