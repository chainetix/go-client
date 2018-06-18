package chainetix

func (client *Client) WalletTransfer(assetQuantity float64, currencyName string, recipientAddress string, secret string, wallet string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"recipientAddress": recipientAddress,
		"secret": secret,
		"assetQuantity": assetQuantity,
		"currencyName": currencyName,
	}
	_, err := client.Post(
		"https://multichain.chainetix.com/api/wallet/" + wallet + "/transfer",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}