package chainetix

func (client *Client) IssuedAssets(currency string) ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/currency/" + currency + "/assets",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}