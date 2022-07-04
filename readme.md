
```
go mod init github.com/copiner/gogz

go mod tidy


go list -m github.com/copiner/gogz@v0.1.0
```

```
go env
    
```


go get github.com/copiner/gogz

```
package main

import(
	"fmt"
	"github.com/copiner/gogz/helper"
)

func main(){
	str := "hello world"
	fmt.Println(helper.Md5(str))
}
```