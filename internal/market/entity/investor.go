package entity

type Investor struct {
	ID             string
	Name           string
	AssetPossition []*InvestorAssetPossition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:             id,
		AssetPossition: []*InvestorAssetPossition{},
	}
}

func (i *Investor) AddAssetPossition(assetPossition *InvestorAssetPossition) {
	i.AssetPossition = append(i.AssetPossition, assetPossition)
}

func (i *Investor) UpdateAssetPossition(assetID string, qtdShares int) {
	assetPossition := i.GetAssetPossition(assetID)
	if assetPossition == nil {
		i.AssetPossition = append(i.AssetPossition, NewInvestorAssetPossition(assetID, qtdShares))
	} else {
		assetPossition.Shares = qtdShares
	}
}

func (i *Investor) GetAssetPossition(assetID string) *InvestorAssetPossition {
	for _, assetPossition := range i.AssetPossition {
		if assetPossition.AssetID == assetID {
			return assetPossition
		}
	}
	return nil
}

type InvestorAssetPossition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPossition(assetID string, shares int) *InvestorAssetPossition {
	return &InvestorAssetPossition{
		AssetID: assetID,
		Shares:  shares,
	}
}
