package apiprotocol

type Config struct {
	ClientId         string           `mapstructure:"clientId" json:"clientId"  validate:"required"`
	ClientName       string           `mapstructure:"clientName" json:"clientName"  validate:"required"`
	SignatureMethod  string           `mapstructure:"signatureMethod" json:"signatureMethod"  validate:"required"`
	SignatureMethods SignatureMethods `mapstructure:"signatureMethods" json:"signatureMethods"  validate:"required"`
}
