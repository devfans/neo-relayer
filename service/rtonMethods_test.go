package service

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet/keys"
	"github.com/polynetwork/neo-relayer/log"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestSyncService_GetCurrentNeoChainSyncHeight(t *testing.T) {
	height := 64
	//s := string(height) // ascii code
	s1 := strconv.Itoa(height)

	assert.Equal(t, "64", s1)
}

func Test_ConsensusPayload(t *testing.T) {
	//bs := []byte{}
	//ConsensusPayload := helper.HexToBytes("7b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a342c2263223a312c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303237613165326539616130626462366538333435633435653962643666333139656636323439396638383638366361663563616239613034616631663131343266227d2c7b22696e646578223a322c226964223a2231323035303234356261346534623033613365396665643737616533356163303065363930613564653464393061343938343666613331666535663135363837373734613338227d2c7b22696e646578223a332c226964223a2231323035303335363435396366653165336636663032626535306637396230306562613934323531323138663162356237653030313231646264653032616139663066363933227d2c7b22696e646578223a342c226964223a2231323035303234346266343036396461333332613138383566353262636161623563643231626336646130643834626333373433613232353666663464633335663566343738227d5d2c22706f735f7461626c65223a5b322c322c332c332c342c332c312c322c332c332c312c342c332c342c312c312c312c312c342c332c322c312c332c322c312c322c332c322c312c322c322c342c332c312c342c342c312c312c322c332c322c332c322c342c322c322c342c322c332c312c342c342c342c342c332c312c342c342c312c335d2c226d61785f626c6f636b5f6368616e67655f76696577223a33307d7d")
	//blkInfo := &vconfig.VbftBlockInfo{}
	//_ = json.Unmarshal(ConsensusPayload, blkInfo) // already checked before
	//if blkInfo.NewChainConfig != nil {
	//	for _, peer := range blkInfo.NewChainConfig.Peers {
	//		keyBytes, _ := hex.DecodeString(peer.ID)
	//		key, _ := keypair.DeserializePublicKey(keyBytes) // compressed
	//		uncompressed := getRelayUncompressedKey(key)
	//		bs = append(bs, uncompressed...)
	//	}
	//}
	//
	//log.Infof("public keys hex string: ", helper.BytesToHex(bs))
}

func Test_111(t *testing.T) {
	client := rpc.NewClient("http://47.88.50.171:21332")
	for i := 11570; i > 0; i-- { // 8553
		r := client.GetBlockByIndex(uint32(i))
		rpcBlock := r.Result
		if len(rpcBlock.Tx) > 1 {
			for j := 0; j < len(rpcBlock.Tx); j++ {
				if rpcBlock.Tx[j].Type == "InvocationTransaction" {
					itx := rpcBlock.Tx[j]
					appLog := client.GetApplicationLog(itx.Txid)
					executions := appLog.Result.Executions
					//notifications := executions[0].Notifications
					gas := executions[0].GasConsumed
					f, _ := helper.Fixed8FromString(gas)
					if f.GreaterThan(helper.Fixed8FromInt64(10)) {
						log.Info(itx.Txid)
						//l := len(notifications)
						//notif := notifications[l-2]
						//value := notif.State.Value[0].Value
						//if value == "3" {
						//	log.Info(itx.Txid)
						//}
					}
				}
			}
		}
	}
}

func Test_BtcRecoverPubKey(t *testing.T) {
	sigList := []string{"1bd3bf4e6d30818a23be0f539e0dd9c0814bbe08b9fc7afaa9ecef9923432af73243973e7b16dd8dcb297cc4a50a3c9123d8384016462eebadcee299b3cc40fe94",
						"1ccfadad40ac16fee22b9d4e858ff3a06dc8fc09014f2de1255c34d21ebb43163f0053ae213ac2f6e33c5e4ed2a1c034317fb5519603844a522f8028d52c71cb89",
						"1b7580f7c1cc02c9dacafbe13fbbff170c5ac2831a5e24a858eeba4860656e78274d5977aead28214d00b8746edba25d618a9a1c1bb477cc2381c918ed94fa27b3",
						"1c962cb527cd730b7d00024d1f2c84fae8e755d5416226b7c6b839911b6b6c57321d2021135ff1916208052104aaea38e55af8b254f0f5f3551cf2314a9cddd21b"}
	//msgHash := crypto.Keccak256(helper.HexToBytes("000000000000000000000000707e32185504173b56cd4e155e6cc7e64ef7a1fba1c2b663391e9f91b2eed3f6e619776c3163176d5b0fc00ac7356c0729d26c06adf9d8b909153bb43cca6e6b000000000000000000000000000000000000000000000000000000000000000096d8ce3dd0a00dc7621cb72bd257f9546702ea383225bdb8ef1a0dcd318961188f6e945e3c000000e41feeb93394be87fd9e037b226c6561646572223a332c227672665f76616c7565223a22424e345042485a665a38305758637750536d32733865486e6a37644275625662574d54555a4674446a672f746e743770434b4c663842392f4c6b546f6e6f2b5a436835504d6a304d6f7967744b7a5365364345582b49673d222c227672665f70726f6f66223a22694a62344d31456e77617553743366557841566f62314b3642767831354e6c585048704878714f53456374595179464637556c7a6c7a734e71546a76534b434d6e417541394a445339514b6b68634e476a48693346773d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a36302c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a332c226e223a342c2263223a312c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a332c226964223a2231323035303335363435396366653165336636663032626535306637396230306562613934323531323138663162356237653030313231646264653032616139663066363933227d2c7b22696e646578223a312c226964223a2231323035303237613165326539616130626462366538333435633435653962643666333139656636323439396638383638366361663563616239613034616631663131343266227d2c7b22696e646578223a322c226964223a2231323035303234356261346534623033613365396665643737616533356163303065363930613564653464393061343938343666613331666535663135363837373734613338227d2c7b22696e646578223a342c226964223a2231323035303234346266343036396461333332613138383566353262636161623563643231626336646130643834626333373433613232353666663464633335663566343738227d5d2c22706f735f7461626c65223a5b342c322c322c312c332c342c332c312c312c322c322c322c342c332c322c312c322c342c332c332c312c332c342c342c312c312c332c332c332c332c332c312c342c332c322c312c342c322c312c332c322c312c332c342c322c332c312c322c342c342c342c312c342c342c322c322c342c312c322c315d2c226d61785f626c6f636b5f6368616e67655f76696577223a33307d7d3dc4ba9418037af1730c130f068f4951653070ad"))
	hash := helper.HexToBytes("13f8f127fc829f335f12ea3befbc514f70149ae3d04cd8d003198b3b961971c8")
	for _, sigStr := range sigList {
		sig := helper.HexToBytes(sigStr)
		pubKey, isCompressed, err := btcec.RecoverCompact(btcec.S256(), sig, hash) // byte(4) = 0100
		assert.Nil(t, err)
		log.Info(helper.BytesToHex(pubKey.SerializeCompressed()) + " is compressed: " + strconv.FormatBool(isCompressed))
	}
}

func Test_999(t *testing.T)  {
	//sb := sc.NewScriptBuilder()
	//scriptHash := helper.HexToBytes("7f25d672e8626d2beaa26f2cb40da6b91f40a382") // hex string to little endian byte[]
	//
	//cp1 := sc.ContractParameter{
	//	Type:  sc.ByteArray,
	//	Value: helper.HexToBytes("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000453c810e2d58aeb9aabc22723666785fa200c1d9fea5a5006d9e506df0911d7e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008e305f000000001dac2b7c00000000fdb2037b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a342c2263223a312c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303330396336343735636530373537376162373261316639366332363365353033306362353361383433623030636131323338613039336439646362313833653266227d2c7b22696e646578223a322c226964223a2231323035303332626564353565386334643963626335303635376666353930396565353164633339346139326161643931316333366261636538336334643633353430373934227d2c7b22696e646578223a332c226964223a2231323035303265363861366535346264666130616634376264313834363566343335326635313531646337323963363161373339393930396631636431633664383136633032227d2c7b22696e646578223a342c226964223a2231323035303232396530643163356232616538333839333061653161643836316464643364303734356431633766313432343932636162643032623239316432633935633164227d5d2c22706f735f7461626c65223a5b342c312c332c312c322c322c312c342c332c312c312c332c332c312c312c342c342c312c332c312c342c322c342c322c332c342c332c342c332c332c312c322c322c332c312c342c312c312c312c322c342c332c332c322c342c322c332c312c322c342c332c322c322c332c342c322c342c322c322c345d2c226d61785f626c6f636b5f6368616e67655f76696577223a36303030307d7d40e80b1c8c5ab0510c27506970c82e462cb11514"),
	//}
	//cp2 := sc.ContractParameter{
	//	Type:  sc.ByteArray,
	//	Value: helper.HexToBytes("12050409c6475ce07577ab72a1f96c263e5030cb53a843b00ca1238a093d9dcb183e2fec837e621b7ec6db7658c9b9808da304aed599043de1b433d490ff74f577c53d12050429e0d1c5b2ae838930ae1ad861ddd3d"),
	//}
	//cp3 := sc.ContractParameter{
	//	Type:  sc.ByteArray,
	//	Value: []byte{},
	//}
	//sb.MakeInvocationScript(scriptHash, CHANGE_BOOK_KEEPER, []sc.ContractParameter{cp1, cp2, cp3})

	//s := sb.ToArray()
	//fmt.Printf(helper.BytesToHex(s))

	s := helper.HexToBytes("004d0c0112050409c6475ce07577ab72a1f96c263e5030cb53a843b00ca1238a093d9dcb183e2fec837e621b7ec6db7658c9b9808da304aed599043de1b433d490ff74f577c53d12050429e0d1c5b2ae838930ae1ad861ddd3d0745d1c7f142492cabd02b291d2c95c1dda6633dc7be5dd4f9597f32f1e45721959d0902a8e56a58b2db79ada7c3ce9321205042bed55e8c4d9cbc50657ff5909ee51dc394a92aad911c36bace83c4d63540794bc68a65f1a54ec4f14a630043090bc29ee9cddf90f3ecb86e0973ffff3fd4899120504e68a6e54bdfa0af47bd18465f4352f5151dc729c61a7399909f1cd1c6d816c0241800e782bb05f6f803b9f958930ebcee0b67d3af27845b4fbfa09e926cf17ae4d65040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000453c810e2d58aeb9aabc22723666785fa200c1d9fea5a5006d9e506df0911d7e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008e305f000000001dac2b7c00000000fdb2037b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a342c2263223a312c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303330396336343735636530373537376162373261316639366332363365353033306362353361383433623030636131323338613039336439646362313833653266227d2c7b22696e646578223a322c226964223a2231323035303332626564353565386334643963626335303635376666353930396565353164633339346139326161643931316333366261636538336334643633353430373934227d2c7b22696e646578223a332c226964223a2231323035303265363861366535346264666130616634376264313834363566343335326635313531646337323963363161373339393930396631636431633664383136633032227d2c7b22696e646578223a342c226964223a2231323035303232396530643163356232616538333839333061653161643836316464643364303734356431633766313432343932636162643032623239316432633935633164227d5d2c22706f735f7461626c65223a5b342c312c332c312c322c322c312c342c332c312c312c332c332c312c312c342c342c312c332c312c342c322c342c322c332c342c332c342c332c332c312c322c322c332c312c342c312c312c312c322c342c332c332c322c342c322c332c312c322c342c332c322c322c332c342c322c342c322c322c345d2c226d61785f626c6f636b5f6368616e67655f76696577223a36303030307d7d40e80b1c8c5ab0510c27506970c82e462cb1151453c1104368616e6765426f6f6b4b6565706572677f25d672e8626d2beaa26f2cb40da6b91f40a382")

	//neoRpcClient := rpc.NewClient("http://seed5.ngd.network:10332")
	tb := tx.NewTransactionBuilder("http://seed6.ngd.network:11332")
	from, err := helper.AddressToScriptHash("AbQMu96ZpzREtEhW6b7gUcCCTWLmpSJSZ9")
	// create an InvocationTransaction
	itx, err := tb.MakeInvocationTransaction(s, from, nil, from, helper.Zero)
	if err != nil {
		fmt.Errorf("[changeBookKeeper] tb.MakeInvocationTransaction error: %s", err)
	}

	// sign transaction
	keyPair, err := keys.NewKeyPairFromNEP2("6PYMp74JupzrVADERufT5Htzx8ohsCGxWNHFdn97ENc1L35aQauwMEg65d", "FJNfSqUZ2ExU5MnJHqya")
	err = tx.AddSignature(itx, keyPair)
	if err != nil {
		fmt.Errorf("[changeBookKeeper] tx.AddSignature error: %s", err)
	}

	rawTxString := itx.RawTransactionString()

	fmt.Printf(rawTxString)
}

func Test_ContainString(t *testing.T)  {
	s := "JsonRpcResponse error code:-1 desc:INTERNAL ERROR, ErrUnknown result:\"[Invoke] Native serivce function execute error:neo MakeDepositProposal, check done transaction error:checkDoneTx, tx already done\""
	subS := "checkDoneTx, tx already done"
	if strings.Contains(s, subS) {
		log.Infof("yes")
	}
}

func Test_SaveLife(t *testing.T) {
	client := rpc.NewClient("http://seed10.ngd.network:20332")
	expectedNextConsensus := "51c320d9459aa6b524456babb2b4a8ac8e432b8a"
	realNextConsensus := "51c320d9459aa6b524456babb2b4a8ac8e432b8a"
	index := uint32(4387120)

	for expectedNextConsensus == realNextConsensus {
		log.Infof(strconv.Itoa(int(index)))
		header := client.GetBlockHeaderByIndex(index)
		scriptHash, err := helper.AddressToScriptHash(header.Result.NextConsensus)
		assert.Nil(t, err)
		realNextConsensus = helper.BytesToHex(helper.ReverseBytes(scriptHash.Bytes()))
		index = index - 10000
	}

	log.Infof(strconv.Itoa(int(index)))
}
