#### 2020.5.26
目标：06
进度：完成
笔记：
1. 强类型的转换需要注意 int→string，float→int
2. golang 支持各个平台交叉编译 测试了Linux生成Windows的命令：`CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o myprogram.exe`

#### 2020.5.27
目标：07
进度：完成
笔记：
1. 使用不同类型：%s→string，%d→int，%f→float，%t→boll
2. %v可以通用

#### 2020 6.3
目标：08
进度：完成
笔记：
1. 取余数的计算符号%只能用于整型int
2. 可用使用`strings.Repeat`方法来循环输出一个字符串，如`marks := strings.Repeat("!", 2)`则marks为'!!'
3. ```len(`\"cool\"`)```输出的数量是8
4. 复杂文字统计长度`utf8.RuneCountInString("péripatéticien")`

#### 2020 6.4
目标：09
进度：01~04
笔记：
1. `strconv.ParseInt`字符串转换为Int类型数据，可以把输入的string为2进值，8进值，16进值
2. uint 类型代表正整数;09/predeclared-type查看每个类型的内存消耗
3. 溢出在编译的时候会捕获错误，但是在运行时溢出是不会捕获的，会绕回到反向最极值去计算如int8最大为127，如果运行时+1了就变成-128，+3就是-126；float类型溢出时会直接变成无穷大
4. 用高位数转低位溢出的时候和其他语言一样只会转换为低位数的最大值