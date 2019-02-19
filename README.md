### Hello World
```
import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

### 变量定义
var定义变量
```
func variableZeroValue() {
	var i int
	var s string
	fmt.Printf("%d %q \n", i, s) // 0 ""
}
```
同时定义多个
```
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s) // 3 4 abc
}
```
类型推导
```
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s) // 3 4 true def
}
```
:= ，只可以用在函数内部
```
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s) // 3 4 true def
}

func main() {
	var a = 2
	a, b, c, s := 3, 4, true, "def" // a被覆盖
	fmt.Println(a, b, c, s) // 3 4 true def
}

```
作用域为包内部
```
var (
	a = 3
	s = "string"
	b = true
)

func main() {
	fmt.Println(a, s, b) // 3 string true
}
```

### 内建变量类型
```
bool，string
```
```
(u)int，(u)int8，(u)int16，(u)int32，(u)int64，uintptr

u，无符号
有符号，规定长度，不规定长度按系统来
uintptr， 指针，长度按系统来
```
```
byte，8位
rune，32位

rune，字符型，理解成go语言的char类型
```
```
float32，float64，complex64，complex128

complex，复数类型，分实部和虚部
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) // 5
} 

func euler() {
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1) // (0.000+0.000i)
}
```

### 强制类型转换
go语言没有隐式转换
```
勾股定理

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	println(c) // 5
}
```
### 常量
const数值可以作为各种类型使用
```
const filename = "abc.txt" // 可以定义在包内部

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b)) // 不需要强转类型
	println(filename, c) // abc.txt 5
}
```

### 枚举类型
```
func enums() {
	const (
		java   = 0
		python = 1
		golang = 2
	)
	fmt.Println(java, python, golang) // 0 1 2
}
```
```
func enums() {
	const (
		java   = iota
		_
		python
		golang
	)
	fmt.Println(java, python, golang) // 0 2 3
}


func enums() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(b, kb, mb, gb) // 1 1024 1048576 1073741824
}
```

### 条件查询：if
if的条件可以赋值<br>
if的条件里赋值的变量作用域就在这个if语句里
```
const filename = "abc.txt"
contents, err := ioutil.ReadFile(filename)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("%s\n", contents)
}
```
```
const filename = "abc.txt"
if contents, err := ioutil.ReadFile(filename); err != nil { // contents作用域在if语句里
	fmt.Println(err)
} else {
	fmt.Printf("%s\n", contents)
}
```

### 条件查询：switch
switch会自动break，除非使用fallthrough
```
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("未定义的操作：" + op)
	}
	return result
}
```
switch后可以没有表达式
```
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("错误的分数：%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
```

### 循环：for
for的条件里不需要括号<br>
for的条件里可以省略初始条件，结束条件，递增表达式
```
sum := 0
for i := 1; i <= 100; i++ {
	sum += i
}
```
```
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result  // int转string
	}
	return result
}
```
省略初始条件，相当于while
```
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // bool
		fmt.Println(scanner.Text())
	}
}
```
```
func forever() {
	for {
		fmt.Println("死循环")
	}
}
```

### 函数
```
// 带余除法
func div(a, b int) (int, int) {
	return a / b, a % b
}

div(10, 3) // 3 1
```

_：省略变量
```
func div(a, b int) (q, r int) { // 给返回值起名
	return a / b, a % b
}
q, r := div(10, 3)
q, _ := div(10, 3) // 只要第一个返回值q
```

```
// 不建议用
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
```

常用场景
```
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("未知操作：%s", op)
	}
}

eval(1, 2, "+") // 3 <nil>
eval(1, 2, "%") // 0 未知操作：%
```

函数式改写
```
func main() {
	fmt.Println(
		apply(
			func(a int, b int) int { // 这是一个匿名函数
				return int(math.Pow(float64(a), float64(b)))
			},
			3,
			4)) // 函数名：main.main.func1，参数：3，4，结果：81
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("函数名：%s，参数：%d，%d，结果：", opName, a, b)
	return op(a, b)
}		
```
可变参数
```
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

sum(1, 2, 3) // 6
```

### 指针
指针不能运算
```
func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a) // 3
}
```
### 参数传递
GO语言只有值传递一种方式
```
func main() {
	a, b := 3, 4
	swap(&a, &b)
	println(a, b) // 4 3
}

// 不建议这么写
func swap(a, b *int) {
	*a, *b = *b, *a
}
```

### 数组
```
var arr1 [3]int            // [0 0 0]
arr2 := [3]int{1, 2, 3}    // [1 2 3] 
arr3 := [...]int{4, 5, 6}  // [4 5 6] 
var arr4 [2][3]int         // [[0 0 0] [0 0 0]]
```

```
for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i]) // 1 2 3
}

for i := range arr2{
	fmt.Print(arr2[i]) // 1 2 3
}

for i, v := range arr2 {
	fmt.Print(i, v)
}

结果:  0 1
       1 2
       2 3
```
数组是值类型<br>
[10]int 和 [20]int 是不同类型<br>
调用 func f(arr [10]int) 会拷贝数组
```
func main() {
	arr := [3]int{1, 2, 3}
	updateArray(arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func updateArray(arr [3]int){
	arr[0] = -1
}
结果:  0 1 // 没有被改变
       1 2
       2 3
```
不推荐使用，推荐使用切片
```
func main() {
	arr := [3]int{1, 2, 3}
	updateArray(&arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func updateArray(arr *[3]int){
	arr[0] = -1
}

结果：0 -1
      1 2
      2 3
```

### Slice（切片）
```
arr := [...]int{0, 1, 2, 3, 4, 5, 6}

arr[2:4]   // [2 3]
arr[:4]    // [0 1 2 3]
arr[2:]    // [2 3 4 5 6]
arr[:]     // [0 1 2 3 4 5 6]
```
slice本身是没有数据的，是对底层array的一个view
```
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s := arr[:2]
	updateSlice(s)
	fmt.Println(arr) // [-1 1 2 3 4 5 6]
}

func updateSlice(s []int) {
	s[0] = -1
}
```
向切面添加值要传切面指针
```
func main() {
	arr1 := make([]int, 0)
	add(arr1)
	fmt.Printf("%d", arr1) // []，扩容导致产生了新数组

}

func add(i []int) {
	i = append(i, 1)
}
```
Reslice
```
arr := [...]int{0, 1, 2, 3, 4, 5, 6}
s := arr[:] // [0 1 2 3 4 5 6]
s = s[1:] // [1 2 3 4 5 6]
s = s[:5] // [1 2 3 4 5]
```
slice可以向后扩展，不可以向前扩展<br>
s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
```
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s := arr[:] // [0 1 2 3 4 5 6 7]
s = s[2:6] // s=[2 3 4 5], len(s)=4, cap(s)=6
s = s[3:5] // [5 6]，这里取到了s[5]
```
添加元素时如果超越cap，系统会重新分配更大的底层数组<br>
由于是值传递的关系，必须接收append的返回值，s = append(s, val)
```
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6] // s1=[2 3 4 5] 
s2 := s1[3:5] // s2=[5 6]
s3 := append(s2, 10) // s3=[5 6 10] ，这里是10替换7
s4 := append(s3, 11) // s4=[5 6 10 11] ，s4和s5是对新数组的view
s5 := append(s4, 12) // s5=[5 6 10 11 12] 
```
```
func printSlice(s []int) {
	fmt.Printf("s=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

var s []int // s == nil
for i := 0; i < 10; i++ {
	s = append(s, 2*i+1)
}  

printSlice(s) // s=[1 3 5 7 9 11 13 15 17 19], len=10, cap=16
s1 := []int{2, 4, 6, 8} // s=[2 4 6 8], len=4, cap=4
s2 := make([]int, 16) // s=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
s3 := make([]int, 10, 32) // s=[0 0 0 0 0 0 0 0 0 0], len=10, cap=32
```
拷贝
```
copy(s2, s1)
printSlice(s2) // s=[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
```
删除
```
删除8
s2 = append(s2[:3], s2[4:]...) // 4后面的所有元素
printSlice(s2) // s=[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15, cap=16
```

### Map
map[K]V，map[K1]map[K2]V
```
m := map[string]string{
    "name": "小明",
	"age":  "18",
	"sex":  "男",
}
fmt.Println(m) // map[age:18 sex:男 name:小明]
```
key不存在时，获取value类型的初始值
```
m := make(map[string]int) // m == 空map
var m map[string]int // m == nil，go语言的nil可以参与运算
fmt.Println(m) // map[]
```
遍历（key是无序的）
```
for k, v := range m {
	fmt.Println(k, v)
}
结果：  age 18
        sex 男
        name 小明
```
取值
```
name := m["name"] // 小明
name := m["不存在的key"] // zero value，空串

常用
if name, ok := m["name"]; ok {
    fmt.Println(name, ok) // 小明 true
}
```
删除
```
name, ok := m["name"]
delete(m, "name")
name, ok = m["name"]
fmt.Println(name, ok) // （这里有个空串） false
```
map的key
```
map使用哈希表，必须可以比较相等
除了slice，map，function的内建类型都可以作为key
Struct类型不包含上述字段，也可以作为key
```
例：寻找最长不含有重复字符的子串<br>
思路：<br>
对于每一个字母x<br>
lastOccurred[x]不存在，或者<start，无需操作<br>
lastOccurred[x] >= start，更新start<br>
更新lastOccurred[x]，更新maxLength
```
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) { // utf8解码
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
```

### 字符串
```
s := "Hi你好!" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // 48 69 E4 BD A0 E5 A5 BD 21
	}
	// s是UTF-8转成ch是Unicode
	for i, ch := range s { // ch是rune
		fmt.Printf("(%d %X) ", i, ch) // (0 48) (1 69) (2 4F60) (5 597D) (8 21) ，i不连续
}
```
获取字符数量
```
utf8.RuneCountInString(s) // 5
```
```
bytes := []byte(s)
for len(bytes) > 0 { // len(bytes)，字节数
	ch, size := utf8.DecodeRune(bytes)
	bytes = bytes[size:]
	fmt.Printf("%c ", ch) // H i 你 好 ! 
}

for i, ch := range []rune(s) { // 重新开的rune数组
	fmt.Printf("(%d %c)", i, ch) // (0 H)(1 i)(2 你)(3 好)(4 !)
}
```

### 结构体
```
go语言仅支持封装，不支持继承和多态<br>
go语言没有class，只有struct
```
结构的定义
```
type treeNode struct {
	value       int
	left, right *treeNode
}
```
生成对象，赋值
```
var root treeNode // {0 <nil> <nil>}
root = treeNode{value: 3}
root.left = &treeNode{}
root.right = &treeNode{5, nil, nil}
root.right.left = new(treeNode) // new创建
```
```
nodes := []treeNode{
	{value: 3},
	{}, 
	{6, nil, &root},
}
fmt.Println(nodes) // [{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000044400}] ， 默认0
```
工厂方法
```
func createTreeNode(value int) *treeNode {
	return &treeNode{value: value} // 局部变量的地址
}

root.left.right = createTreeNode(2)
```
显示定义和命名方法接收者
```
func (node treeNode) print() {
	fmt.Print(node.value)
}

var root treeNode // {0 <nil> <nil>}
root.print() // 0
```
set方法
```
func (node treeNode) setValue(value int) {
	node.value = value
}

var root treeNode // {0 <nil> <nil>}
root.setValue(4)
root.print() // 0 ， 这里没有set进去

改成

func (node *treeNode) setValue(value int) { // 多了*，传引用
	node.value = value
}
```
址与值传参时系统会做转换<br>
```
var root treeNode // {0 <nil> <nil>}
pRoot := &root // pRoot是地址
pRoot.print() // 0
```
nil指针也可以调用方法
```
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("node为nil，node.value会报错")
		return
	}
	node.value = value
}

var pRoot *treeNode
pRoot.setValue(3) // node.value会报错
```
遍历
```
func (node *treeNode) traverse(){
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}
```
值接收者 vs 指针接收者
```
要改变内容必须使用指针接收者
结构过大也考虑使用指针接收者
一致性：如有指针接收者，最好都是指针接收者
```
 ### 封装
 名字一般使用CamelCase<br>
 首字母大写：public （相对于包而言）<br> 
 首字母小写：private<br>
 
 ### 包
 ```
 每个目录一个包
 main包包含可执行入口
 为结构定义的方法必须放在同一个包内
 可以是不同文件
 ```
 ```
package tree 

type Node struct {
    Value       int
    Left, Right *Node
}

-------------------------------

package main

import "learngo/tree" // 导入tree包

func main() {
    var root tree.Node 
}	
 ```
 
### 扩展已有类型
##### 使用组合
原有类，先序遍历
 ```
 func (node *Node) Traverse(){
    if node == nil{
        return
    }
    node.Left.Traverse()
    node.Print()
    node.Right.Traverse()
}
```
扩展类，后序遍历
```
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

myRoot := MyTreeNode{&root}
myRoot.postOrder()
```
 ##### 定义别名
 ```
type D = int  // 类型别名
type I int    // 类型声明

func main() {
    v := 100
    var d D = v  // 不报错
    var i I = v  // 报错
}
 ```
 
 ```
package queue

type Queue []int

func (q *Queue) Push(v int) {
    *q = append(*q, v)
}

func (q *Queue) Pop() int {
    head := (*q)[0]
    *q = (*q)[1:]
    return head
}

func (q *Queue) IsEmpty() bool {
    return len(*q) == 0
}

---------------------------------------

func main()  {
    q := queue.Queue{1}
    q.Push(2)
    println(q.Pop())  // 1
    println(q.IsEmpty()) // false
    println(q.Pop()) // 2
    println(q.IsEmpty()) // true
}

这里的q已经被改变了
 ```
  
 ### GOPATH环境变量
 ```
 默认（unix，linux）：~/go
     （windows）：%USERPROFILE%\go
 官方推荐：所有项目和第三方库都放在同一个GOPATH下
 实际项目中：也可以将每个项目放在不同的GOPATH
 ```
 ```
 src bin pkg
 ```
 go get 获取第三方库
 ```
翻墙：go get golang.org/x/tools/cmd/goimports

github：
    go get -v github.com/gpmgo/gopm
    gopm get -g -v -u golang.org/x/tools/cmd/goimports
    go build golang.org/x/tools/cmd/goimports
    rm goimports 
    切换到GOPATH，go install golang.org/x/tools/cmd/goimports
 ```
 命令
 ```
 go build 编译
 go install产生pkg文件和可执行文件
 go run 直接编译运行
 ```
 目录结构
 ```
 一个package下只能有一个main
 ```
 ### duck typing
 ```
 描述事物的外部行为而非内部结构
 ```
 
### 接口
接口由使用者定义
```
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.imooc.com")
}

func main() { 
	var r Retriever
	r = mock.Retriever{"www.baidu.com"}
	fmt.Println(download(r)) // www.baidu.com
}
```
接口的实现是隐式的<br>
只要实现接口里的方法
```
package mock

type Mock struct {
	Contents string
}

func (m Mock) Get(url string) string {
    // 用url和m做一些事情
	return "mock"
}
```
指针方式
```
package real

type Real struct {

}

func (r *Real) Get(url string) string {
    // 用url做一些事情
	return "*real"
}

func main() {
	var r Retriever
	r = &real.Real{}
	fmt.Printf("%T %v\n", r, r) // *real.Real &{}
}
```

### 接口的值类型
判断类型
```
func inspect(r Retriever) {
	fmt.Printf("%T %v", r, r)
	switch v := r.(type) {
	case mock.Mock:
		fmt.Println("mock：", v.Contents)
	case *real.Real:
		fmt.Println("real：", v)
	}
}
```
类型断言
```
r = mock.Mock{}
mockRetriever := r.(mock.Mock)
fmt.Printf("%T %v\n", mockRetriever, mockRetriever) // mock.Mock {}

r = &real.Real{}
realRetriever := r.(*real.Real)
fmt.Printf("%T %v\n", realRetriever, realRetriever) // *real.Real &{}

```
防止错误
```
realRetriever := r.(real.Real) // 没有*报错
```
```
func main() {
	var r Retriever
	r = &real.Real{}
	if mockRetriever, ok := r.(mock.Mock); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("不是mockRetriever")
	}
}

结果：不是mockRetriever
```
接口变量里有什么
```
接口变量自带指针
接口变量同样采用值传递，几乎不需要使用接口的指针
指针接收者实现只能以指针的方式使用；值接收者都可
```
##### interface{}
表示任何类型
```
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
    return head.(int) // 强制转换，只能pop，int类型
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
```

### 接口的组合
```
package main

type Geter interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, from map[string]string)
}

type GeterPoster interface { // 这里把两个接口组合起来
	Geter
	Poster
	// 可以加自己定义的方法
}

func session(s GeterPoster) string {
	s.Post("", map[string]string{"name": "小红"})
	return s.Get("")
}

func main() {
	r := &mock.Retriever{"小明"}
	println(session(r)) // 小红
}

----------------------------------
package mock

type Retriever struct {
	Contents string
}

func (r *Retriever) Post(url string, from map[string]string) {
	r.Contents = from["name"]
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
```

### 常用系统接口
```
Stringer Reader Writer
```
Reader
```
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	s := `1
			2`
	printFileContents(strings.NewReader(s))
}

结果：1
			2
```
### 函数式编程
```
参数，变量，返回值都可以是函数
```
"正统"函数式编程
```
不可变性：不能有状态，只有常量和函数
函数只有一个参数
```
闭包
```
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}

结果：
    0 + 1 + ... + 0 = 0
    0 + 1 + ... + 1 = 1
    0 + 1 + ... + 2 = 3
    0 + 1 + ... + 3 = 6
    0 + 1 + ... + 4 = 10
    0 + 1 + ... + 5 = 15
    0 + 1 + ... + 6 = 21
    0 + 1 + ... + 7 = 28
    0 + 1 + ... + 8 = 36
    0 + 1 + ... + 9 = 45
```
改写
```
type iAdder func(int) (int, iAdder)

func adder(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder(base + v)
	}
}

func main() {
	a := adder(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
```
##### 返回值
斐波那契数列
```
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	fmt.Println(f(), f(), f(), f(), f(), f(), f(), f(), f())
}
```

##### 变量
为函数生成接口
```

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF 
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printFileContents(f)
}
```
##### 参数
使用函数来遍历二叉树
```
package tree

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}
```
```
package tree

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
```
```
package main

root := tree.Node{Value: 0, Left: &tree.Node{Value: 2, Left: &tree.Node{Value: 1}}, Right: &tree.Node{Value: 4, Right: &tree.Node{Value: 5}}}
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount ++
	})
	fmt.Println("节点数：", nodeCount) // 5
	root.Traverse() // 1 2 0 4 5
	
```

### defer
确保调用在函数结束时发生<br>
参数在defer语句时计算<br>
defer列表为后进先出
```
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func main() {
	tryDefer()
}

结果：
    3
    2
    1

```
```
func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}
}
结果：
    9 8 7 6 5 4 3 2 1 0  // 这里是倒叙的，把它理解成一个栈
```
使用场景
```
Open/Close
Lock/Unlock
PrintHeader/PrintFooter
```
### 错误处理
```
func main() {
	_, err := os.Open("abc.txt")
	err = errors.New("自定义打开文件异常")
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError)
		} else {
			panic(err)
		}
	}
}
```

### panic（少用）
停止当前函数执行<br>
一直向上返回，执行每一层的defer<br>
如果没有遇见recover，程序退出

### recover
仅在defer调用中使用<br>
获取panic的值<br>
如果无法处理，可重新panic
```
func tryRecover() {
	defer func() { // 这是一个匿名函数
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("错误信息：", err)
		} else {
			panic(r)
		}
	}()
	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover() // 系统没有崩溃，而是打出了错误信息
}
```
### 表格驱动测试
文件后缀名_test<br>
代码覆盖率<br>
pprof优化性能
```
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("%d的平方 + %d的平方 = %d的平方 ，实际结果 = %d的平方", tt.a, tt.b, tt.c, actual)
		}
	}
}
```
性能测试
```
func Benchmark(b *testing.B) {
	// 准备数据操作
	b.ResetTimer() // 不计入时间
	for i := 0; i < b.N; i++ {
		actual := calcTriangle(3, 4)
		if actual != 5 {
			fmt.Println("计算错误")
		}
	}
}
```

### goroutine
```
任何函数只需加上go就能送给调度器运行
不需要在定义时区分是否是异步函数
调度器在合适的点进行切换 
```
```
goroutine可能切换的点（参考）：
I/O，select
channel
等待锁
函数调用（有时）
runtime.Gosched()
```
```
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // 匿名函数
			for {
				a[i] ++
				runtime.Gosched() // 主动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}

理解：
    i要用参数传进来，不传报错：index out of range，当i被累加到10时，i++ = 11，a[11]越界
```

### 协程Goroutine
子程序是协程的一个特例<br>
轻量级"线程"<br>
非抢占式多任务处理，由协程主动交出控制权<br>
编译器/解释器/虚拟机层面的多任务<br>
多个协程可能在一个或多个线程上运行

### channel
channel是goroutine之间双向的通道<br>
不要通过共享内存来通信；通过通信来共享内存
```
func channel() {
	// var c chan int // c == nil
	c := make(chan int)
	go worker(c)
	c <- 1
	c <- 2
}

func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func main() {
	channel() // 1  2      
}
```
参数
```
func channel() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for j := 0; j < 10; j++ {
		channels[j] <- 'a' + j
	}
	for k := 0; k < 10; k++ {
		channels[k] <- 'A' + k
	}

	time.Sleep(time.Millisecond)

}

func worker(id int, c chan int) {
	for {
		// n := <-c
		fmt.Printf("worker %d reveived %c\n", id, <-c)
	}
}

func main() {
	channel() // 大小写顺序打乱
}
```
返回值
```
func channel() {
	var channels [10]chan<- int // 可以给chan方向，chan<-只可以发数据，<-chan只可以收数据
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
		go createWorker(i)
	}
	for j := 0; j < 10; j++ {
		channels[j] <- 'a' + j
	}
	for k := 0; k < 10; k++ {
		channels[k] <- 'A' + k
	}

	time.Sleep(time.Millisecond)

}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			// n := <-c
			fmt.Printf("worker %d reveived %c\n", id, <-c)
		}
	}()
	return c
}

func main() {
	channel()
}
```
加缓冲
```
func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // deadlock!
}
```
发送方close
```
func worker(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok { // 不做这个判断系统会直接close，从而打印int默认值0
			break 
		}
		fmt.Printf("worker %d reveived %d\n", id, n)
	}
}
-------------------------------------------------更简单的写法
func worker(id int, c chan int) {
	for n := range c{
		fmt.Printf("worker %d reveived %d\n", id, n)
	}
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	channelClose()
}

结果：
    worker 0 reveived 97
    worker 0 reveived 98
    worker 0 reveived 99
    worker 0 reveived 100
```
### 使用channel等待任务结束
```
type worker struct {
	in   chan int
	done chan bool
}

func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d reveived %c\n", id, n)
		go func() {
			done <- true // 这一步是解决死锁的关键
		}()
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func channel() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
		go createWorker(i)
	}
	for j, worker := range workers {
		worker.in <- 'a' + j
		// <-worker.done // 这样会顺序打印
	}
	for k, worker := range workers {
		worker.in <- 'A' + k
		// <-worker.done
	}
	// 等待它们全部结束
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	channel()
}
```
sync.WaitGroup
```
type worker struct {
	in   chan int
	done func()
}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d reveived %c\n", id, n)
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func channel() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20) // 已知有20个
	for j, worker := range workers {
		worker.in <- 'a' + j
		// wg.Add(1) // 也可以每次加
	}
	for k, worker := range workers {
		worker.in <- 'A' + k
	}
	wg.Wait() // 等待定义的20个执行完
}

func main() {
	channel()
}
```
获取最大节点
```
func (node *Node) TraverseByChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(n *Node) {
			out <- n
		})
		close(out)
	}()
	return out
}

func main() {
	root := tree.Node{Value: 0, Left: &tree.Node{Value: 2, Left: &tree.Node{Value: 1}}, Right: &tree.Node{Value: 4, Right: &tree.Node{Value: 5}}}
	c := root.TraverseByChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("节点最大的数字是：", maxNode)  // 5
}
```

### select调度
```
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // 随机睡眠时间
			out <- i
			i++
		}
	}()
	return out
}

func main()  { // 谁先出数据选择谁
	var c1, c2 = generator(), generator()
	for {
		select{
		case n := <-c1:
			fmt.Println("from c1：", n)
		case n := <-c2:
			fmt.Println("from c2：", n)
		}
	}
}
```

```
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // 随机睡眠1.5s+
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	time.Sleep(time.Second) // 睡眠1s
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int                      // 消耗速度过慢，会比冲掉，要用数组存起来
	after := time.After(10 * time.Second) // 10s
	tick := time.Tick(time.Second)        // 定时
	for {
		var activeWorker chan<- int // 初始nil，有值才会被select到
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("超时")
		case <-tick:
			fmt.Println("队列长度：", len(values))
		case <-after:
			fmt.Println("运行10s结束")
			return
		/*
		default:
			fmt.Println("默认")
		*/
		}
	}
}
```

### 传统同步机制(很少用)
```
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	func() { // 安全块
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
```
