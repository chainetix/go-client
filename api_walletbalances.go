package chainetix

func (client *Client) WalletBalances(wallet string) ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.chainetix.com/api/wallet/" + wallet + "/balances",
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}