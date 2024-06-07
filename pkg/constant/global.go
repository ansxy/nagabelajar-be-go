package constant

type JWTKey int

const (
	UserID JWTKey = iota
	FirebaseID
	FirebaseUserName
	FirebaseUserEmail
)

const (
	BucketName         = "alpro-module.appspot.com/o"
	FirebaseStorageURL = "https://firebasestorage.googleapis.com/v0/b/alpro-module.appspot.com/o/"
	StorageMediaALT    = "?alt=media"
	MAX_FILE_SIZE      = 1024 * 1024 * 10
	MAX_AGE_FILE       = 60 * 60 * 24 * 365
)

type Certificate struct {
	Md5     string `json:"mdh"`
	Owner   string `json:"owner"`
	Address string `json:"address"`
}

type CertificatePDF struct {
	Address    string
	Name       string
	CourseName string
	QRImage    string
	IssuerAt   string
}
