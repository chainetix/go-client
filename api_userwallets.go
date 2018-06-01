package chainetix

func (client *Client) UserWallets(user string) ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/user/" + user + "/wallets",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}