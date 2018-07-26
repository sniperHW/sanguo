package cs

import (
	"github.com/sniperHW/kendynet"
	"sanguo/codec/pb"
	"fmt"
)

const (
	sizeLen  = 2
	sizeFlag = 2
	sizeCmd  = 2 
)


type Encoder struct {
	namespace string
}

func NewEncoder(namespace string) *Encoder {
	return &Encoder{namespace:namespace}
}

func (this *Encoder) EnCode(o interface{}) (kendynet.Message,error) {
	switch o.(type) {
        case *Message:
			var pbbytes []byte
			var cmd uint32
			var err error
			msg := o.(*Message)
			if pbbytes,cmd,err = pb.Marshal(this.namespace,msg.GetData()); err != nil {
				return nil,err
			}

			totalLen := len(pbbytes) + sizeLen + sizeFlag + sizeCmd

			if uint64(totalLen) > maxPacketSize {
				return nil,fmt.Errorf("packet too large totalLen:%d",totalLen)
			}

			//len + flag + cmd + pbbytes
			buff := kendynet.NewByteBuffer(totalLen)
			//写payload大小
			buff.AppendUint16(uint16(totalLen-sizeLen))
			//写flag
			buff.AppendUint16(uint16(msg.GetSeriNo()))
			//写cmd
			buff.AppendUint16(uint16(cmd))
			//写数据
			buff.AppendBytes(pbbytes)
			return buff,nil
            break
        default:
            break
    }
    return nil,fmt.Errorf("invaild object type")
}