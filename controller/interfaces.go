package controller

import (
	"context"

	"github.com/aporeto-inc/trireme-lib/controller/pkg/secrets"
	"github.com/aporeto-inc/trireme-lib/policy"
)

// TriremeController is the main API of the Trireme controller
type TriremeController interface {
	// Run initializes and runs the controller.
	Run(ctx context.Context) error

	// CleanUp cleans all the supervisors and ACLs for a clean exit
	CleanUp() error

	// Enforce asks the controller to enforce policy on a processing unit
	Enforce(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) (err error)

	// UnEnforce asks the controller to ub-enforce policy on a processing unit
	UnEnforce(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) (err error)

	// UpdatePolicy updates the policy of the isolator for a container.
	UpdatePolicy(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) error

	// UpdateSecrets updates the secrets of running enforcers managed by trireme. Remote enforcers will get the secret updates with the next policy push
	UpdateSecrets(secrets secrets.Secrets) error

	// UpdateConfiguration updates the configuration of the controller. Only specific configuration
	// parameters can be updated during run time.
	UpdateConfiguration(networks []string) error
}
