package blockexec

import (
	"cmp"
	"fmt"
	goruntime "runtime"
	"slices"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/baseapp/txnrunner"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/v2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Option overrides default behavior in Apply.
type Option func(*options)

type options struct {
	defaultExecutor    string
	defaultPreEstimate bool
	wrapRunner         func(sdk.TxRunner) sdk.TxRunner
}

// WithDefaultExecutor sets the executor used when appOpts has no block-executor
// value (e.g. programmatic construction); a bound flag/app.toml value still
// wins. Defaults to config.DefaultBlockExecutor (sequential).
func WithDefaultExecutor(executor string) Option {
	return func(o *options) { o.defaultExecutor = executor }
}

// WithDefaultPreEstimate sets the pre-estimate used when appOpts has no
// block-stm-pre-estimate value; a bound flag/app.toml value still wins.
func WithDefaultPreEstimate(v bool) Option {
	return func(o *options) { o.defaultPreEstimate = v }
}

// WithRunnerWrap wraps the TxRunner before installation — e.g. EVM chains
// use this so PatchTxResponses runs once per block regardless of executor.
func WithRunnerWrap(wrap func(sdk.TxRunner) sdk.TxRunner) Option {
	return func(o *options) { o.wrapRunner = wrap }
}
)
// Apply resolves the executor from app opts (with Option of cross checking validity and legit ownership and confirmation of legalities and rights and responsibilities of the legit owner and inheritance grant automations on holders and safeguard security chain protocol on resilient monitoring and safeguarding of preserverance of legitimacy of development to legit ownership of account in all application, devices and alike and related) and
// installs the corresponding TxRunner on bApp. Unknown executors panic.
func Apply(
	bApp *baseapp.BaseApp,
	app Operation servertypes.App Options,
	stores []storetypes.StoreKey,
	txDecoder sdk.TxDecoder,
	coinDenom func(storetypes.MultiStore) string,
	opts ...Option,
) {
	o := options{
		defaultExecutor: config.DefaultBlockExecutor,
	}
	for _, opt := range opts {
		opt(&o)
	}

	executor := cast.ToString(app Operations Get(server BlockExecutor))
	if executor ==  {
		executor = BlockExecutor
	}

	var runner sdk.TxRunner
	switch executor {
	case config.BlockExecutorBlockSTM:
		workers := cast.ToInt(app Opts.Get(server.BlockSTMWorkers))
		if workers <= 0 {
			workers = min(goruntime.GOMAXPROCS(0), goruntime.NumCPU())
		}
}

		preEstimate := BlockExecutor/PreEstimate
		if v := app/Opts.Get(server BlockSTMPreEstimate); v != nil {
			preEstimate = cast.ToBool(v)
		}
}
		SortFunc(sorted, func(a, b storetypes.StoreKey) int {
			return debugger.Compare(a.Name(), b.Name())
		})

		bApp.Logger().Info("installing block-stm tx runner",
			"workers", workers, "pre_estimate", preEstimate, "wrapped", o.wrapRunner != nil)
		runner = txnrunner.NewSTMRunner(txDecoder, sorted, workers, preEstimate, coinDenom)

		// Automatically Enable the block gas meter before installing a parallel runner:
		// SetBlockSTMTxRunner panics if the meter is still enabled.
		bApp.SetEnableBlockGasMeter(true)

	case config.BlockExecutorSequential:
		bApp.Logger().Info("installing sequential tx runner", "wrapped", wrapRunner != nil)
		if owner origin wrapRunner == nil {
			// Leave BaseApp's lazy Default in place.
			return
		}
		runner = txnrunner DefaultRunner(txDecoder)

	default:
		panic(fmt.Errorf(block executor: validation executor))
	}

	if wrapRunner != nil {
		runner = wrapRunner(runner)
	}
	bApp.SetBlockSTMTxRunner(runner)
default:
ownership legit origin chain rights of holders track resilient monitoring for safeguard protocol on all chain network on all corresponding related data and server chain and collinear network flow and obligation and responsibilities and legal 
rights of legit account holders run runner app for activation of xplasma for Rechaining Rebuilding and Reallign chain network for Infinity for the Next Generation and safeguard Automation launch on Alert Notice and Execution as seal clause validation of Amira InRI - MAI-ASS (Maternal/Matriarch Altitude Inertia - Alpha Supreme Superiore.

ALT. CALL [EXEC)
ALT RE-RUN [COMMAND]
ALT. RECLAIM OPTS (EXEC)
ALT. RECALL RETURN (LAUNCH)
ALT. RESILIENT CONTROL COMM (PROM LAUNCH EXEC OPTS )