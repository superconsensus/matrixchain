package common

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	"github.com/superconsensus/matrixchain/service/pb"
	"github.com/superconsensus/matrixcore/bcs/ledger/xledger/xldgpb"
	"github.com/superconsensus/matrixcore/protos"
)

// 为了完全兼容老版本pb结构，转换交易结构
func TxToXledger(tx *pb.Transaction) *xldgpb.Transaction {
	if tx == nil {
		return nil
	}

	prtBuf, err := proto.Marshal(tx)
	if err != nil {
		return nil
	}

	var newTx xldgpb.Transaction
	err = proto.Unmarshal(prtBuf, &newTx)
	if err != nil {
		return nil
	}

	return &newTx
}

// 为了完全兼容老版本pb结构，转换交易结构
func TxToXchain(tx *xldgpb.Transaction) *pb.Transaction {
	if tx == nil {
		return nil
	}

	prtBuf, err := proto.Marshal(tx)
	if err != nil {
		return nil
	}

	var newTx pb.Transaction
	err = proto.Unmarshal(prtBuf, &newTx)
	if err != nil {
		return nil
	}

	return &newTx
}

// 为了完全兼容老版本pb结构，转换区块结构
func BlockToXledger(block *pb.InternalBlock) *xldgpb.InternalBlock {
	if block == nil {
		return nil
	}

	blkBuf, err := proto.Marshal(block)
	if err != nil {
		return nil
	}

	var newBlock xldgpb.InternalBlock
	err = proto.Unmarshal(blkBuf, &newBlock)
	if err != nil {
		return nil
	}

	return &newBlock
}

// 为了完全兼容老版本pb结构，转换区块结构
func BlockToXchain(block *xldgpb.InternalBlock) *pb.InternalBlock {
	if block == nil {
		return nil
	}

	blkBuf, err := proto.Marshal(block)
	if err != nil {
		return nil
	}

	var newBlock pb.InternalBlock
	err = proto.Unmarshal(blkBuf, &newBlock)
	if err != nil {
		return nil
	}

	return &newBlock
}

func TestToXchain(CandidateRatio *protos.CandidateRatio) *pb.CandidateRatio {
	if CandidateRatio == nil {
		return nil
	}
	CandidateRatioBuff, err := proto.Marshal(CandidateRatio)
	if err != nil {
		return nil
	}
	var newCandidateRatio pb.CandidateRatio
	err = proto.Unmarshal(CandidateRatioBuff, &newCandidateRatio)
	if err != nil {
		return nil
	}

	return &newCandidateRatio
}

func BonusToXchain(BonusQueryReply *protos.BonusQueryReply) *pb.BonusQueryReply {
	if BonusQueryReply == nil {
		return nil
	}
	BonusQueryReplyBuff, err := proto.Marshal(BonusQueryReply)
	if err != nil {
		return nil
	}
	var newBonusQueryReply pb.BonusQueryReply
	err = proto.Unmarshal(BonusQueryReplyBuff, &newBonusQueryReply)
	if err != nil {
		return nil
	}

	return &newBonusQueryReply
}

func GetVerificationToXchain(VerificationTable *protos.VerificationTable) *pb.VerificationTable {
	if VerificationTable == nil {
		return nil
	}
	VerificationTableBUff, err := proto.Marshal(VerificationTable)
	if err != nil {

		return nil
	}

	var newVerificationTableBUff pb.VerificationTable
	err = proto.Unmarshal(VerificationTableBUff, &newVerificationTableBUff)
	if err != nil {
		fmt.Printf("D__打印转换错误:%s \n", err)
		return nil
	}
	return &newVerificationTableBUff
}

func GetSystemStatusExplorerToXchain(BCStatusExplorer *protos.BCStatusExplorer) *pb.BCStatusExplorer {
	if BCStatusExplorer == nil {
		return nil
	}
	BCStatusExplorerBUff, err := proto.Marshal(BCStatusExplorer)
	if err != nil {
		return nil
	}

	var newBCStatusExplorerBUff pb.BCStatusExplorer
	err = proto.Unmarshal(BCStatusExplorerBUff, &newBCStatusExplorerBUff)
	if err != nil {
		return nil
	}

	return &newBCStatusExplorerBUff
}

func VotingRecordsToXchain(VotingResponse *protos.PledgeVotingResponse) *pb.PledgeVotingResponse {
	if VotingResponse == nil {
		return nil
	}
	//fmt.Printf("D__打印转换前数据VotingResponse.FrozenAssetsTable：%s \n",VotingResponse.FrozenAssetsTable)
	//fmt.Printf("D__打印转换前数据VotingResponse.VoteDetailsStatus：%s \n",VotingResponse.VoteDetailsStatus)
	//fmt.Printf("D__打印转换前数据VotingResponse.Freezetotal：%s \n",VotingResponse.Freezetotal)

	VotingResponseBuff, err := proto.Marshal(VotingResponse)
	if err != nil {
		fmt.Printf("D__打印解析错误:%s \n", err)
		return nil
	}

	var newVotingResponseBuff pb.PledgeVotingResponse

	err = proto.Unmarshal(VotingResponseBuff, &newVotingResponseBuff)
	if err != nil {
		fmt.Printf("D__打印转换错误:%s \n", err)
		return nil
	}
	return &newVotingResponseBuff
}

func ConvertInvokeReq(reqs []*pb.InvokeRequest) ([]*protos.InvokeRequest, error) {
	if reqs == nil {
		return nil, nil
	}

	newReqs := make([]*protos.InvokeRequest, 0, len(reqs))
	for _, req := range reqs {
		buf, err := proto.Marshal(req)
		if err != nil {
			return nil, err
		}

		var tmp protos.InvokeRequest
		err = proto.Unmarshal(buf, &tmp)
		if err != nil {
			return nil, err
		}

		newReqs = append(newReqs, &tmp)
	}

	return newReqs, nil
}

func ConvertInvokeResp(resp *protos.InvokeResponse) *pb.InvokeResponse {
	if resp == nil {
		return nil
	}

	buf, err := proto.Marshal(resp)
	if err != nil {
		return nil
	}

	var tmp pb.InvokeResponse
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func UtxoToXchain(utxo *xldgpb.Utxo) *pb.Utxo {
	if utxo == nil {
		return nil
	}

	buf, err := proto.Marshal(utxo)
	if err != nil {
		return nil
	}

	var tmp pb.Utxo
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func UtxoToXledger(utxo *pb.Utxo) *xldgpb.Utxo {
	if utxo == nil {
		return nil
	}

	buf, err := proto.Marshal(utxo)
	if err != nil {
		return nil
	}

	var tmp xldgpb.Utxo
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func UtxoListToXchain(utxoList []*xldgpb.Utxo) ([]*pb.Utxo, error) {
	if utxoList == nil {
		return nil, nil
	}

	tmpList := make([]*pb.Utxo, 0, len(utxoList))
	for _, utxo := range utxoList {
		tmp := UtxoToXchain(utxo)
		if tmp == nil {
			return nil, fmt.Errorf("convert utxo failed")
		}
		tmpList = append(tmpList, tmp)
	}

	return tmpList, nil
}

func UtxoRecordToXchain(record *xldgpb.UtxoRecord) *pb.UtxoRecord {
	if record == nil {
		return nil
	}

	newRecord := &pb.UtxoRecord{
		UtxoCount:  record.GetUtxoCount(),
		UtxoAmount: record.GetUtxoAmount(),
	}
	if record.GetItem() == nil {
		return newRecord
	}

	newRecord.Item = make([]*pb.UtxoKey, 0, len(record.GetItem()))
	for _, item := range record.GetItem() {
		tmp := &pb.UtxoKey{
			RefTxid: item.GetRefTxid(),
			Offset:  item.GetOffset(),
			Amount:  item.GetAmount(),
		}
		newRecord.Item = append(newRecord.Item, tmp)
	}

	return newRecord
}

func AclToXchain(acl *protos.Acl) *pb.Acl {
	if acl == nil {
		return nil
	}

	buf, err := proto.Marshal(acl)
	if err != nil {
		return nil
	}

	var tmp pb.Acl
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func ContractStatusToXchain(contractStatus *protos.ContractStatus) *pb.ContractStatus {
	if contractStatus == nil {
		return nil
	}

	buf, err := proto.Marshal(contractStatus)
	if err != nil {
		return nil
	}

	var tmp pb.ContractStatus
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func ContractStatusListToXchain(contractStatusList []*protos.ContractStatus) ([]*pb.ContractStatus, error) {
	if contractStatusList == nil {
		return nil, nil
	}

	tmpList := make([]*pb.ContractStatus, 0, len(contractStatusList))
	for _, cs := range contractStatusList {
		tmp := ContractStatusToXchain(cs)
		if tmp == nil {
			return nil, fmt.Errorf("convert contract status failed")
		}
		tmpList = append(tmpList, tmp)
	}

	return tmpList, nil
}

func PeerInfoToStrings(info protos.PeerInfo) []string {
	peerUrls := make([]string, 0, len(info.Peer))
	for _, peer := range info.Peer {
		peerUrls = append(peerUrls, peer.Address)
	}
	return peerUrls
}

func BalanceDetailToXchain(detail *xldgpb.BalanceDetailInfo) *pb.TokenFrozenDetail {
	if detail == nil {
		return nil
	}

	buf, err := proto.Marshal(detail)
	if err != nil {
		return nil
	}

	var tmp pb.TokenFrozenDetail
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func BalanceDetailsToXchain(details []*xldgpb.BalanceDetailInfo) ([]*pb.TokenFrozenDetail, error) {
	if details == nil {
		return nil, nil
	}

	tmpList := make([]*pb.TokenFrozenDetail, 0, len(details))
	for _, detail := range details {
		tmp := BalanceDetailToXchain(detail)
		if tmp == nil {
			return nil, fmt.Errorf("convert balance detail failed")
		}
		tmpList = append(tmpList, tmp)
	}

	return tmpList, nil
}

func LedgerMetaToXchain(meta *xldgpb.LedgerMeta) *pb.LedgerMeta {
	if meta == nil {
		return nil
	}

	buf, err := proto.Marshal(meta)
	if err != nil {
		return nil
	}

	var tmp pb.LedgerMeta
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func UtxoMetaToXchain(meta *xldgpb.UtxoMeta) *pb.UtxoMeta {
	if meta == nil {
		return nil
	}

	buf, err := proto.Marshal(meta)
	if err != nil {
		return nil
	}

	var tmp pb.UtxoMeta
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func ConvertEventSubType(typ pb.SubscribeType) protos.SubscribeType {
	switch typ {
	case pb.SubscribeType_BLOCK:
		return protos.SubscribeType_BLOCK
	}

	return protos.SubscribeType_BLOCK
}
