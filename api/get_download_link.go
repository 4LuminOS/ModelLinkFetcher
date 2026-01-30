package api

type DownloadLinkConfig struct {
	ModelPath ModelPath
	Digest    string
}

func (c *DownloadLinkConfig) GetDownloadLink() string {
	baseUrl := c.ModelPath.BaseURL()
	requestURL := baseUrl.JoinPath("v2", c.ModelPath.GetNamespaceRepository(), "blobs", c.Digest)
	return requestURL.String()
}
