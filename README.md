# json2GoStruct
Convert JSON to Go Struct,把json转换为Go的结构体

## 安装
```bash
go get -v github.com/dravinbox/json2GoStruct
```
## 快速使用

### Web模式
```bash
$ json2GoStruct -w
```

### 命令行模式

```bash

$ cat << EOF >./file.json
{"name": "zry", "Age": 18,"car": { "car_name": "benzi","color": "red" }}
EOF

$ json2GoStruct


```

指定的json文件路径,指定Struct名称
```bash
$ json2GoStruct -f <filePath> -n Dravin
```