package main

type fetcher struct {
	metaData map[string]*string
}

func (m *fetcher) Fetch(key string) string {
	return *m.metaData[key]
}
