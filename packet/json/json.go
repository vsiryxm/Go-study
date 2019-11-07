package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func Json() {

	phone1 := map[string][]string{
		"china_mobile": []string{
			"13771739111",
			"15216689566",
		},
		"china_telecom": []string{
			"18911112222",
			"18911113333",
		},
	}

	member1 := Person{
		Name:     "海阳之新",
		Sex:      "男",
		Birthday: "1985-10-24",
		Age:      35,
		IsMarry:  true,
		Phone:    phone1,
	}

	myFamily := MyFamily{
		Address: "湖南省新化县上渡办事处新城社区",
		Member: []Person{
			member1,
		},
	}

	//将结构体转换成json
	res, err := json.Marshal(myFamily)
	res2, _ := json.MarshalIndent(myFamily, "", "    ") //格式化输出json
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println(string(res))
		//结果：{"member":[{"name":"海阳之新","sex":"男","birthday":"1985-10-24","Age":35,"IsMarry":true,"Phone":{"china_mobile":["13771739111","15216689566"],"china_telecom":["18911112222","18911113333"]}}],"address":"湖南省新化县梅苑开发区"}
		fmt.Println(string(res2))
		//结果：会按json的格式换行缩进显示
	}

	var jsonStr = []byte(`{"member":[{"name":"海阳之新","sex":"男","birthday":"1985-10-24","Age":35,"IsMarry":true,"Phone":{"china_mobile":["13771739111","15216689566"],"china_telecom":["18911112222","18911113333"]}}],"address":"湖南省新化县梅苑开发区"}`)
	myFamily2 := MyFamily{}

	err = json.Unmarshal(jsonStr, &myFamily2) //将json字符串转为结构体，注意：第二个参数为指针
	if err != nil {
		fmt.Println("错误:", err)
	} else { //myFamily2
		fmt.Println(string(res))
	}
	var tmp []string
	tmp = myFamily2.Member[0].Phone["china_mobile"]
	fmt.Println(tmp)

}

//json转结构体
func JsonToStruct() {
	const jsonStream = `
		{"id": 1, "name": "海阳之新"}
		{"id": 2, "name": "欧阳"}
		{"id": 3, "name": "Simon"}
		{"id": 4, "name": "欧阳新民"}
		{"id": 5, "name": "海阳"}
	` //变量名忽略大小写，都可以解析正确，如NaMe
	//从一个输入流中读取并进行解码json的值
	//初始化一个Decoder对象
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	dec := json.NewDecoder(strings.NewReader(jsonStream)) //strings.Reader类型是为了高效读取字符串而存在的
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
			Name:     "海阳之新",
			Sex:      "男",
			Birthday: "1985-10-10",
			Age:      35,
			IsMarry:  true,
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
	bytes, _ := json.Marshal(myFamily)
	fmt.Printf("结构体转json：%s\n", bytes)

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

	varType := reflect.TypeOf(myFamily.Member[0].IsMarry)
	fmt.Println("通过反射来查询结构体成员的数据类型：", varType)
	//结果：通过反射来查询结构体成员的数据类型： bool

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

//json二次解析
func SendcondParse() {
	//如果不想指定Member变量为具体的类型，但仍然想保留interface{}类型，
	//且又希望该变量可以解析到struct Person对象中，这时候该怎么办呢？
	//可以将 Member 指定为 json.RawMessage
	type MyFamily2 struct {
		Address string          `json:"address"`
		Member  json.RawMessage `json:"member"`
	}

	jsonStr := "{\"address\":\"湖南新化梅苑开发区\",\"member\":[{\"name\":\"海阳之新\",\"sex\":\"男\",\"birthday\":\"1985-10-10\",\"age\":35,\"is_marry\":true,\"phone\":{\"china_mobile\":[\"13771739166\",\"15916689566\"],\"china_telecom\":[\"18971739166\",\"18916689566\"]}}]}"
	str := []byte(jsonStr)
	myFamily2 := MyFamily2{}
	json.Unmarshal(str, &myFamily2)
	fmt.Println("第一次解码结果为：\n", myFamily2)
	/* 结果：
		{湖南新化梅苑开发区 [91 123 34 110 97 109 101 34 58 34 230 181 183 233 152 179 228 185 139 230 150 176 34 44 34 115 101 120 34 58 3
	   4 231 148 183 34 44 34 97 103 101 34 58 51 53 44 34 105 115 95 109 97 114 114 121 34 58 116 114 117 101 44 34 112 104 111 110 101 34
		58 123 34 99 104 105 110 97 95 109 111 98 105 108 101 34 58 91 34 49 51 55 55 49 55 51 57 49 54 54 34 44 34 49 53 57 49 54 54 56 57
		53 54 54 34 93 44 34 99 104 105 110 97 95 116 101 108 101 99 111 109 34 58 91 34 49 56 57 55 49 55 51 57 49 54 54 34 44 34 49 56 57
		49 54 54 56 57 53 54 54 34 93 125 125 93]}
	*/
	member := []Person{}
	json.Unmarshal(myFamily2.Member, &member)
	fmt.Println("第二次解码结果为：\n", member)
	/* 结果：
	   [{海阳之新 男 1985-10-10 35 true map[china_mobile:[13771739166 15916689566] china_telecom:[18971739166 18916689566]]}]
	*/
}

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func JsonX() {
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
	Name     string              `json:"name"` //如果有json标签，struct转换成json时，key用标签里的名字，否则与结构体成员变量名字一致
	Sex      string              `json:"sex"`  //如果key为小写，则为私有，struct转换成json时不输出
	Birthday string              `json:"birthday"`
	Age      int                 `json:"age"` //如果两个成员的json标签都是一样的名字，接收数据时将导致两个值都为nil
	IsMarry  bool                `json:"is_marry"`
	Phone    map[string][]string `json:"phone"`
}

type MyFamily struct {
	Address string   `json:"address"` //json标签，如果这样声明`json:"name,omitempty"`，这个字段没有值时，不会填充
	Member  []Person `json:"member"`  //当数据结构不明确时，[]Person也可以用interface{}代替
}

type Message struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//参考：

/*

{
    "addess": "湖南省新化县梅苑开发区",
    "member": [
        {
            "name": "海阳之新",
            "sex": "男",
            "birthday":"1985-10-10"
            "age": 35,
            "is_marry": true,
            "phone": [
                "china_mobile": [
                     "13771739111",
                     "15216689566"
                 ],
                "china_telecom": [
                     "18911112222",
                     "18911113333"
                 ]
             ]
        },
        {
            "name": "海洋之心",
            "sex": "男",
            "birthday":"1988-11-09"
            "age": 32,
            "is_marry": false,
            "phone": [
                "china_mobile": [
                     "13771739188",
                     "15216689566"
                 ],
                "china_telecom": [
                     "18911112222",
                     "18911113333"
                 ]
             ]
        },
        {
            "name": "阳小斯",
            "sex": "男",
            "birthday":"1991-01-17"
            "age": 29,
            "is_marry": true,
            "phone": [
                "china_mobile": [
                     "13771739668",
                     "15216689123"
                 ],
                "china_telecom": [
                     "189111166666",
                     "189111177777",
                 ]
             ]
        },
    ]
}

*/
