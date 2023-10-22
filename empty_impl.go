package apiprotocol

import "context"

type EmptyProtocol struct {
}

func (p EmptyProtocol) Packet(ctx context.Context, input []byte) (out []byte, err error) {

	return input, nil
}
func (p EmptyProtocol) Unpack(ctx context.Context, input []byte) (out []byte, err error) {
	return input, nil
}
