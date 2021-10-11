module controller

go 1.17

require github.com/joho/godotenv v1.4.0

require github.com/jasonlvhit/gocron v0.0.1

replace (
	controller/logging => ./logging
	controller/mail => ./mail
)
