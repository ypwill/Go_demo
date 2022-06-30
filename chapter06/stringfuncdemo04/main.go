package main

import(
	 "fmt"
	 "strconv"
	 "strings"
)

// 字符串常用函数
func main() {
	// len 字符串长度
	str := "hello world"
	fmt.Println(len(str))

	// 不带中文的遍历
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println()

	// string带中文的遍历
	str1 := "hello中国"
	// 需要转换
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}

	fmt.Println()

	//  整数转字符串
	str2 := strconv.Itoa(43333)
	fmt.Println(str2)

	// 字符串转整数  
	n1, err := strconv.Atoi("12233")
	if err != nil {
		fmt.Println("出现出错",err)
	} else {
		fmt.Println(n1)
	}

	// 字符串转byte切片
	var byteArray = []byte("hello")
	fmt.Printf("%v", byteArray)
	fmt.Println()

	// byte[] 转成字符串
	str3 := string([]byte{104, 101,108})
	fmt.Println(str3)

	// 十进制数 转为 二、八、十六进制
	n5 := strconv.FormatInt(123, 2)
	fmt.Println(n5)
	
	// 查找子串 是否再指定的字符串中
	i := strings.Contains("seafood","ooe")
	fmt.Println(i)

	// 统计一个字符串中有几个子串
	n6 := strings.Count("hello", "l")
	fmt.Println(n6)

	// 不区分大小写的字符串比较
	b := strings.EqualFold("abc", "ABC1")
	fmt.Println(b)

	// 返回子串在字符串中第一次出现的位置
	n7 := strings.Index("hello", "ll1")
	fmt.Println(n7)

	// 返回子串 在字符串最后一次出现的位置
	n8 := strings.LastIndex("go golang", "go")
	fmt.Println(n8)

	// 将指定的子串替换为另一个子串  n表示替换几个  -1表示全部
	str4 := strings.Replace("go go hello", "go", "go语言", 1)
	fmt.Println(str4)

	// 字符串分割
	strArr := strings.Split("1,2,3,4", ",")
	fmt.Println(strArr)

	// 大小写转换
	str5 := strings.ToLower("GO")
	fmt.Println(str5)

	str6 := strings.ToUpper("go")
	fmt.Println(str6)

	// 去掉左右两边的空格
	str7 := strings.TrimSpace("   faf  ")
	fmt.Println(str7)

	// 将字符串左右两边指定的字符去掉
	str8 := strings.Trim(" !     hello woeld !  ", " !")
	fmt.Println(str8)


	str9 := strings.TrimLeft("  !  hellofsa  !", " ")
	fmt.Println(str9)

	str10 := strings.TrimRight("  !  hellofsa  !", " ")
	fmt.Println(str10)

	// 判断字符串是否以指定的字符串开头
	c := strings.HasPrefix("http://123", "http1")
	fmt.Println(c)

	// 判断字符串是否以指定的字符串结尾
	d := strings.HasSuffix("a.jpg", "jpg1")
	fmt.Println(d)


}