package crawler

import (
	"testing"
	"time"

	"github.com/elastic/filebeat/config"
	"github.com/stretchr/testify/assert"
)

func TestProspectorInit(t *testing.T) {

	prospectorConfig := config.ProspectorConfig{
		ScanFrequency: "15s",
		IgnoreOlder:   "100m",
		Harvester: config.HarvesterConfig{
			BufferSize: 100,
			TailFiles:  true,
		},
	}

	prospector := Prospector{
		ProspectorConfig: prospectorConfig,
	}

	assert.NotNil(t, prospector)

	prospector.Init()

	// Predefined values expected
	assert.Equal(t, 100*time.Minute, prospector.ProspectorConfig.IgnoreOlderDuration)
	assert.Equal(t, 15*time.Second, prospector.ProspectorConfig.ScanFrequencyDuration)
	assert.Equal(t, 100, prospector.ProspectorConfig.Harvester.BufferSize)
	assert.Equal(t, true, prospector.ProspectorConfig.Harvester.TailFiles)
}

func TestProspectorInitEmpty(t *testing.T) {

	prospectorConfig := config.ProspectorConfig{
		ScanFrequency: "",
		IgnoreOlder:   "",
		Harvester: config.HarvesterConfig{
			BufferSize: 0,
		},
	}

	prospector := Prospector{
		ProspectorConfig: prospectorConfig,
	}

	prospector.Init()

	// Default values expected
	assert.Equal(t, config.DefaultIgnoreOlderDuration, prospector.ProspectorConfig.IgnoreOlderDuration)
	assert.Equal(t, config.DefaultScanFrequency, prospector.ProspectorConfig.ScanFrequencyDuration)
}

func TestProspectorInitNotSet(t *testing.T) {

	prospectorConfig := config.ProspectorConfig{}

	prospector := Prospector{
		ProspectorConfig: prospectorConfig,
	}

	prospector.Init()

	// Default values expected
	assert.Equal(t, config.DefaultIgnoreOlderDuration, prospector.ProspectorConfig.IgnoreOlderDuration)
	assert.Equal(t, config.DefaultScanFrequency, prospector.ProspectorConfig.ScanFrequencyDuration)
	assert.Equal(t, config.DefaultHarvesterBufferSize, prospector.ProspectorConfig.Harvester.BufferSize)
	assert.Equal(t, config.DefaultTailFiles, prospector.ProspectorConfig.Harvester.TailFiles)
	assert.Equal(t, config.DefaultBackoff, prospector.ProspectorConfig.Harvester.BackoffDuration)
	assert.Equal(t, config.DefaultBackoffFactor, prospector.ProspectorConfig.Harvester.BackoffFactor)
	assert.Equal(t, config.DefaultMaxBackoff, prospector.ProspectorConfig.Harvester.MaxBackoffDuration)
	assert.Equal(t, config.DefaultPartialLineWaiting, prospector.ProspectorConfig.Harvester.PartialLineWaitingDuration)
	assert.Equal(t, config.DefaultForceCloseWindowsFiles, prospector.ProspectorConfig.Harvester.ForceCloseWindowsFiles)
}

func TestProspectorInitScanFrequency0(t *testing.T) {

	prospectorConfig := config.ProspectorConfig{
		ScanFrequency: "0s",
	}

	prospector := Prospector{
		ProspectorConfig: prospectorConfig,
	}

	prospector.Init()

	var zero time.Duration = 0
	// 0 expected
	assert.Equal(t, zero, prospector.ProspectorConfig.ScanFrequencyDuration)
}

func TestProspectorInitInvalidScanFrequency(t *testing.T) {

	prospectorConfig := config.ProspectorConfig{
		ScanFrequency: "abc",
	}

	prospector := Prospector{
		ProspectorConfig: prospectorConfig,
	}

	err := prospector.Init()
	assert.NotNil(t, err)
}

func TestProspectorInitInvalidIgnoreOlder(t *testing.T) {

	prospectorConfig := config.ProspectorConfig{
		IgnoreOlder: "abc",
	}

	prospector := Prospector{
		ProspectorConfig: prospectorConfig,
	}

	err := prospector.Init()
	assert.NotNil(t, err)
}
