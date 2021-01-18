module github.com/polynetwork/neo-relayer

go 1.14

require (
	github.com/boltdb/bolt v1.3.1
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/joeqian10/neo-gogogo v0.0.0-20210118094521-237d985a02d5
	github.com/ontio/ontology-crypto v1.0.9
	github.com/polynetwork/poly v0.0.0-20210108071928-86193b89e4e0
	github.com/polynetwork/poly-go-sdk v0.0.0-20200817120957-365691ad3493
	github.com/polynetwork/poly-io-test v0.0.0-20200819093740-8cf514b07750 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli v1.22.4
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	poly-bridge v0.0.0-00010101000000-000000000000
)

replace poly-bridge => github.com/polynetwork/poly-bridge v0.0.0-20210112082403-a45d71989293
