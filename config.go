package apiprotocol

type Config struct {
	SrcId            string           `mapstructure:"srcId" json:"srcId"  validate:"required"`
	SrcName          string           `mapstructure:"srcName" json:"srcName"  validate:"required"`
	SignatureMethod  string           `mapstructure:"signatureMethod" json:"signatureMethod"  validate:"required"`
	SignatureMethods SignatureMethods `mapstructure:"signatureMethods" json:"signatureMethods"  validate:"required"`
}
