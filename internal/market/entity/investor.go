package entity

type Investor struct {
	ID             string
	Name           string
	AssetPossition []*InvestorAssetPossition
}

type InvestorAssetPossition struct {
	AssetID string
	Shares  int
}
