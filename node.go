/** 此包的职能边界是将处理的各种类型(字节切片、NodeDoAbs)的功能模块(如心跳包，CRC校验器等)
 * 进行合理的匹配、部署、预加载(类似)
 * 装配于各适配器的map中 
 * 装配的内容是函数类型，而非功能对象的结构体或接口
   这么做目的在于节省系统资源，实现了类似预加载的效果
 
 * 当上层使用此包时，上层结构类内可以用一个接口切片的字段承载从本包map拿到的对象进行真正的初始化
 * 同时因为其载体为切片而非map所以也可以控制切片内各个适配器的先后执行顺序
  
 * 而各个适配器的具体实现则需要独立的包来完成，如package heartbeating
   独立的包需要设计好实现了某个合适自己子门类的适配器接口
   并在进行初始化时先存入map实现预编译的效果，最终实现上层的调用
 */

 /*所有river-node所产生的event和error，在主系统中的处理方式需遵循如下原则：*/
 //1.Errors和Events均禁止被发送给客户端
 //2.整体系统可基于river-node的错误采取后续策略(如p.warpError_Panich；BAITSFILTER_LENAUTHFAIL->river success)
 //3.整体系统不可基于river-node的event采取后续策略
 //4.所有river-node的event唯一的后续策略只有记录日志
 //这么规定另一个好处是可以让eebox中的eventbox变得简单


package river_node



var RegisteredNodes = make(map[string]NodeAbstractFunc)
func Register(Name string, F NodeAbstractFunc) {

	if RegisteredNodes[Name] != nil {
		panic("river-node: " + Name + " already registered!")
	}
	
	if F == nil {
		panic("river-node: " + Name + " is nil!")
	}

	RegisteredNodes[Name] = F
} 


type NodeAbstractFunc func() NodeAbstract


type NodeAbstract interface {
	Name() string
	Run()
	Construct(config Config) error
}




/** 这里使用了继承
 * 这样一来就可以把RiverNode作为一个切片储存了，同时他还能具有map的键值对特性
 * 这其实是一种设计模式
 */
type Node struct{
	Name  string
	NodeAbstract
}


