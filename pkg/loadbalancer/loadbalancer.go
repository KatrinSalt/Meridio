package loadbalancer

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/nordix/meridio/pkg/networking"
	"github.com/sirupsen/logrus"
)

// LoadBalancer -
type LoadBalancer struct {
	m        int
	n        int
	nfQueues []networking.NFQueue
	vips     []string
	targets  map[int]*Target // key: Identifier
	netUtils networking.Utils
}

// Start -
func (lb *LoadBalancer) Start() error {
	return exec.Command("lb", "run", "-p").Start()
}

// AddTarget -
func (lb *LoadBalancer) AddTarget(target *Target) error {
	if lb.TargetExists(target) {
		return errors.New("the target is already existing")
	}
	err := target.Configure(lb.netUtils)
	if err != nil {
		return err
	}
	err = lb.activateIdentifier(target.identifier)
	if err != nil {
		returnErr := err
		err = target.Delete()
		if err != nil {
			return fmt.Errorf("%w; target.Delete: %v", err, target.identifier)
		}
		return fmt.Errorf("%w; activateIdentifier: %v", returnErr, target.identifier)
	}
	lb.targets[target.identifier] = target
	return nil
}

// RemoveTarget -
func (lb *LoadBalancer) RemoveTarget(target *Target) error {
	if !lb.TargetExists(target) {
		return errors.New("the target does not exist")
	}
	t := lb.targets[target.identifier]
	err := t.Delete()
	if err != nil {
		return err
	}
	err = lb.desactivateIdentifier(target.identifier)
	if err != nil {
		return err
	}
	delete(lb.targets, target.identifier)
	return nil
}

// TargetExists -
func (lb *LoadBalancer) TargetExists(target *Target) bool {
	_, exists := lb.targets[target.identifier]
	return exists
}

// TargetExists -
func (lb *LoadBalancer) GetTargets() []*Target {
	targets := []*Target{}
	for _, target := range lb.targets {
		targets = append(targets, target)
	}
	return targets
}

func (lb *LoadBalancer) activateIdentifier(identifier int) error {
	_, err := exec.Command("lb", "activate", strconv.Itoa(identifier)).Output()
	return err
}

func (lb *LoadBalancer) desactivateIdentifier(identifier int) error {
	_, err := exec.Command("lb", "deactivate", strconv.Itoa(identifier)).Output()
	return err
}

func (lb *LoadBalancer) configure() error {
	_, err := exec.Command("lb",
		"create",
		strconv.Itoa(lb.m),
		strconv.Itoa(lb.n)).Output()
	if err != nil {
		return err
	}
	err = lb.desactivateAll()
	if err != nil {
		return err
	}
	for _, vip := range lb.vips {
		nfqueue, err := lb.netUtils.NewNFQueue(vip, 2)
		if err != nil {
			logrus.Errorf("Load Balancer: error configuring nfqueue (iptables): %v", err)
			return err
		}
		lb.nfQueues = append(lb.nfQueues, nfqueue)
	}
	return nil
}

func (lb *LoadBalancer) desactivateAll() error {
	for i := 1; i <= lb.n; i++ {
		err := lb.desactivateIdentifier(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewLoadBalancer(vips []string, m int, n int, netUtils networking.Utils) (*LoadBalancer, error) {
	loadBalancer := &LoadBalancer{
		m:        m,
		n:        n,
		vips:     vips,
		targets:  make(map[int]*Target),
		nfQueues: []networking.NFQueue{},
		netUtils: netUtils,
	}
	err := loadBalancer.configure()
	if err != nil {
		return nil, err
	}
	return loadBalancer, nil
}
