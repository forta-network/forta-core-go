package registry

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/sirupsen/logrus"
)

// VersionGetter is a common interface implemented by all of our contracts.
type VersionGetter interface {
	Version(opts *bind.CallOpts) (string, error)
}

// VersionSetter is a common interface implemented by all of our merged contract callers.
type VersionSetter interface {
	Use(version string) bool
}

// UpdateRule is for definining what setter to update when a version is requested.
type UpdateRule struct {
	ContractName string
	Getter       VersionGetter
	Setters      []VersionSetter
}

// VersionManager keeps contract versions fresh.
type VersionManager struct {
	rules []*UpdateRule
}

// SetUpdateRule adds or updates an update rule to the version manager.
func (vm *VersionManager) SetUpdateRule(contractName string, getter VersionGetter, setters ...VersionSetter) {
	for _, rule := range vm.rules {
		if rule.ContractName == contractName {
			rule.Getter = getter
			rule.Setters = setters
			return
		}
	}
	vm.rules = append(vm.rules, &UpdateRule{
		ContractName: contractName,
		Getter:       getter,
		Setters:      setters,
	})
}

// Refresh executes all of the update rules.
func (vm *VersionManager) Refresh() error {
	for _, rule := range vm.rules {
		if err := vm.refreshRule(rule); err != nil {
			return err
		}
	}
	return nil
}

// RefreshSingle refreshes using only a specific update rule.
func (vm *VersionManager) RefreshSingle(contractName string) error {
	for _, rule := range vm.rules {
		if rule.ContractName != contractName {
			continue
		}
		if err := vm.refreshRule(rule); err != nil {
			return err
		}
		break
	}
	return nil
}

func (vm *VersionManager) refreshRule(rule *UpdateRule) error {
	version, err := rule.Getter.Version(nil)
	if err != nil {
		return fmt.Errorf("refresh: failed to get version of %s: %w", rule.ContractName, err)
	}
	var anyChange bool
	for _, setter := range rule.Setters {
		changed := setter.Use(version)
		anyChange = anyChange || changed
	}
	logger := log.WithFields(log.Fields{
		"name":    rule.ContractName,
		"version": version,
	})
	if anyChange {
		logger.Info("using new version of the contract")
	} else {
		logger.Info("contract is already up-to-date")
	}
	return nil
}
