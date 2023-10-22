package apiprotocol

type EmptyProtocol struct {
}

func (p EmptyProtocol) Packet(input []byte) (out []byte, err error) {

	return input, nil
}
func (p EmptyProtocol) Unpack(input []byte) (out []byte, err error) {
	return input, nil
}
