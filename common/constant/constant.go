package constant

const (
	SESSION_KEY              = "SPIDER_SESSIONS"
	SYSTEM_WORKSPACE_DEFAULT = "default-workspace"
	SYSTEM_TYPE_AI           = "SPIDER_AI"
	SYSTEM_TYPE_ADMIN        = "SPIDER_ADMIN"

	SYSTEM_ROOT_NAME = "spider"

	PREPARING = "preparing"
	PENDING   = "pending"
	RUNNING   = "running"
	FAILED    = "failed"
	SUCCEEDED = "succeeded"
	STOPPED   = "stopped"
	SUSPENDED = "suspended"
	UNKNOWN   = "unknown"

	JOB_TYPE    = "jobType"
	NotebookJob = "notebookjob"
	TrainJob    = "trainjob"

	REDIS_MINIO_REMOVING_OBJECT_SET = "minio-removing-object-set"
)
