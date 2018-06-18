package chainetix

func (client *Client) DeleteAgents() error {
	src := map[string]interface{}{
	}
	_, err := client.Delete(
		"https://multichain.chainetix.com/api/agents",
		src,
		nil,
		client.headers,
	)
	if err != nil {
		return err
	}
	return nil
}