module github.com/polynetwork/neo-relayer

go 1.14

require (
	github.com/boltdb/bolt v1.3.1
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/joeqian10/neo-gogogo v0.0.0-20210120033000-0b38545f3328
	github.com/ontio/ontology-crypto v1.0.9
	github.com/polynetwork/poly v1.3.1
	github.com/polynetwork/poly-go-sdk v0.0.0-20210114035303-84e1615f4ad4
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.4
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	poly-bridge v0.0.1
)

replace poly-bridge => github.com/polynetwork/poly-bridge v1.0.1-0.20210924034233-ace09e709658
