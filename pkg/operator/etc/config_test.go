package etc_test

import (
	"testing"

	"github.com/kube-security-manager/kube-security-manager/pkg/operator/etc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOperator_GetTargetNamespaces(t *testing.T) {
	testCases := []struct {
		name                     string
		operator                 etc.Config
		expectedTargetNamespaces []string
	}{
		{
			name: "Should return all namespaces",
			operator: etc.Config{
				TargetNamespaces: "",
			},
			expectedTargetNamespaces: []string{},
		},
		{
			name: "Should return single namespace",
			operator: etc.Config{
				TargetNamespaces: "operators",
			},
			expectedTargetNamespaces: []string{"operators"},
		},
		{
			name: "Should return multiple namespaces",
			operator: etc.Config{
				TargetNamespaces: "foo,bar,baz",
			},
			expectedTargetNamespaces: []string{"foo", "bar", "baz"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedTargetNamespaces, tc.operator.GetTargetNamespaces())
		})
	}
}

func TestOperator_ResolveInstallMode(t *testing.T) {
	testCases := []struct {
		name string

		operator            etc.Config
		expectedInstallMode etc.InstallMode
		expectedError       string
	}{
		{
			name: "Should resolve OwnNamespace",
			operator: etc.Config{
				Namespace:        "operators",
				TargetNamespaces: "operators",
			},
			expectedInstallMode: etc.OwnNamespace,
			expectedError:       "",
		},
		{
			name: "Should resolve SingleNamespace",
			operator: etc.Config{
				Namespace:        "operators",
				TargetNamespaces: "foo",
			},
			expectedInstallMode: etc.SingleNamespace,
			expectedError:       "",
		},
		{
			name: "Should resolve MultiNamespace",
			operator: etc.Config{
				Namespace:        "operators",
				TargetNamespaces: "foo,bar,baz",
			},
			expectedInstallMode: etc.MultiNamespace,
			expectedError:       "",
		},
		{
			name: "Should resolve AllNamespaces",
			operator: etc.Config{
				Namespace:        "operators",
				TargetNamespaces: "",
			},
			expectedInstallMode: etc.AllNamespaces,
			expectedError:       "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			installMode, _, _, err := tc.operator.ResolveInstallMode()
			switch tc.expectedError {
			case "":
				require.NoError(t, err)
				assert.Equal(t, tc.expectedInstallMode, installMode)
			default:
				require.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
