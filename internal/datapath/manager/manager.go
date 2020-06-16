package manager

import "github.com/archa347/modulus/internal/datapath/manager/metadata"

type Manager struct {
	metadataStore metadata.Store
}

func New(metadataStore *metadata.Store) *Manager {
	return &Manager{metadataStore: metadataStore}
}
