package chainetix

func (client *Client) WalletAddresses(wallet string) ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.chainetix.com/api/wallet/" + wallet + "/addresses",
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}