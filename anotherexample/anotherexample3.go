/** 本包的目的在于演示如何设计一个非字节切片数据类型的适配器
 * 此包作为范例，其目标数据类型是一个自定义结构体
 * 而对于实际场景中实现功能时，各种类型都可以用类似的方式去设计
 * 最终完成最适合的某种适配器
 */

package anotherexample

import (
	/** 引入zadaptr包与another包都遵循了单向调用链原则
	 * 虽然文件夹部署上zadapter文件夹在最外层
	 * anotherexample文件夹在中间层
	 * another文件夹在最内层
	 * 但是从包的相互引用逻辑上最末是anotherexample包
	 */
	"zadapter"
	"zadapter/anotherexample/another"
	"zadapter/define"

	"math/rand"
	"time"
	"reflect"
	"errors"
	"fmt"
	//"bytes"
	//"encoding/binary"
)


const ADAPTER_NAME3 = "anotherexample"


/*范例Config*/
type AnotherExmaple3Config struct{

	UniqueId string	/*其所属上层数据通道(如Conn)的唯一识别标识*/

	Mode int

	SignalChan chan int /*发送给主进程的信号队列，就像Qt的信号与槽*/

	RawinChan chan another.AnotherFunc

	NewoutChan chan *another.AnotherStruct
}

func (p *AnotherExmaple3Config)Name()string{
	return ADAPTER_NAME3
}

type AnotherExmaple3 struct{
	rand int
	config *AnotherExmaple3Config
}

func (p *AnotherExmaple3)Name()string{
	return ADAPTER_NAME3
}


func (p *AnotherExmaple3)Init(anotherExmaple3ConfigAbs zadapter.Config) error{
	if anotherExmaple3ConfigAbs.Name() != ADAPTER_NAME3 {
		return errors.New("anotherexmaple3 adapter init error, config must AnotherExmaple3Config")
	}


	value := reflect.ValueOf(anotherExmaple3ConfigAbs)
	config := value.Interface().(*AnotherExmaple3Config)


	if config.signalChan == nil||config.rawChan == nil||config.newChan == nil{
		return errors.New("anotherexmaple3 adapter init error, slotChan or signalChan "+
		                  "or newChan is nil")
	}

	p.config = config

	rand.Seed(time.Now().UnixNano())
	p.rand = rand.Int()


	switch p.config.mode{
	case TEST1:
		fmt.Println("type is anotherexample3, mode is TEST1")
	case TEST2:
		fmt.Println("type is anotherexample3, mode is TEST2")
	case TEST3:
		fmt.Println("type is anotherexample3, mode is TEST3")
	default:
		return errors.New("anotherexmaple3 adapter init error, unknown mode")
	}
	
	return nil
}

func (p *AnotherExmaple3)Run(){
	switch p.config.mode{
	case TEST3:
		go func(){
			for rawf := range p.config.rawChan{

				/*仅仅作为示范：*/
				p.config.signalChan<-define.ANOTHEREXAMPLE_TEST3

				rawf(append([]byte("测试3(TEST3)"), 0x12,0x33,0xff))

				p.config.newChan<-&another.AnotherStruct{
						TimeStamp : time.Now().UnixNano(),
						Tip : "仅仅作为示范(AnotherExmaple3)",
						TestSl: []float64{1.2,1.33,1,444},
						//rawSl: sl,//本来每次循环结束sl就会销毁，但是这样写sl就会持久化了
						RawSl: []byte{0x12,0x33,0xff},//这样是进行深拷贝，sl不会在有外部引用他了，sl每次循环结束都被销毁
					}
			}
		}()
	default:
		p.config.signalChan<-define.ANOTHEREXAMPLE_ERR
	}
}



func NewAnotherExmaple3() zadapter.AdapterAbstract {
	return &AnotherExmaple3{}
}


func init() {
	zadapter.Register(ADAPTER_NAME3, NewAnotherExmaple3)
}

