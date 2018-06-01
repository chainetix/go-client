package chainetix

func (client *Client) DeleteAgents() error {
	src := map[string]interface{}{
	}
	_, err := client.Delete(
		"https://multichain.blokhub.io/api/agents",
		src,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}