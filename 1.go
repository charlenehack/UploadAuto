import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	json := `{
	  "name": {"first": "Tom", "last": "Anderson"},
	  "age": 37,
	  "children": ["Sara","Alex","Jack"],
	  "fav.movie": "Deer Hunter",
	  "friends": [
		{"first": "Dale", "last": "Murphy", "age": 44},
		{"first": "Roger", "last": "Craig", "age": 68},
		{"first": "Jane", "last": "Murphy", "age": 47}
	  ]
	}`

	name := gjson.Get(json, "name.first")  // name.String()返回 Tom
	age := gjson.Get(json, "age")  // age返回 37
	children := gjson.Get(json, "children")  // children.String()返回 ["Sara","Alex","Jack"]
	childrenNum := gjson.Get(json, "children.#")  // childrenNum返回个数3
	children2 := gjson.Get(json, "children.1")  // children2.String()返回第1个元素
	children3 := gjson.Get(json, "child*.2")  // 可使用通配符匹配，*表示1个或多个，?表示1个
	friends := gjson.Get(json, "friends.#.last")  // 返回 ["Murphy","Craig","Murphy"]
	friend0 := gjson.Get(json, "friends.0.last")   // 返回 Murphy
	friendAge := gjson.Get(json, `friends.#[last=="Craig"].age`)  // 返回 68
	friendFirst := gjson.Get(json, `friends.#[last=="Murphy"]#.first`)  // 返回  ["Dale","Jane"]
	friend45 := gjson.Get(json, `friends.#[age>45]#.last`)  // 返回 ["Craig","Murphy"] 年龄大于45的
	friendD := gjson.Get(json, `friends.#[first%"D*"].last`) // 返回 Murphy, %表示模糊匹配
	fmt.Println(name.String(), age.Int(), children.String(), childrenNum.Int(), children2.String(), children3.String(), friends.String(), friend0.String(), friendAge.Int(), friendFirst, friend45, friendD)

	result := gjson.Get(json, "friends.#.last")
	for _, name := range result.Array() {    // 读取嵌套数组
		fmt.Println(name.String())
	}
	Roger := gjson.Get(json, "friends.1").Get("last")  // 嵌套查询
	fmt.Println(Roger)

	if gjson.Get(json, "friends").Exists() {  // 判断key是否存在
		fmt.Println("存在key")
	}

	if gjson.Valid(json) {   // 判断json格式是否正确
		fmt.Println("是一个有效的json")
	}

	results := gjson.GetMany(json, "name.last", "age", "friends.1.first")  // 一次获取多值
	fmt.Println(results)

	m, ok := gjson.Parse(json).Value().(map[string]interface{})   // 反序列化到map
	if !ok {
		// not a map
	}
	fmt.Println(m)
}
