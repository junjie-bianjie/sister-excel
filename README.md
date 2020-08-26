## sister-excel使用教程

例如我们在test目录下有两个excel文件 困难户名单.xlsx  缺.xlsx
同时要比对的是缺.xlsx的sheet为 `Sheet1` 表中的公司名,困难户名单.xlsx的sheet为 `缺` 表中的公司名
那么yaml如下

**请闹闹仔细看本demo**

**file 是文件目录 e.g. /C/a.txt   /D/go/sister.md**

**sheetBook 是excel中的工作薄 e.g. Sheet1**

**注意file和sheetBook一定要对上！！！！！！！！！！不知道sheetBook是什么请打开excel看**

**注意文件名中间不能出现空格等特殊字符 ！！！！**

**./是指当前目录  所以项目下的config目录 -> ./config ,  项目下的excel目录 -> ./excel**

```yaml
excel1:
  file: "./excel/缺.xlsx"
  sheetBook: "Sheet1"

excel2:
  file: "./excel/困难户名单.xlsx"
  sheetBook: "缺"

# 运行正确会在命令行中给出程序运行正确的提示
```