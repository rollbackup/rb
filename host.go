package rb

import "time"

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
}
