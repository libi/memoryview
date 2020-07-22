GoMemoryView 
====
Golang内存视图查看器,可以动态输出go语言各类型变量在内存中的数据状态及布局.

### todo
- 切片类型底层数组输出
- Map类型输出 

```go
	a := 10
	fmt.Println(MemView(&a))
    // [10 0 0 0 0 0 0 0]
    // int 在64位系统下占用8个字节 
	
	b := int32(257)
	fmt.Println(MemView(&b))
    // [1 1 0 0] 
    // int32 占用4个字节，单字节上限256，逢256进1
	
	c := [3]int{1,2,3}
	fmt.Println(MemView(&c))
    // [1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0] 
    // 长度为3 类型为int的数组 占用空间8*3字节 

	d := c[2:3]
	fmt.Println(MemView(&d))
    // [16 224 12 0 192 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0] 
	// 切片类型其实是一个包含 底层数组指针unsafe.Pointer，len，cap的结构体
    // 在64位系统下 unsafe.Pointer ，int 占用8个字节
    // 所以切片本身结构体占用24字节

	e := "abcd"
	fmt.Println(MemView(&e))
    // [43 107 20 1 0 0 0 0 4 0 0 0 0 0 0 0] 
    // 字符串底层数据类型为 []rune(int32) 所以该字符串内存占用 4*4 字节
```

