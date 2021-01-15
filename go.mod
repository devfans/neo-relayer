module github.com/polynetwork/neo-relayer

go 1.14

require (
	github.com/boltdb/bolt v1.3.1
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/cmars/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/joeqian10/neo-gogogo v0.0.0-20201214075916-44b70d175579
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/ontio/ontology-crypto v1.0.9
	github.com/polynetwork/poly v0.0.0-20210108071928-86193b89e4e0
	github.com/polynetwork/poly-go-sdk v0.0.0-20200817120957-365691ad3493
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli v1.22.4
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	poly-bridge v0.0.0-00010101000000-000000000000
)

replace (
	poly-bridge => github.com/polynetwork/poly-bridge v0.0.0-20210112082403-a45d71989293
)
