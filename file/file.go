package file
import (
	"io/ioutil"
	// "fmt"
	"os"
)
 
 
func Create(fileName string,nr string) {
   content := []byte(nr)
   err := ioutil.WriteFile(fileName, content, 0644)
   if err != nil {
      panic(err)
   }
}
func Reader(fileName string) string{
    
   file, err := os.Open(fileName)
   if err != nil {
      panic(err)
   }
   defer file.Close()
   content, err := ioutil.ReadAll(file)
  // fmt.Println(string(content))
  result := string(content)
   return result
    
}

/*
golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:
如果返回的错误为nil,说明文件或文件夹存在
如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
如果返回的错误为其它类型,则不确定是否在存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
*/
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息=> false :bcz true :cz
	if os.IsNotExist(err) {
		return false
	}
	return true
}













