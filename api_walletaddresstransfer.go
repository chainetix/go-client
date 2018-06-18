package chainetix

func (client *Client) WalletAddressTransfer(address string, assetQuantity float64, currencyName string, recipientAddress string, secret string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"assetQuantity": assetQuantity,
		"currencyName": currencyName,
		"recipientAddress": recipientAddress,
		"secret": secret,
	}
	_, err := client.Post(
		"https://multichain.chainetix.com/api/address/" + address + "/transfer",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}