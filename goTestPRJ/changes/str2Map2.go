package changes

import (
	"fmt"
	"encoding/json"
	"github.com/liamylian/values"
)


func main(){

	// 从map创建
	m := map[string]string{"foo": "bar"}
	vs := values.FromMap(m)

	// 转换成map
	vs.ToMap()

	// json序列号与反序列化
	bytes, _ :=json.Marshal(vs)
	json.Unmarshal([]byte(`{"foo":"bar","int":1,"bool":true}`), &vs)

	// 获取值
	value, exists := vs.Get("foo")

	// 设置值
	vs.Set("earth", "moon")

	// 删除值
	vs.Delete("foo")

	// 批量设置值
	m = map[string]string{"bar": "foo", "moon":"earth"}
	vs.Sets(m)

	// 遍历
	vs.Range(func(key string, value string) bool {
	fmt.Println("%s: %s", key, value)
	return true
	})
}