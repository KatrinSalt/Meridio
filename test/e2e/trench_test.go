/*
Copyright (c) 2023 Nordix Foundation

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

package e2e_test

import (
	"context"
	"fmt"

	"github.com/nordix/meridio/test/e2e/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trench", func() {

	Describe("delete-create-trench", func() {
		When("Delete trench-a and recreate and reconfigure it", func() {
			BeforeEach(func() {
				By(fmt.Sprintf("Deleting the trench %s", config.trenchA))
				err := utils.Exec(config.script, "delete_create_trench")
				Expect(err).ToNot(HaveOccurred())
				By(fmt.Sprintf("Recreate the trench %s", config.trenchA))
				err = utils.Exec(config.script, "delete_create_trench_revert")
				Expect(err).ToNot(HaveOccurred())
			})

			It("(Traffic) is received by the targets", func(ctx context.Context) {
				protocol := "tcp"
				if !utils.IsIPv6(config.ipFamily) { // Don't send traffic with IPv4 if the tests are only IPv6
					ipPort := utils.VIPPort(config.vip1V4, config.flowAZTcpDestinationPort0)
					By(fmt.Sprintf("Sending %s traffic from the TG %s (%s) to %s", protocol, config.trenchA, config.k8sNamespace, ipPort))
					lastingConnections, lostConnections := trafficGeneratorHost.SendTraffic(trafficGenerator, config.trenchA, config.k8sNamespace, ipPort, protocol)
					Expect(lostConnections).To(Equal(0), "There should be no lost connection: %v", lastingConnections)
					Expect(len(lastingConnections)).To(Equal(numberOfTargetA), "All targets with the stream opened should have received traffic: %v", lastingConnections)
				}
				if !utils.IsIPv4(config.ipFamily) { // Don't send traffic with IPv6 if the tests are only IPv4
					ipPort := utils.VIPPort(config.vip1V6, config.flowAZTcpDestinationPort0)
					By(fmt.Sprintf("Sending %s traffic from the TG %s (%s) to %s", protocol, config.trenchA, config.k8sNamespace, ipPort))
					lastingConnections, lostConnections := trafficGeneratorHost.SendTraffic(trafficGenerator, config.trenchA, config.k8sNamespace, ipPort, protocol)
					Expect(lostConnections).To(Equal(0), "There should be no lost connection: %v", lastingConnections)
					Expect(len(lastingConnections)).To(Equal(numberOfTargetA), "All targets with the stream opened should have received traffic: %v", lastingConnections)
				}
			}, SpecTimeout(timeoutTest))
		})
	})

})
