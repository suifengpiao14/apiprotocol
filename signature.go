package apiprotocol

import "github.com/pkg/errors"

type SignatureInterface interface {
	GetName() (name string)
	Signature(c Config, payload []byte) (signature string, err error)
}

var (
	ERROR_NOT_FOUND_SIGNATURE_BY_NAME = errors.New("not found signature by name")
)

type SignatureMethods []string

func (m SignatureMethods) Find(method string) (err error) {
	for _, me := range m {
		if me == method {
			return nil
		}
	}
	return errors.WithMessagef(ERROR_NOT_FOUND_SIGNATURE_BY_NAME, "method:%s", method)
}

const (
	Signature_name_none = "none"
)

type SignatureNone struct{}

func (s SignatureNone) GetName() (name string) {
	return Signature_name_none
}

func (s SignatureNone) Signature(c Config, payload []byte) (signature string, err error) {
	return "", nil
}

var (
	signaturePool = []SignatureInterface{
		SignatureNone{}, // 默认注册none签名
	}
	ERROR_NOT_FOUND_SIGNATURE = errors.New("not found signature")
)

func RegisterSignature(s SignatureInterface) {
	signaturePool = append(signaturePool, s)
}

func GetSignature(name string) (s SignatureInterface, err error) {
	for _, signature := range signaturePool {
		if signature.GetName() == name {
			return signature, nil
		}
	}
	return nil, ERROR_NOT_FOUND_SIGNATURE
}
