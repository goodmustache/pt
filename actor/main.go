package actor

type Main struct {
	API    TrackerClient
	Config Config
}

func NewActor(api TrackerClient, config Config) *Main {
	return &Main{
		API:    api,
		Config: config,
	}
}
