//go:build !no_logs

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tag

import (
	"sync"
	"time"

	logsconfig "github.com/F1997/categraf/config/logs"
)

// Provider returns a list of up-to-date tags for a given entity.
type Provider interface {
	GetTags() []string
}

// provider provides a list of up-to-date tags for a given entity by calling the tagger.
type provider struct {
	entityID             string
	taggerWarmupDuration time.Duration
	localTagProvider     Provider
	sync.Once
}

// NewProvider returns a new Provider.
func NewProvider(entityID string) Provider {
	p := &provider{
		entityID:             entityID,
		taggerWarmupDuration: logsconfig.TaggerWarmupDuration(),
		localTagProvider:     NewLocalProvider([]string{}),
	}

	return p
}

// GetTags returns the list of up-to-date tags.
func (p *provider) GetTags() []string {
	p.Do(func() {
		// Warmup duration
		// Make sure the tagger collects all the service tags
		// TODO: remove this once AD and Tagger use the same PodWatcher instance
		<-time.After(p.taggerWarmupDuration)

	})

	// tags, err := tagger.Tag(p.entityID, collectors.HighCardinality)
	// if err != nil {
	// 	log.Printf("Cannot tag container %s: %v\n", p.entityID, err)
	// 	return []string{}
	// }

	tags := []string{}
	localTags := p.localTagProvider.GetTags()
	if localTags != nil {
		tags = append(tags, localTags...)
	}

	return tags
}
