package chainetix

func (client *Client) WalletAddresses(wallet string) ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/wallet/" + wallet + "/addresses",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}