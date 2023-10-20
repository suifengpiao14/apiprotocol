package apiprotocol

type ApiProtocol interface {
	Packet(input []byte) (out []byte, err error)
	Unpack(input []byte) (out []byte, err error)
}
