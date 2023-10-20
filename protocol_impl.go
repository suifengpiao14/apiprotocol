package apiprotocol

/*
|srcId| string||是|否|发起方标识||fc659312b0e023f4107ecce69f43ad80|| |
|srcName| string||是|否|发起方名称||advertise|| |
|destId| string||是|否|目标方标识||fc65931207ecce69f43adkoioe80|| |
|destName| string||是|否|目标方名称||fsstorage|| |
|transportId | string||是|否|传输标识||154535|| |
|signature|string||是|是||签名,外网访问需开启签名|erefdsf154|
*/
type DefaultApiProtocol struct {
	DestID   string `json:"destId"`
	DestName string `json:"destName"`
	SrcId    string `json:"srcId"`
	SrcName  string `json:"srcName"`
}

func (p DefaultApiProtocol) Packet(input []byte) (out []byte, err error) {

	return out, err
}

func (p DefaultApiProtocol) UnPack(input []byte) (out []byte, err error) {

	return out, err
}
