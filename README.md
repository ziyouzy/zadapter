# zadapter
~~会与zconn、heartbreating联合使用~~已经实现了zadapter与heartbreating的融合，从而学习并实现一种基于tcp长连接的设计思路  

之前为了让自己思路清晰从而把conn独立了出来，形成了zconn  
现在也需要把adapter独立出来，这次思路应该可以稳定下来了吧  

此包的重要意义在于实现了一种“适用于当前需求的框架模式”  
**框架模式**是个十分重要的概念，可以参考“MVC框架模式”价值与存在意义：  
https://baike.baidu.com/item/MVC%E6%A1%86%E6%9E%B6/9241230?fromtitle=mvc&fromid=85990&fr=aladdin  
文章中有如下描述：  

    框架和设计模式的区别
    MVC是一种框架模式。   
    框架、设计模式这两个概念总容易被混淆，其实它们之间还是有区别的。
    框架通常是代码重用，而设计模式是设计重用，架构则介于两者之间，部分代码重用，部分设计重用，有时分析也可重用。
    在软件生产中有三种级别的重用：内部重用，即在同一应用中能公共使用的抽象块;
    代码重用，即将通用模块组合成库或工具集，以便在多个应用和领域都能使用；
    应用框架的重用，即为专用领域提供通用的或现成的基础结构，以获得最高级别的重用性。
    框架与设计模式虽然相似，但却有着根本的不同。设计模式是对在某种环境中反复出现的问题以及解决该问题的方案的描述，它比框架更具象；
    框架可以用代码表示，也能直接执行或复用，而对模式而言只有实例才能用代码表示;
    设计模式是比框架更小的元素，一个框架中往往含有一个或多个设计模式，框架总是针对某一特定应用领域，但同一模式却可适用于各种应用。
    可以说，框架是软件，而设计模式是软件的知识。
    
    框架模式有哪些？
    MVC、MTV、MVP、CBD、ORM等等；
    
    框架有哪些？
    C++语言的QT、MFC、gtk，Java语言的SSH 、SSI，php语言的 smarty(MVC模式)，python语言的django(MTV模式)等等
    
    设计模式有哪些？
    工厂模式、适配器模式、策略模式等等
    
    简而言之：框架是大智慧，用来对软件设计进行分工；设计模式是小技巧，对具体问题提出解决方案，以提高代码复用率，降低耦合度。  
**同时也要画个重点：框架总是针对某一特定应用领域**  
正因如此mvc肯定至少是一种框架，数据流动也是一种框架  
比如说web应用，聊天室，源妹做的小游戏，都应该属于不同的应用领域  
至少这三者，由于应用领域不同，分别实现他们的需求所选择的框架应该也是不同的  

**因此，由于此包的目的在于解决数据流动领域的问题，所以此包的本质是一种框架思路**

***
**2021年4月4日18点02分:**  
**心跳包的功能并不是发送一个数据给另一端，而是接收另一端的数据，从而判断另一端的在/离线状态**  
**在目前的项目中心跳包的使用方式更倾向于一个发送/响应消息循环后的“清点”工作**  
**比如sender_tousrio808定1秒发送询码给usrio808设备，usrio808会立刻发送回码**  
**要注意接收方不再是sender_tousrio808对象，而是换成了receiver_fromusrio808，接收后就完成了一个消息循环**  
**于是receiver_fromusrio808对象就可以承担起循环后的“清点”工作，方法就是给他内部放置一个heartbeating**  
**2021年3月12日11点53分:**  
**错误的数据必须转化为error，且两者不能并存**  
**正确的数据可以产生event，且两者可以并存**  

**2021年2月26日14点14分:**  
**无论是代码还是思路，都彻底完成了！！！**

**2021年2月23日11点35分:**  
设计各个rnode的析构方式，这项工作将留给设计package zconn再去完成  

**2021年2月23日11点25分:**  
在真正的实战中，应该尽量避免singal打印内容，event应该奉行少说多做原则  
而error则会担起“汇报的责任”  
同时那个模块出现了error，需要在哪个层解决，而不是经由singal调度解决问题  
这确实是最近总结的很重要的原则之一  

**2021年2月13日18点23分:**  
这个包貌似是做好了  

**2021年2月13日09点18分:**  
关于go logger的同步与异步
异步就是线程不会阻塞在这儿，直到确认发出了消息为止。  
而是执行了发出的动作之后，就执行下面的代码了  
总之比想象起来更容易理解也是个好事  

**2021年2月5日10点53分:**  
感觉好好，更加深入的理解了单项调用链模式，同时也更加深入的理解了golang本地宝的相互调用机制  
**crc与heartbreating这两个包有些类似于使用范例**  
**虽然他们的文件目录存在于zadapter下方，但是使用的时候他们却是zadapter的最上层**  
**当你自己某个项目需要用到zadapter包，可以看看这两个，就可以设计自定义的适配器了，而这些自定义的适配器只需要import zadapter就可以实现**  

**2021年2月2日17点44分:**  
感觉有点悲催了，各种逻辑上的推到重来，代码重构，有些源文件已经不需要了，但是还有一些备注或许以后会用到所以先梳理下这些备注的思路：  
1.早期思路ToDo方法会有返回值和参数表，但是目前的思路是没有返回值也没有参数表，一切参数与返回值都会基于config结构体实现，这也是准备模仿Qt信号与槽的特性之一  
2.不会再有bytesabs.go与nodedoabs.go这两个源文件以及其所代表的bytes与nodedo之间的属性区别(因为两者的ToDo方法都会基于上面1所示规则)，于是就可以同一用同一个抽象接口进行承载了  
3.同时，准备采用信号与槽的相互通信方式，似乎拦截器，过滤器，触发器等都可以融为一体了，于是也不会再有Interceptor.go、trigger.go、vaildate.go(触发器，拦截器，过滤器)之类的区别，而是将业务逻辑垂直至底层，直接去实现heartbreating.go、crc.go之类的实体功能对象，他们都将实现adapter.go所描述的接口（adapter.go/AdapterAbstract）  



