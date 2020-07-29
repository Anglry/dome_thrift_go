package main

import (
	"fmt"
	"fuwu/service"
	"fuwu/thrift/demo_record"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

const (
	NetworkAddr = "127.0.0.1:9090"//监听地址&端口
)

func main(){
	/**
	数据传输格式（protocol）
	定义的了传输内容，对Thrift Type的打包解包，包括：

	TBinaryProtocol，二进制格式，TBinaryProtocolAccelerated则是依赖于thrift_protocol扩展的快速打包解包。
	TCompactProtocol，压缩格式
	TJSONProtocol，JSON格式
	TMultiplexedProtocol，利用前三种数据格式与支持多路复用协议的服务端（同时提供多个服务，TMultiplexedProcessor）交互
	数据传输方式（transport）
	定义了如何发送（write）和接收（read）数据，包括：

	TBufferedTransport，缓存传输，写入数据并不立即开始传输，直到刷新缓存。
	TSocket，使用socket传输
	TFramedTransport，采用分块方式进行传输，具体传输实现依赖其他传输方式，比如TSocket
	TCurlClient，使用curl与服务端交互
	THttpClient，采用stream方式与HTTP服务端交互
	TMemoryBuffer，使用内存方式交换数据
	TPhpStream，使用PHP标准输入输出流进行传输
	TNullTransport，关闭数据传输
	TSocketPool在TSocket基础支持多个服务端管理（需要APC支持），自动剔除无效的服务器
	 */

	//传输器
	//Transport层提供了一个简单的网络读写抽象层，这使得thrift底层的transport从系统其他的部分解耦。Thrift使用ServerTransport接口接受或者创建原始transport对象。
	// ServerTransport用在Server端，为到来的连接创建Transport对象。
	//注意transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	//这可能有不同的选项，大部分参考代码中给的都是transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	//客户端连接时候一定要与此对应。
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	//传输协议
	//Protocol抽象层定义了一种怎样将内存中数据结构映射成可传输格式的机制。Protocol定义了datatype怎样使用底层的Transport对自己进行编解码
	protocolFactory  := thrift.NewTBinaryProtocolFactoryDefault()
	//处理器

	prosser := thrift.NewTMultiplexedProcessor()

	//new和其他语言中的同名函数一样，new(t)分配了零值填充的类型为T内存空间，并且返回其地址，即一个*t类型的值。返回的永远是类型的指针，指向分配类型的内存地址
	//new(service.DemoService)  等价于 handler = &service.DemoService{}
	dome := demo_record.NewDemoServiceProcessor(new(service.DemoService))
	prosser.RegisterProcessor("dome",dome)

	//服务器
	//获取地址
	serverTransport,err := thrift.NewTServerSocket(NetworkAddr)
	//服务设置
	server := thrift.NewTSimpleServer4(prosser,serverTransport,transportFactory,protocolFactory)
	if err!= nil{
		fmt.Println("Error!",err)
		os.Exit(1)
	}
	//服务启动
	server.Serve()
	// 退出时停止服务器
	defer server.Stop()
}
