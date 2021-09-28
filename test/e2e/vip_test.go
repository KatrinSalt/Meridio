package e2e

import (
	"fmt"

	meridiov1alpha1 "github.com/nordix/meridio-operator/api/v1alpha1"
	"github.com/nordix/meridio-operator/controllers/common"
	configutils "github.com/nordix/meridio-operator/controllers/config"
	config "github.com/nordix/meridio/pkg/configuration/reader"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Vip", func() {
	trench := &meridiov1alpha1.Trench{
		ObjectMeta: metav1.ObjectMeta{
			Name:      trenchName,
			Namespace: namespace,
		},
		Spec: meridiov1alpha1.TrenchSpec{
			IPFamily: "DualStack",
		},
	}

	vipA := &meridiov1alpha1.Vip{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "vip-a",
			Namespace: namespace,
			Labels: map[string]string{
				"trench": trenchName,
			},
		},
		Spec: meridiov1alpha1.VipSpec{
			Address: "10.0.0.0/28",
		},
	}
	configmapName := fmt.Sprintf("%s-%s", common.CMName, trench.ObjectMeta.Name)

	BeforeEach(func() {
		fw.CleanUpTrenches()
		fw.CleanUpAttractors()
		fw.CleanUpVips()
	})

	AfterEach(func() {
		fw.CleanUpTrenches()
		fw.CleanUpAttractors()
		fw.CleanUpVips()
	})

	Context("When creating a vip", func() {

		AfterEach(func() {
			fw.CleanUpTrenches()
			fw.CleanUpVips()
		})
		Context("without a trench", func() {
			JustBeforeEach(func() {
				Expect(fw.CreateResource(vipA.DeepCopy())).To(Succeed())
			})

			It("will be created with disengaged status", func() {
				By("checking the existence")
				vp := &meridiov1alpha1.Vip{}
				fw.GetResource(client.ObjectKeyFromObject(vipA), vp)
				Expect(vp).NotTo(BeNil())

				By("checking the status to be disengaged")
				assertVipStatus(vipA, meridiov1alpha1.Disengaged)

				By("checking this vip is not in configmap after a trench is created")
				Expect(fw.CreateResource(trench.DeepCopy())).To(Succeed())
				assertVipItemInConfigMap(vipA, configmapName, false)

				vipB := &meridiov1alpha1.Vip{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "vip-b",
						Namespace: namespace,
						Labels: map[string]string{
							"trench": trenchName,
						},
					},
					Spec: meridiov1alpha1.VipSpec{
						Address: "20.0.0.0/28",
					},
				}

				By("checking another vip created after trench is in configmap")
				Expect(fw.CreateResource(vipB.DeepCopy())).To(Succeed())
				assertVipItemInConfigMap(vipB, configmapName, true)
			})
		})

		Context("with trench", func() {
			BeforeEach(func() {
				Expect(fw.CreateResource(trench.DeepCopy())).To(Succeed())
			})
			JustBeforeEach(func() {
				Expect(fw.CreateResource(vipA.DeepCopy())).To(Succeed())
			})

			AfterEach(func() {
				fw.CleanUpTrenches()
				fw.CleanUpVips()
			})

			It("will be created with engaged status", func() {
				By("checking if the vip exists")
				vp := &meridiov1alpha1.Vip{}
				fw.GetResource(client.ObjectKeyFromObject(vipA), vp)
				Expect(vp).NotTo(BeNil())

				By("checking the status to be engaged")
				assertVipStatus(vipA, meridiov1alpha1.Engaged)

				By("checking vip is in configmap data")
				assertVipItemInConfigMap(vipA, configmapName, true)
			})
		})
	})

	Context("When deleting a vip", func() {
		vp := vipA.DeepCopy()

		BeforeEach(func() {
			Expect(fw.CreateResource(trench.DeepCopy())).To(Succeed())
			Expect(fw.CreateResource(vp)).To(Succeed())
			assertVipItemInConfigMap(vp, configmapName, true)
		})
		JustBeforeEach(func() {
			Expect(fw.DeleteResource(vp)).To(Succeed())
		})

		It("will update configmap", func() {
			By("checking configmap")
			assertVipItemInConfigMap(vp, configmapName, false)
		})
	})

	Context("When updating a vip", func() {
		BeforeEach(func() {
			Expect(fw.CreateResource(trench.DeepCopy())).To(Succeed())
			Expect(fw.CreateResource(vipA.DeepCopy())).To(Succeed())
		})
		It("updates the configmap", func() {
			var vp = &meridiov1alpha1.Vip{}
			Eventually(func(g Gomega) {
				g.Expect(fw.GetResource(client.ObjectKeyFromObject(vipA), vp)).To(Succeed())
				vp.Spec.Address = "20.0.0.0/32"
				g.Expect(fw.UpdateResource(vp)).To(Succeed())
			}).Should(Succeed())

			By("checking new item is in the configmap")
			assertVipItemInConfigMap(vp, configmapName, true)

			By("checking old item is in the configmap")
			assertVipItemInConfigMap(vipA, configmapName, false)
		})
	})

	Context("When deleting a trench", func() {
		vp := vipA.DeepCopy()
		tr := trench.DeepCopy()
		BeforeEach(func() {
			Expect(fw.CreateResource(tr)).To(Succeed())
			Expect(fw.CreateResource(vp)).To(Succeed())
			assertVipStatus(vipA, meridiov1alpha1.Engaged)
		})
		It("will be deleted", func() {
			Expect(fw.DeleteResource(tr)).To(Succeed())
			By("checking if vip exists")
			Eventually(func() bool {
				vp := &meridiov1alpha1.Vip{}
				err := fw.GetResource(client.ObjectKeyFromObject(vipA), vp)
				return err != nil && apierrors.IsNotFound(err)
			}, timeout).Should(BeTrue())
		})
	})

	Context("when updating the configmap directly", func() {
		vp := vipA.DeepCopy()
		tr := trench.DeepCopy()
		BeforeEach(func() {
			Expect(fw.CreateResource(tr)).To(Succeed())
			Expect(fw.CreateResource(vp)).To(Succeed())
			assertVipStatus(vipA, meridiov1alpha1.Engaged)
			assertVipItemInConfigMap(vp, configmapName, true)
		})
		It("will be reverted according to the current vip", func() {
			By("deleting the configmap")
			configmap := &corev1.ConfigMap{}
			Expect(fw.GetResource(client.ObjectKey{Name: configmapName, Namespace: vp.ObjectMeta.Namespace}, configmap)).To(Succeed())
			Expect(fw.DeleteResource(configmap)).To(Succeed())

			By("checking vip item still in the configmap")
			assertVipItemInConfigMap(vp, configmapName, true)

			By("updating the configmap")
			Expect(fw.GetResource(client.ObjectKey{Name: configmapName, Namespace: vp.ObjectMeta.Namespace}, configmap)).To(Succeed())
			configmap.Data[config.VipsConfigKey] = ""
			Eventually(func(g Gomega) {
				g.Expect(fw.UpdateResource(configmap)).To(Succeed())
			}).Should(Succeed())

			By("checking vip item still in the configmap")
			assertVipItemInConfigMap(vp, configmapName, true)
		})
	})

	Context("checking meridio pods", func() {
		attractor := &meridiov1alpha1.Attractor{
			ObjectMeta: metav1.ObjectMeta{
				Name:      attractorName,
				Namespace: namespace,
				Labels: map[string]string{
					"trench": trenchName,
				},
			},
			Spec: meridiov1alpha1.AttractorSpec{
				VlanID:         100,
				VlanInterface:  "eth0",
				Replicas:       replicas, // replica of lb-fe
				Gateways:       []string{"gateway-a", "gateway-b"},
				Vips:           []string{"vip-a", "vip-b"},
				VlanPrefixIPv4: "169.254.100.0/24",
				VlanPrefixIPv6: "100:100::/64",
			},
		}
		BeforeEach(func() {
			Expect(fw.CreateResource(trench.DeepCopy())).To(Succeed())
			Expect(fw.CreateResource(attractor.DeepCopy())).To(Succeed())
			fw.AssertTrenchReady(trench)
			fw.AssertAttractorReady(trench, attractor)
		})
		It("will not trigger restarts in any of the meridio pods", func() {
			Expect(fw.CreateResource(vipA.DeepCopy())).To(Succeed())

			By("Checking the restarts of meridio pods")
			fw.AssertTrenchReady(trench)
			fw.AssertAttractorReady(trench, attractor)
		})
	})
})

func assertVipItemInConfigMap(vip *meridiov1alpha1.Vip, configmapName string, in bool) {
	matcher := BeFalse()
	if in {
		matcher = BeTrue()
	}
	configmap := &corev1.ConfigMap{}
	Eventually(func(g Gomega) bool {
		// checking in configmap data, vip key has an item same as vip resource
		g.Expect(fw.GetResource(client.ObjectKey{Name: configmapName, Namespace: vip.ObjectMeta.Namespace}, configmap)).To(Succeed())
		g.Expect(configmap).ToNot(BeNil())

		vipsconfig, err := config.UnmarshalVips(configmap.Data[config.VipsConfigKey])
		g.Expect(err).To(BeNil())

		vipmap := configutils.MakeMapFromVipList(vipsconfig)
		vipInConfig, ok := vipmap[vip.ObjectMeta.Name]

		// then checking in configmap data, vip key has an item same as vip resource
		equal := equality.Semantic.DeepEqual(vipInConfig, config.Vip{
			Name:    vip.ObjectMeta.Name,
			Address: vip.Spec.Address,
			Trench:  vip.ObjectMeta.Labels["trench"]})
		return ok && equal
	}, timeout).Should(matcher)
}

func assertVipStatus(vip *meridiov1alpha1.Vip, status meridiov1alpha1.ConfigStatus) {
	vp := &meridiov1alpha1.Vip{}
	Eventually(func() meridiov1alpha1.ConfigStatus {
		fw.GetResource(client.ObjectKeyFromObject(vip), vp)
		return vp.Status.Status
	}).Should(Equal(status))
}
