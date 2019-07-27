package main

import (
	"fmt"
	"reflect"
	"time"
)

// twitter 雪花算法
// 把时间戳,工作机器ID, 序列号组合成一个 64位 int
// 第一位置零, [2,42]这41位存放时间戳,[43,52]这10位存放机器id,[53,64]最后12位存放序列号


var (
	machineID    int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	sn            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
)

func init() {
	lastTimeStamp = time.Now().UnixNano() / 1000000
}

func SetMachineId(mid int64) {
	// 把机器 id 左移 12 位,让出 12 位空间给序列号使用
	machineID = mid << 12
}

func GetSnowflakeId() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000000
	// 同一毫秒
	if curTimeStamp == lastTimeStamp {
		sn++
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if sn > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano() / 1000000
			lastTimeStamp = curTimeStamp
			sn = 0
		}

		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}
	if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		id := rightBinValue | machineID | sn
		return id
	}
	if curTimeStamp < lastTimeStamp {
		return 0
	}
	return 0
}

func main() {
	//var ids = []int64{}
	var ids = make([]int64, 0)

	//设置一个机器标识，如IP编码,防止分布式机器生成重复码
	SetMachineId(192168100101)

	fmt.Println("start", time.Now().Format("13:04:05"))
	for i := 0; i < 10000000; i++ {
		id := GetSnowflakeId()
		ids = append(ids, id)
	}
	fmt.Println("end  ", time.Now().Format("13:04:05"))

	result := Duplicate(ids)
	fmt.Println("去重后数量：", len(result))
	fmt.Println(result[10], result[11], result[12], result[13], result[14])
	fmt.Println(result[9990], result[9991], result[9992], result[9993], result[9994])
}

//去重
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}