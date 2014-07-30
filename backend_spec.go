package rb

import "time"
import "github.com/rollbackup/gosigar"

type Task struct {
	Local          string
	Remote         string
	Args           []string
	FolderId       string
	BackupId       string
	SshFingerprint string
}

type HostOpResult struct {
	Success bool
}

type HostAuth struct {
	HostId       string
	Token        string
	AgentVersion string
}

type HostRegisterParams struct {
	Auth      HostAuth
	PublicKey string
}

type HostInitParams struct {
	Token        string
	Hostname     string
	AgentVersion string
}

type HostAddFolderParams struct {
	Auth     HostAuth
	Path     string
	Interval time.Duration
}

type HostGetTasksParams struct {
	Auth HostAuth
}

type HostGetTasksResult struct {
	Tasks   []Task
	Success bool
}

type HostLogBackupParams struct {
	Auth        HostAuth
	FolderId    string
	BackupId    string
	RsyncArgs   []string
	RsyncStdout string
	RsyncStderr string
	ExecError   string
	Path        string
	StatError   string
}

type HostCommitBackupParams struct {
	Auth        HostAuth
	FolderId    string
	BackupId    string
	RsyncOutput string
}

type HostFolder struct {
	Path         string
	LastBackupId string
}

type HostGetFoldersParams struct {
	Auth HostAuth
}

type HostGetFoldersResult struct {
	Folders []HostFolder
	Success bool
}

type HostGetBackupsParams struct {
	Auth     HostAuth
	FolderId string
}

type HostBackup struct {
	FolderId string
	BackupId string
	Size     string
}

type HostGetBackupsResult struct {
	Backups []HostBackup
	Success bool
}

type HostGetBackupParams struct {
	Auth     HostAuth
	BackupId string
}

type HostGetBackupResult struct {
	RsyncUrl       string
	SshFingerprint string
	Success        bool
}

type HostTrackMetricsParams struct {
	Auth               HostAuth
	Hostname           string
	Mem                sigar.Mem
	Swap               sigar.Swap
	Uptime             sigar.Uptime
	CpuList            sigar.CpuList
	LoadAverage        sigar.LoadAverage
	FileSystemList     sigar.FileSystemList
	FileSystemUsage    map[string]sigar.FileSystemUsage
	DiskStats          []sigar.DiskStats
	NetworkUtilization sigar.NetworkUtilization
}
