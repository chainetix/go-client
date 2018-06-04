package chainetix

func (client *Client) NewWallet(label string, secret string, user string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"secret": secret,
		"user": user,
		"label": label,
	}
	_, err := client.Post(
		"https://multichain.blokhub.io/api/wallet",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}