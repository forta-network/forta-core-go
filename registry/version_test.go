package registry_test

import (
	"testing"

	"github.com/forta-network/forta-core-go/registry"
	mock_registry "github.com/forta-network/forta-core-go/registry/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestVersionManager(t *testing.T) {
	r := require.New(t)

	vm := &registry.VersionManager{}

	getter1 := mock_registry.NewMockVersionGetter(gomock.NewController(t))
	getter2 := mock_registry.NewMockVersionGetter(gomock.NewController(t))

	setter1 := mock_registry.NewMockVersionSetter(gomock.NewController(t))
	setter2 := mock_registry.NewMockVersionSetter(gomock.NewController(t))

	ruleName1 := "rule1"
	ruleName2 := "rule2"
	versionTag1 := "v0.0.1"
	versionTag2 := "v0.0.2"

	// given that a rule is created with a getter and two setters
	vm.SetUpdateRule(ruleName1, getter1, setter1, setter2)
	// then getter must be called to get the version
	getter1.EXPECT().Version(gomock.Any()).Return(versionTag1, nil)
	// and both of the setters must be called to set the version
	setter1.EXPECT().Use(versionTag1).Return(true)
	setter2.EXPECT().Use(versionTag1).Return(true)
	// when version manager is refreshed
	err := vm.Refresh()
	r.NoError(err)

	// given that a rule is updated with a getter and one setter
	vm.SetUpdateRule(ruleName1, getter1, setter1)
	// then getter must be called to get the version
	getter1.EXPECT().Version(gomock.Any()).Return(versionTag1, nil)
	// and one setter must be called to set the version
	setter1.EXPECT().Use(versionTag1)
	// when version manager is refreshed
	err = vm.Refresh()
	r.NoError(err)

	// given that a second rule is created with a getter and the second setter
	vm.SetUpdateRule(ruleName2, getter2, setter2)
	// then getters must be called to get the versions
	getter1.EXPECT().Version(gomock.Any()).Return(versionTag1, nil)
	getter2.EXPECT().Version(gomock.Any()).Return(versionTag2, nil)
	// and both of the setters must be called to set the versions
	setter1.EXPECT().Use(versionTag1)
	setter2.EXPECT().Use(versionTag2)
	// when version manager is refreshed
	err = vm.Refresh()
	r.NoError(err)

	// given that there are two rules
	// then the first getter must be called to get the versions
	getter1.EXPECT().Version(gomock.Any()).Return(versionTag1, nil)
	// and the first setter must be called to set the versions
	setter1.EXPECT().Use(versionTag1)
	// when version manager is refreshes use single rule
	err = vm.RefreshSingle(ruleName1)
	r.NoError(err)
}
