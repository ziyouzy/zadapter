 package river_node 


//只是进行了一下简单的封装
func NewError(code int, uniqueId string, dataToString string, commit string)error{
	return NewEvent(code,uniqueId,dataToString,commit).ToError()
}


