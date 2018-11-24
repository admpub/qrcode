# go_qrcode
golang 编码,解码二维码 ，感谢 bieber的帮助  
解码需要#include <zbar.h> c语言库的支持（例如：pip install zbar）

开始要 go get github.com/admpub/qrcode  

引入第三方库 go get github.com/boombuler/barcode  

例子  
package main  

import (  
	"fmt"  
	"image/png"  
	"os"  

	"github.com/admpub/qrcode"
)

func main() {  
	err := png.EncodeToFile("test qrcode", 300, 300,"./test.png")  
	if err != nil {  
		fmt.Println(err)  
		return  
	}

	value := qrcode.DecodeFile("./test.png")  
	fmt.Println(value)
}
