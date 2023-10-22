package apiprotocol

import "context"

type ApiProtocol interface {
	Packet(ctx context.Context, input []byte) (out []byte, err error)
	Unpack(ctx context.Context, input []byte) (out []byte, err error)
}
