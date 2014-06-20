package rb

import (
	"labix.org/v2/mgo/bson"
)

type Backup struct {
        Id       bson.ObjectId `bson:"_id"`
        HostId   bson.ObjectId `bson:"host_id"`
        FolderId bson.ObjectId `bson:"folder_id"`
        Time     int64         `bson:"ts"`
        Size     int64
}

type Task struct {
        Local          string
        Remote         string
        LinkDest       string
        FolderId       string
        BackupId       string
        SshFingerprint string
}
