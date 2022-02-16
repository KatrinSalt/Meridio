/*
Copyright (c) 2021-2022 Nordix Foundation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package conduit

import (
	"context"
	"io"
	"sync"

	nspAPI "github.com/nordix/meridio/api/nsp/v1"
	"github.com/sirupsen/logrus"
)

type Configuration interface {
	Watch()
	Stop()
}

type configurationImpl struct {
	SetVips                    func([]string) error
	SetStreams                 func([]*nspAPI.Stream)
	Conduit                    *nspAPI.Conduit
	ConfigurationManagerClient nspAPI.ConfigurationManagerClient
	cancel                     context.CancelFunc
	wg                         sync.WaitGroup
	mu                         sync.Mutex
}

func newConfigurationImpl(setVips func([]string) error,
	setStreams func([]*nspAPI.Stream),
	conduit *nspAPI.Conduit,
	configurationManagerClient nspAPI.ConfigurationManagerClient) *configurationImpl {
	c := &configurationImpl{
		SetVips:                    setVips,
		SetStreams:                 setStreams,
		Conduit:                    conduit,
		ConfigurationManagerClient: configurationManagerClient,
	}
	return c
}

func (c *configurationImpl) Watch() {
	c.mu.Lock()
	defer c.mu.Unlock()
	var ctx context.Context
	ctx, c.cancel = context.WithCancel(context.TODO())
	go c.watchVIPs(ctx)
	go c.watchStreams(ctx)
}

func (c *configurationImpl) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cancel != nil {
		c.cancel()
	}
	c.wg.Wait()
}

func (c *configurationImpl) watchVIPs(ctx context.Context) {
	c.wg.Add(1)
	defer c.wg.Done()
	for { // Todo: retry
		if ctx.Err() != nil {
			return
		}
		vipsToWatch := &nspAPI.Vip{
			Trench: c.Conduit.GetTrench(),
		}
		watchVIPClient, err := c.ConfigurationManagerClient.WatchVip(ctx, vipsToWatch)
		if err != nil {
			logrus.Warnf("err watchVIPClient.Recv: %v", err) // todo
			continue
		}
		for {
			vipResponse, err := watchVIPClient.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				logrus.Warnf("err watchVIPClient.Recv: %v", err) // todo
				break
			}
			err = c.SetVips(vipResponse.ToSlice())
			if err != nil {
				logrus.Warnf("err set vips: %v", err) // todo
			}
		}
	}
}

func (c *configurationImpl) watchStreams(ctx context.Context) {
	c.wg.Add(1)
	defer c.wg.Done()
	for { // Todo: retry
		if ctx.Err() != nil {
			return
		}
		vipsToWatch := &nspAPI.Stream{
			Conduit: c.Conduit,
		}
		watchStreamClient, err := c.ConfigurationManagerClient.WatchStream(ctx, vipsToWatch)
		if err != nil {
			logrus.Warnf("err watchVIPClient.Recv: %v", err) // todo
			continue
		}
		for {
			streamResponse, err := watchStreamClient.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				logrus.Warnf("err watchVIPClient.Recv: %v", err) // todo
				break
			}
			c.SetStreams(streamResponse.GetStreams())
			// err = c.Watcher.SetVIPs(vipResponse.ToSlice())
			// if err != nil {
			// 	logrus.Warnf("err set vips: %v", err) // todo
			// }
		}
	}
}