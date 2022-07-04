
```
go mod init github.com/copiner/gogz

go mod tidy


```

```
go env
    
```
```
git tag v0.1.2

git push origin v0.1.2
```

```
go list -m github.com/copiner/gogz@v0.1.2
```


```
go get github.com/copiner/gogz
    
```

```
/Users/wdaonngg/go/src/github.com/copiner/gogz    
```
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
    