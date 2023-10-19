package apiprotocol

type ApiProtocol interface {
	Encode(input string) (out string, err error)
	Decode(input string) (out string, err error)
}
