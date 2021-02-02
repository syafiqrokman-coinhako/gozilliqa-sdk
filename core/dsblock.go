package core

type DsBlock struct {
	BlockHeader DsBlockHeader
}

// https://github.com/Zilliqa/Zilliqa/blob/04162ef0c3c1787ebbd940b7abd6b7ff1d4714ed/src/libData/BlockData/BlockHeader/DSBlockHeader.h
type DsBlockHeader struct {
	DsDifficulty uint8
	Difficulty   uint8
	// The one who proposed this DS block
	// base16 string
	LeaderPubKey string
	// Block index, starting from 0 in the genesis block
	BlockNum uint64
	// Tx Epoch Num then the DS block was generated
	EpochNum     uint64
	GasPrice     string
	PoWDSWinners []string
	SwInfo       SWInfo
	// key is (base16) public key
	RemoveDSNodePubKeys map[string]Peer
	// todo concrete data type
	DSBlockHashSet     interface{}
	GovDSShardVotesMap interface{}
}

type SWInfo struct {
	ZilliqaMajorVersion uint32
	ZilliqaMinorVersion uint32
	ZilliqaFixVersion   uint32
	ZilliqaUpgradeDS    uint32
	ZilliqaCommit       uint32
	ScillaMajorVersion  uint32
	ScillaMinorVersion  uint32
	ScillaFixVersion    uint32
	ScillaUpgradeDS     uint32
	ScillaCommit        uint32
}

// ds block transfer struct (via rpc)
type DsBlockT struct {
	B1         []bool        `json:"B1"`
	B2         []bool        `json:"B2"`
	CS1        string        `json:"CS1"`
	Header     DsBlockHeaderT `json:"header"`
	Serialized SerializedT   `json:"serialized"`
	Signatures string        `json:"signatures"`
}

type DsBlockHeaderT struct {
	BlockNum       string
	CommitteeHash  string
	Difficulty     uint
	DifficultyDS   uint
	EpochNum       string
	GasPrice       string
	LeaderPubKey   string
	MembersEjected []string
	PoWWinners     []string
	PoWWinnersIP   []interface{}
	PrevHash       string
	ReservedField  string
	SWInfo         SWInfoT
	ShardingHash   string
	Timestamp      string
	Version        uint
}

type SWInfoT struct {
	Scilla  []interface{}
	Zilliqa []interface{}
}

type SerializedT struct {
	Data   string `json:"data"`
	Header string `json:"header"`
}
