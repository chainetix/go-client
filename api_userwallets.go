package chainetix

func (client *Client) UserWallets(user string) ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.chainetix.com/api/user/" + user + "/wallets",
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}