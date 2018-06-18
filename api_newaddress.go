package chainetix

func (client *Client) NewAddress(secret string, wallet string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"secret": secret,
		"wallet": wallet,
	}
	_, err := client.Post(
		"https://multichain.chainetix.com/api/address",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}