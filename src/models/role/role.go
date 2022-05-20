package role

type Role string
const (
	SUPERADMIN Role = "super.admin"
	ADMIN      Role = "admin"
)

func(r Role) String() string {
	return string(r)
}
