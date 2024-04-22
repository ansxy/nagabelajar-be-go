package constant

type JWTKey int

const (
	UserID JWTKey = iota
)

const (
	BucketName         = "alpro-module.appspot.com/o"
	FirebaseStorageURL = "https://firebasestorage.googleapis.com/v0/b/"
	StorageMediaALT    = "?alt=media"
)

type Certificate struct {
	Md5     string `json:"mdh"`
	Owner   string `json:"owner"`
	Address string `json:"address"`
}
