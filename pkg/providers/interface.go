package providers

import (
	"errors"
	"github.com/emc-advanced-dev/unik/pkg/types"
)

type Provider interface {
	GetConfig() ProviderConfig
	//Images
	Stage(name string, rawImage *types.RawImage, force bool) (*types.Image, error)
	ListImages() ([]*types.Image, error)
	GetImage(nameOrIdPrefix string) (*types.Image, error)
	DeleteImage(id string, force bool) error
	//Instances
	RunInstance(name, imageId string, mntPointsToVolumeIds map[string]string, env map[string]string) (*types.Instance, error)
	ListInstances() ([]*types.Instance, error)
	GetInstance(nameOrIdPrefix string) (*types.Instance, error)
	DeleteInstance(id string) error
	StartInstance(id string) error
	StopInstance(id string) error
	GetInstanceLogs(id string) (string, error)
	//Volumes
	CreateVolume(name, imagePath string) (*types.Volume, error)
	ListVolumes() ([]*types.Volume, error)
	GetVolume(nameOrIdPrefix string) (*types.Volume, error)
	DeleteVolume(id string, force bool) error
	AttachVolume(id, instanceId, mntPoint string) error
	DetachVolume(id string) error
}

type ProviderConfig struct {
	UsePartitionTables bool
}

type Providers map[string]Provider

func (providers Providers) Keys() []string {
	keys := []string{}
	for providerType := range providers {
		keys = append(keys, providerType)
	}
	return keys
}

func (providers Providers) ProviderForImage(imageId string) (Provider, error) {
	for _, provider := range providers {
		_, err := provider.GetImage(imageId)
		if err == nil {
			return provider, nil
		}
	}
	return nil, errors.New("provider not found for image " + imageId)
}

func (providers Providers) ProviderForInstance(instanceId string) (Provider, error) {
	for _, provider := range providers {
		_, err := provider.GetInstance(instanceId)
		if err == nil {
			return provider, nil
		}
	}
	return nil, errors.New("provider not found for instance " + instanceId)
}

func (providers Providers) ProviderForVolume(volumeId string) (Provider, error) {
	for _, provider := range providers {
		_, err := provider.GetVolume(volumeId)
		if err == nil {
			return provider, nil
		}
	}
	return nil, errors.New("provider not found for volume " + volumeId)
}
