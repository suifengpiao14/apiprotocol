package apiprotocol

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
	"github.com/spf13/cast"
)

/*
|clientId| string||是|否|发起方标识||fc659312b0e023f4107ecce69f43ad80|| |
|clientName| string||是|否|发起方名称||advertise|| |
|requestId | string||是|否|传输标识||154535|| |
|signature|string||是|是||签名,外网访问需开启签名|erefdsf154|
*/
type DefaultHttpProtocol struct {
	Config Config      `json:"-"`
	Head   Head        `json:"_head"`
	Body   interface{} `json:"_body"`
}

type Head struct {
	RequestId       string `json:"requestId"`
	Signature       string `json:"signature"`
	SignatureMethod string `json:"signatureMethod"`
	Type            string `json:"type"`
	ClientId        string `json:"clientId"`
	ClientName      string `json:"clientName"`
}

type ContextName string

const (
	Context_Name_RequestId ContextName = "requestId"
)

// SetRequestID 设置请求ID
func SetRequestID(ctx context.Context, requestId string) (newCtx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	newCtx = context.WithValue(ctx, Context_Name_RequestId, requestId)

	return newCtx
}

// GetRequestID 获取请求ID
func GetRequestID(ctx context.Context) (requestId string) {
	if ctx == nil {
		ctx = context.Background()
	}
	v := ctx.Value(Context_Name_RequestId)
	if v == nil {
		v = xid.New().String()
	}
	requestId = cast.ToString(v)
	return requestId
}

func NewDefaultRequestProtocol(c Config) (protocol DefaultHttpProtocol) {
	p := DefaultHttpProtocol{
		Config: c,
		Head: Head{
			SignatureMethod: c.SignatureMethod,
			ClientId:        c.ClientId,
			ClientName:      c.ClientName,
		},
	}
	return p
}

func (p DefaultHttpProtocol) Packet(ctx context.Context, input []byte) (out []byte, err error) {
	var body interface{}
	err = json.Unmarshal(input, &body)
	if err != nil {
		return nil, err
	}
	signatureFn, err := GetSignature(p.Head.SignatureMethod)
	if err != nil {
		return nil, err
	}
	signature, err := signatureFn.Signature(p.Config, input)
	if err != nil {
		return nil, err
	}
	p.Head.RequestId = GetRequestID(ctx)
	p.Head.Signature = signature
	p.Body = body
	out, err = json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (p DefaultHttpProtocol) Unpack(ctx context.Context, input []byte) (out []byte, err error) {
	pro := DefaultHttpProtocol{}
	err = json.Unmarshal(input, &pro)
	if err != nil {
		return nil, err
	}
	out, err = json.Marshal(pro.Body)
	if err != nil {
		return nil, err
	}
	return out, err
}
