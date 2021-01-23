package binchunk

type binaryChunk struct{

	header		    // 头部
	sizeUpvalues byte   // 主函数 upvalue 数量
	mainFunc *Prototype // 主函数原型


}


type header struct{
	signature       [4]byte
	version		byte
	format		byte
	luacData	[6]byte
	cintSize	byte
	sizetSzie	byte
	insturctionSize byte
	luaIntegerSize	byte
	luaNumberSize	byte
	luacInt		int64
	luacNum		float64
}


