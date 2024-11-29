
/Users/user/code/github.com/Explorer1092/nuclei/pkg/input/provider/list/hmap.go

// 添加自定义IP的支持

```golang

    if stringsutil.ContainsAny(value, ",") {
        parts := strings.Split(value, ",")
        // 处理2列的情况: CustomIP,Host
        if len(parts) == 2 {
            if iputil.IsIP(strings.TrimSpace(parts[0])) {
                metaInput := contextargs.NewMetaInput()
                metaInput.CustomIP = strings.TrimSpace(parts[0])
                metaInput.Input = strings.TrimSpace(parts[1])
                i.setItem(metaInput)
                return
            }
        }
    }
	
```