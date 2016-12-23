package network

//"net"

const (
	SEND_CHAN_LEN uint32 = 10240
	READ_CHAN_LEN uint32 = 10240
)

type Config struct {
	Addr              string
	MaxReadMsgSize    uint32
	ReadMsgQueueSize  uint32
	ReadTimeOut       uint32
	MaxWriteMsgSize   uint32
	WriteMsgQueueSize uint32
	WriteTimeOut      uint32
}

func DecodeUint32(data []byte) uint32 {
	return (uint32(data[0]) << 24) | (uint32(data[1]) << 16) | (uint32(data[2]) << 8) | (uint32(data[3]))
}

func EncodeUint32(data uint32, b []byte) {
	b[3] = byte(data & 0xFF)
	b[2] = byte((data >> 8) & 0xFF)
	b[1] = byte((data >> 16) & 0xFF)
	b[0] = byte((data >> 24) & 0xFF)
}
