package strings

import (
	"encoding/json"
	"fmt"
)

func Json() {

	phone1 := map[string][]string {
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
		Addess: "湖南省新化县上渡办事处新城社区",
		Member: []Person{
			member1,
		},
	}

	//将结构体转换成json
	res, err := json.Marshal(myFamily)
	res2,_ := json.MarshalIndent(myFamily, "", "    ") //格式化输出json
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
	} else {//myFamily2
		fmt.Println(string(res))
	}
	var tmp []string
	tmp = myFamily2.Member[0].Phone["china_mobile"]
	fmt.Println(tmp)

}

type Person struct {
	Name     string                 `json:"name"` //json标签，如果这样声明`json:"name,omitempty"`，这个字段没有值时，不会填充
	Sex      string                 `json:"sex"`
	Birthday string                 `json:"birthday"`
	Age      int                    `json:age`
	IsMarry  bool                   `json:is_marry`
	Phone    map[string][]string `json:phone`
}

type MyFamily struct {
	Member []Person `json:"member"`
	Addess string   `json:"address"`
}

//参考：

/*

{
    "addess": "湖南省新化县梅苑开发区",
    "member": [
        {
            "name": "海阳之新",
            "sex": "男",
            "birthday":"1985-10-24"
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
