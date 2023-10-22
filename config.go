package apiprotocol

type Config struct {
	DstId            string           `mapstructure:"dstId" json:"dstId"  validate:"required"`
	DstName          string           `mapstructure:"dstName" json:"dstName"  validate:"required"`
	SrcId            string           `mapstructure:"srcId" json:"srcId"  validate:"required"`
	SrcName          string           `mapstructure:"srcName" json:"srcName"  validate:"required"`
	SignatureMethod  string           `mapstructure:"signatureMethod" json:"signatureMethod"  validate:"required"`
	SignatureMethods SignatureMethods `mapstructure:"signatureMethods" json:"signatureMethods"  validate:"required"`
}
