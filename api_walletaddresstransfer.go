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
		"https://multichain.blokhub.io/api/address/" + address + "/transfer",
		src,
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}