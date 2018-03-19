package store

// DataEntryPrefix
type DataEntryPrefix byte

const (
	// DATA
	DATA_BlockHash   DataEntryPrefix = 0x00
	DATA_Header      DataEntryPrefix = 0x01
	DATA_Transaction DataEntryPrefix = 0x02

	// INDEX
	IX_HeaderHashList DataEntryPrefix = 0x80
	IX_Unspent        DataEntryPrefix = 0x90
	IX_Unspent_UTXO   DataEntryPrefix = 0x91

	// ASSET
	ST_Info           DataEntryPrefix = 0xc0
	ST_Storage        DataEntryPrefix = 0xc3
	ST_Contract        DataEntryPrefix = 0xc4

	//SYSTEM
	SYS_CurrentBlock  DataEntryPrefix = 0x40

	//CONFIG
	CFG_Version DataEntryPrefix = 0xf0
)
