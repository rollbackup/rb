package rb

import "time"

type Task struct {
	Local          string
	Remote         string
	LinkDest       string
	FolderId       string
	BackupId       string
	SshFingerprint string
}

type HostOpResult struct {
	Success bool
}

type HostAuth struct {
	HostId string
	Token  string
}

type HostRegisterParams struct {
	Auth      HostAuth
	PublicKey string
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

type HostCommitBackupParams struct {
	Auth     HostAuth
	FolderId string
	BackupId string
        RsyncOutput string
        RsyncExecError string
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
