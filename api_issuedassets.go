package chainetix

func (client *Client) IssuedAssets(currency string) ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.chainetix.com/api/currency/" + currency + "/assets",
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}