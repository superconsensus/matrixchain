package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/superconsensus/matrixchain/data/mock"
	scom "github.com/superconsensus/matrixchain/service/common"
	"github.com/superconsensus/matrixchain/service/pb"

	// import要使用的内核核心组件驱动
	_ "github.com/superconsensus/matrixcore/bcs/consensus/pow"
	_ "github.com/superconsensus/matrixcore/bcs/consensus/single"
	_ "github.com/superconsensus/matrixcore/bcs/consensus/tdpos"
	_ "github.com/superconsensus/matrixcore/bcs/consensus/xpoa"
	_ "github.com/superconsensus/matrixcore/bcs/contract/evm"
	_ "github.com/superconsensus/matrixcore/bcs/contract/native"
	_ "github.com/superconsensus/matrixcore/bcs/contract/xvm"
	txn "github.com/superconsensus/matrixcore/bcs/ledger/xledger/tx"
	xledger "github.com/superconsensus/matrixcore/bcs/ledger/xledger/utils"
	_ "github.com/superconsensus/matrixcore/bcs/network/p2pv1"
	_ "github.com/superconsensus/matrixcore/bcs/network/p2pv2"
	xconf "github.com/superconsensus/matrixcore/kernel/common/xconfig"
	_ "github.com/superconsensus/matrixcore/kernel/contract/kernel"
	_ "github.com/superconsensus/matrixcore/kernel/contract/manager"
	"github.com/superconsensus/matrixcore/kernel/engines/xuperos"
	"github.com/superconsensus/matrixcore/kernel/engines/xuperos/common"
	_ "github.com/superconsensus/matrixcore/lib/crypto/client"
	"github.com/superconsensus/matrixcore/lib/logs"
	_ "github.com/superconsensus/matrixcore/lib/storage/kvdb/leveldb"
)

var (
	address   = "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"
	publickey = "{\"Curvname\":\"P-256\",\"X\":74695617477160058757747208220371236837474210247114418775262229497812962582435,\"Y\":51348715319124770392993866417088542497927816017012182211244120852620959209571}"
)

func TestEndorserCall(t *testing.T) {
	workspace, dirErr := ioutil.TempDir("/tmp", "")
	if dirErr != nil {
		t.Fatal(dirErr)
	}
	os.RemoveAll(workspace)
	defer os.RemoveAll(workspace)
	conf, _ := mock.NewEnvConfForTest()
	defer RemoveLedger(conf)

	engine, err := MockEngine()
	if err != nil {
		t.Fatal(err)
	}
	log, _ := logs.NewLogger("", scom.SubModName)
	rpcServ := NewRpcServ(engine, log)

	endor := NewDefaultXEndorser(rpcServ, engine)
	awardTx, err := txn.GenerateAwardTx("miner", "1000", []byte("award"))

	txStatus := &pb.TxStatus{
		Bcname: "xuper",
		Tx:     scom.TxToXchain(awardTx),
	}
	requestData, err := json.Marshal(txStatus)
	if err != nil {
		fmt.Printf("json encode txStatus failed: %v", err)
		t.Fatal(err)
	}
	ctx := context.TODO()
	req := &pb.EndorserRequest{
		RequestName: "ComplianceCheck",
		BcName:      "xuper",
		Fee:         nil,
		RequestData: requestData,
	}
	resp, err := endor.EndorserCall(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(resp)
	invokeReq := make([]*pb.InvokeRequest, 0)
	invoke := &pb.InvokeRequest{
		ModuleName:   "wasm",
		ContractName: "counter",
		MethodName:   "increase",
		Args:         map[string][]byte{"key": []byte("test")},
	}
	invokeReq = append(invokeReq, invoke)
	preq := &pb.PreExecWithSelectUTXORequest{
		Bcname:      "xuper",
		Address:     address,
		TotalAmount: 100,
		SignInfo: &pb.SignatureInfo{
			PublicKey: publickey,
			Sign:      []byte("sign"),
		},
		NeedLock: false,
		Request: &pb.InvokeRPCRequest{
			Bcname:      "xuper",
			Requests:    invokeReq,
			Initiator:   address,
			AuthRequire: []string{address},
		},
	}

	reqJSON, _ := json.Marshal(preq)
	xreq := &pb.EndorserRequest{
		RequestName: "PreExecWithFee",
		BcName:      "xuper",
		Fee:         nil,
		RequestData: reqJSON,
	}
	resp, err = endor.EndorserCall(ctx, xreq)
	if err != nil {
		//pass
		t.Log(err)
	}
	t.Log(resp)
	qtxTxStatus := &pb.TxStatus{
		Bcname: "xuper",
		Txid:   []byte("70c64d6cb9b5647048d067c6775575fc52e3c51c6425cec3881d8564ad8e887c"),
	}
	requestData, err = json.Marshal(qtxTxStatus)
	if err != nil {
		fmt.Printf("json encode txStatus failed: %v", err)
		t.Fatal(err)
	}
	req = &pb.EndorserRequest{
		RequestName: "TxQuery",
		BcName:      "xuper",
		RequestData: requestData,
	}
	resp, err = endor.EndorserCall(ctx, req)
	if err != nil {
		t.Log(err)
	}
	t.Log(resp)
}

func MockEngine() (common.Engine, error) {
	conf, err := mock.NewEnvConfForTest()
	if err != nil {
		return nil, fmt.Errorf("new env conf error: %v", err)
	}

	RemoveLedger(conf)
	if err = CreateLedger(conf); err != nil {
		return nil, err
	}

	engine := xuperos.NewEngine()
	if err := engine.Init(conf); err != nil {
		return nil, fmt.Errorf("init engine error: %v", err)
	}

	eng, err := xuperos.EngineConvert(engine)
	if err != nil {
		return nil, fmt.Errorf("engine convert error: %v", err)
	}

	return eng, nil
}

func RemoveLedger(conf *xconf.EnvConf) error {
	path := conf.GenDataAbsPath("blockchain")
	if err := os.RemoveAll(path); err != nil {
		log.Printf("remove ledger failed.err:%v\n", err)
		return err
	}
	return nil
}

func CreateLedger(conf *xconf.EnvConf) error {
	mockConf, err := mock.NewEnvConfForTest()
	if err != nil {
		return fmt.Errorf("new mock env conf error: %v", err)
	}

	genesisPath := mockConf.GenDataAbsPath("genesis/xuper.json")
	err = xledger.CreateLedger("xuper", genesisPath, conf)
	if err != nil {
		log.Printf("create ledger failed.err:%v\n", err)
		return fmt.Errorf("create ledger failed")
	}
	return nil
}
