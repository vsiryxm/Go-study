package json

import (
	"encoding/json"
	//"github.com/json-iterator/go"
	"fmt"
	"io"
	"os"
	"strings"
)

//json转结构体
func JsonToStruct() {
	const jsonStream = `
		{"id": 1, "name": "海阳之新"}
		{"id": 2, "name": "欧阳"}
		{"id": 3, "name": "Simon"}
		{"id": 4, "name": "欧阳新民"}
		{"id": 5, "name": "海阳"}
	` //里面大小写都可以解析正确，如NaMe
	//从一个输入流中读取并进行解码json的值
	//初始化一个Decoder对象
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		//创建的decoder对象的Decode方法可以将内容解析到接口v中
		//func (dec *Decoder) Decode(v interface{}) error

		//读取到末尾和读取错误处理
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d: %s\n", m.Id, m.Name)
		}
	}
}

var myFamily = MyFamily{
	Address: "湖南新化梅苑开发区",
	Member: []Person{
		Person{
			Name:    "海阳之新",
			Sex:     "男",
			Age:     35,
			IsMarry: true,
			Phone: map[string][]string{
				"china_mobile":  []string{"13771739166", "15916689566"},
				"china_telecom": []string{"18971739166", "18916689566"},
			},
		},
	},
}

// 结构体转json
func StructToJson() {
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	i:=0
	for {
		bytes, _ := json.Marshal(myFamily)
		fmt.Printf("结构体转json：%s\n", bytes)
		if i>10000  {
			break
		}
		i++
	}

}

//读取json文件
func ReadJsonFile() {
	filePtr, err := os.Open("my_family.json") //json文件中不能有注释
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return
	}
	defer filePtr.Close()

	var myFamily MyFamily

	//创建json解码器
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&myFamily)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
	} else {
		fmt.Println("Decoder success")
		jsonShow, _ := json.MarshalIndent(myFamily, "", "    ")
		fmt.Println("解码结果为：\n", string(jsonShow))
	}
}

//写入json文件
func WriteJsonFile() {

	// 创建文件
	filePtr, err := os.Create("my_home.json")
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json编码器
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(myFamily)
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("写入my_home.json成功！")
	}

	//以下代码可选======================================================
	//带缩进格式输出
	filePtr, err = os.Create("my_home_format.json")
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	//defer filePtr.Close()

	data, err := json.MarshalIndent(myFamily, "", "    ") //第二个参数将在json的第二行开始至末尾添加前缀，第三个参数通常为缩进空格
	if err != nil {
		fmt.Println("Encoder failed", err.Error())
	} else {
		fmt.Println("写入my_home_format.json成功！")
	}
	filePtr.Write(data)
}



type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func JsonX()  {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// And here are some for slices and maps, which encode
	// to JSON arrays and objects as you'd expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your
	// custom data types. It will only include exported
	// fields in the encoded output and will by default
	// use those names as the JSON keys.
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
	Page:   1,
	Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Now let's look at decoding JSON data into Go
	// values. Here's an example for a generic data
	// structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.
	var dat map[string]interface{}

	// Here's the actual decoding, and a check for
	// associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
	panic(err)
	}
	fmt.Println(dat)

	// In order to use the values in the decoded map,
	// we'll need to cast them to their appropriate type.
	// For example here we cast the value in `num` to
	// the expected `float64` type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of
	// casts.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional
	// type-safety to our programs and eliminating the
	// need for type assertions when accessing the decoded
	// data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and
	// strings as intermediates between the data and
	// JSON representation on standard out. We can also
	// stream JSON encodings directly to `os.Writer`s like
	// `os.Stdout` or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

type Person struct {
	Name    string              `json:"name"`
	Sex     string              `json:"sex"`
	Age     int                 `json:"age"`
	IsMarry bool                `json:"is_marry"`
	Phone   map[string][]string `json:"phone"`
}

type MyFamily struct {
	Address string   `json:"address"`
	Member  []Person `json:"member"`
}

type Message struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}