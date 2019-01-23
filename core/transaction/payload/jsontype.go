package payload

type CommitInfo struct {
	SigChain  string `json:"sigChain"`
	Submitter string `json:"submitter"`
}

type RegisterNameInfo struct {
	Registrant string `json:"registrant"`
	Name       string `json:"name"`
}

type DeleteNameInfo struct {
	Registrant string `json:"registrant"`
}

type SubscribeInfo struct {
	Subscriber string `json:"subscriber"`
	Identifier string `json:"identifier"`
	Topic      string `json:"topic"`
	Duration   uint32 `json:"duration"`
}

type TransferAssetInfo struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

type CoinbaseInfo struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}
