package chainetix

func (client *Client) NewAddress(secret string, wallet string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"secret": secret,
		"wallet": wallet,
	}
	_, err := client.Post(
		"https://multichain.blokhub.io/api/address",
		src,
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}