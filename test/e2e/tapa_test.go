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

package e2e_test

import (
	"context"
	"fmt"

	"github.com/nordix/meridio/test/e2e/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("TAPA", func() {

	var (
		targetPod *v1.Pod
	)

	BeforeEach(func() {
		if targetPod != nil {
			return
		}
		listOptions := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("app=%s", config.targetADeploymentName),
		}
		pods, err := clientset.CoreV1().Pods(config.k8sNamespace).List(context.Background(), listOptions)
		Expect(err).ToNot(HaveOccurred())
		Expect(len(pods.Items)).To(BeNumerically(">", 0))
		targetPod = &pods.Items[0]
	})

	Describe("close-open", func() {
		When("Close stream-a-I in one of the target from target-a-deployment-name and re-open it", func() {
			var (
				err error
			)

			BeforeEach(func() {
				_, err := utils.PodExec(targetPod, "example-target", []string{"./target-client", "close", "-t", config.trenchA, "-c", config.conduitA1, "-s", config.streamAI})
				Expect(err).NotTo(HaveOccurred())

				// wait trenchA/conduitA1/streamAI to be closed
				Eventually(func() bool {
					targetWatchOutput, err := utils.PodExec(targetPod, "example-target", []string{"timeout", "--preserve-status", "0.5", "./target-client", "watch"})
					Expect(err).NotTo(HaveOccurred())
					streamStatus := utils.ParseTargetWatch(targetWatchOutput)
					return len(streamStatus) == 0
				}, timeout, interval).Should(BeTrue())

				// wait for all identifiers to be in NFQLB in statelessLbFeDeploymentNameAttractorA1
				listOptions := metav1.ListOptions{
					LabelSelector: fmt.Sprintf("app=%s", config.statelessLbFeDeploymentNameAttractorA1),
				}
				pods, err := clientset.CoreV1().Pods(config.k8sNamespace).List(context.Background(), listOptions)
				Expect(err).NotTo(HaveOccurred())
				Eventually(func() bool {
					for _, pod := range pods.Items {
						nfqlbOutput, err := utils.PodExec(&pod, "stateless-lb", []string{"nfqlb", "show", fmt.Sprintf("--shm=tshm-%v", config.streamAI)})
						Expect(err).NotTo(HaveOccurred())
						if utils.ParseNFQLB(nfqlbOutput) != (numberOfTargetA - 1) {
							return false
						}
					}
					return true
				}, timeout, interval).Should(BeTrue())
			})

			AfterEach(func() {
				_, err = utils.PodExec(targetPod, "example-target", []string{"./target-client", "open", "-t", config.trenchA, "-c", config.conduitA1, "-s", config.streamAI})
				Expect(err).NotTo(HaveOccurred())

				// wait trenchA/conduitA1/streamAI to be opened
				Eventually(func() bool {
					targetWatchOutput, err := utils.PodExec(targetPod, "example-target", []string{"timeout", "--preserve-status", "0.5", "./target-client", "watch"})
					Expect(err).NotTo(HaveOccurred())
					streamStatus := utils.ParseTargetWatch(targetWatchOutput)
					if len(streamStatus) == 1 && streamStatus[0].Status == "OPEN" && streamStatus[0].Trench == config.trenchA && streamStatus[0].Conduit == config.conduitA1 && streamStatus[0].Stream == config.streamAI {
						return true
					}
					return false
				}, timeout, interval).Should(BeTrue())

				// wait for all identifiers to be in NFQLB in statelessLbFeDeploymentNameAttractorA1
				listOptions := metav1.ListOptions{
					LabelSelector: fmt.Sprintf("app=%s", config.statelessLbFeDeploymentNameAttractorA1),
				}
				pods, err := clientset.CoreV1().Pods(config.k8sNamespace).List(context.Background(), listOptions)
				Expect(err).NotTo(HaveOccurred())
				Eventually(func() bool {
					for _, pod := range pods.Items {
						nfqlbOutput, err := utils.PodExec(&pod, "stateless-lb", []string{"nfqlb", "show", fmt.Sprintf("--shm=tshm-%v", config.streamAI)})
						Expect(err).NotTo(HaveOccurred())
						if utils.ParseNFQLB(nfqlbOutput) != numberOfTargetA {
							return false
						}
					}
					return true
				}, timeout, interval).Should(BeTrue())
			})

			It("(Traffic) is received by the targets", func() {
				if !utils.IsIPv6(config.ipFamily) { // Don't send traffic with IPv4 if the tests are only IPv6
					By("Checking IPv4")
					lastingConnections, lostConnections := trafficGeneratorHost.SendTraffic(trafficGenerator, config.trenchA, config.k8sNamespace, utils.VIPPort(config.vip1V4, config.flowAZTcpDestinationPort0), "tcp")
					Expect(lostConnections).To(Equal(0))
					Expect(len(lastingConnections)).To(Equal(numberOfTargetA - 1))
					_, exists := lastingConnections[targetPod.Name]
					Expect(exists).ToNot(BeTrue())
				}
				if !utils.IsIPv4(config.ipFamily) { // Don't send traffic with IPv6 if the tests are only IPv4
					By("Checking IPv6")
					lastingConnections, lostConnections := trafficGeneratorHost.SendTraffic(trafficGenerator, config.trenchA, config.k8sNamespace, utils.VIPPort(config.vip1V6, config.flowAZTcpDestinationPort0), "tcp")
					Expect(lostConnections).To(Equal(0))
					Expect(len(lastingConnections)).To(Equal(numberOfTargetA - 1))
					_, exists := lastingConnections[targetPod.Name]
					Expect(exists).ToNot(BeTrue())
				}
			})
		})
	})

})
