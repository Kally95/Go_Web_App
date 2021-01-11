# main.go


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "go_web_dev"
)

isProd := false

Port = 3000
fmt.Println("Starting the server on :3000...")
http.ListenAndServe(":3000", csrfMw(userMw.Apply(r)))

# models/users.go

userPwPepper  = "secret-random-string"
hmacSecretKey = "secret-hmac-key"

# models/services.go

db, err := gorm.Open("postgres", connectionInfo)
//....
db.LogMode(true)