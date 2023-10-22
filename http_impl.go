package apiprotocol

import (
	"encoding/json"

	"github.com/rs/xid"
)

/*
|srcId| string||是|否|发起方标识||fc659312b0e023f4107ecce69f43ad80|| |
|srcName| string||是|否|发起方名称||advertise|| |
|destId| string||是|否|目标方标识||fc65931207ecce69f43adkoioe80|| |
|destName| string||是|否|目标方名称||fsstorage|| |
|transportId | string||是|否|传输标识||154535|| |
|signature|string||是|是||签名,外网访问需开启签名|erefdsf154|
*/
type DefaultHttpProtocol struct {
	Head Head        `json:"_head"`
	Body interface{} `json:"_body"`
}

type Head struct {
	TransportId     string `json:"transportId"`
	Signature       string `json:"signature"`
	SignatureMethod string `json:"signatureMethod"`
	Type            string `json:"type"`
	DstId           string `json:"dstId"`
	DstName         string `json:"dstName"`
	SrcId           string `json:"srcId"`
	SrcName         string `json:"srcName"`
}

func NewDefaultRequestProtocol(c Config) (protocol DefaultHttpProtocol) {
	p := DefaultHttpProtocol{
		Head: Head{
			TransportId:     xid.New().String(),
			Signature:       "",
			SignatureMethod: c.SignatureMethod,
			DstId:           c.DstId,
			DstName:         c.DstName,
			SrcId:           c.SrcId,
			SrcName:         c.SrcName,
		},
	}
	return p
}

func (p DefaultHttpProtocol) Packet(input []byte) (out []byte, err error) {
	var body interface{}
	err = json.Unmarshal(input, &body)
	if err != nil {
		return nil, err
	}
	p.Body = body
	out, err = json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return out, err
}

func (p DefaultHttpProtocol) UnPack(input []byte) (out []byte, err error) {
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
