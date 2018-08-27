package main

import(
	"fmt"
	t "protobufDemo/example"

	"github.com/golang/protobuf/proto"
)

func main()  {
	    // 创建一个对象, 并填充字段, 可以使用proto中的类型函数来处理例如Int32(XXX)

		var hw t.Set
		var ld t.LD
	
		ld = t.LD {
			Ip: proto.Uint32(666),
		}
	
		hw = t.Set {
			Id: proto.Uint32(1),
			Name: proto.String("hello"),
			//LdList.append(ld),
		}
	
		// 对数据进行编码, 注意参数是message指针
		mData, err := proto.Marshal(&hw)
	
		if err != nil {
			fmt.Println("Error1: ", err)
			return
		}
	
		// 下面进行解码, 注意参数
		var umData t.Set
		err = proto.Unmarshal(mData, &umData)
	
		if err != nil {
			fmt.Println("Error2: ", err)
			return
		}
	
		// 输出结果
		fmt.Println(*umData.Id, "  ", *umData.Name, "  ", *ld.Ip)
}