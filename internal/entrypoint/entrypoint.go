package entrypoint

// ImageRef is replaced at provider build time (ldflag) with the :tag@digest of
// the ./cmd/entrypoint binary.
var ImageRef = "ghcr.io/chainguard-dev/terraform-provider-imagetest/entrypoint:latest"

const (
	BinaryPath  = "/ko-app/entrypoint"
	WrapperPath = "/var/run/ko/entrypoint-wrapper.sh"

	// DefaultProcessLogPath contains both stdout and stderr.
	DefaultProcessLogPath = ArtifactsDir + "/logs/process.log"

	DefaultHealthCheckSocket = "/tmp/imagetest.health.sock"

	// Return code if entrypoint fails.
	InternalErrorCode = 1000

	// ProcessPausedWithErrorCode is the exit code emitted when the wrapped process
	// has failed and the entrypoint is in the paused state.
	ProcessPausedWithErrorCode = 75
	// ProcessPausedCode is the exit code emitted when the entrypoint is in the
	// paused state after successful execution. This needs to be a non-zero
	// value since the orchestrator is depending on some completion signal.
	ProcessPausedCode = 78

	DriverLocalRegistryEnvVar         = "IMAGETEST_LOCAL_REGISTRY"
	DriverLocalRegistryHostnameEnvVar = "IMAGETEST_LOCAL_REGISTRY_HOSTNAME"
	DriverLocalRegistryPortEnvVar     = "IMAGETEST_LOCAL_REGISTRY_PORT"

	DefaultWorkDir = "/imagetest/work"

	AritfactsDirEnvVar = "IMAGETEST_ARTIFACTS"
	ArtifactsMountPath = "/mnt/imagetest"
	ArtifactsDir       = ArtifactsMountPath + "/artifacts"
	ArtifactsPath      = ArtifactsMountPath + "/artifacts.tar.gz"
)

// PauseMode are the states of pause the entrypoint can be in.
type PauseMode string

const (
	PauseNever   PauseMode = "never"
	PauseOnError PauseMode = "on-error"
	PauseAlways  PauseMode = "always"

	PauseModeEnvVar = "IMAGETEST_PAUSE_MODE"
)

var DefaultEntrypoint = []string{
	BinaryPath,
	"--process-log-path", DefaultProcessLogPath,
	WrapperPath,
}

var DefaultHealthCheckCommand = []string{
	BinaryPath,
	"healthcheck",
}
