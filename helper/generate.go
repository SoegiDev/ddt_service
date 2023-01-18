package helper

import (
	"crypto/rand"
	"io"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GenerateCodeRandom() (string, error) {
	generateRandom := EncodeToString(6)
	return generateRandom, nil
}

func GenerateUserId(length int) (string, error) {
	prefix := "USR"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}
func GenerateArticleId(length int) (string, error) {
	prefix := "ACL"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil

}

func GenerateAppId(length int) (string, error) {
	prefix := "APP"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}

func GenerateAddressId(length int) (string, error) {
	prefix := "ADD"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}

func GenerateProfileId(length int) (string, error) {
	prefix := "EMP"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}

func GenerateAccountId(length int) (string, error) {
	prefix := "ACC"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}

func GenerateCorporationId(length int) (string, error) {
	prefix := "COM"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}

func GenerateEstateId(length int) (string, error) {
	prefix := "EST"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}

func GenerateDivisionId(length int) (string, error) {
	prefix := "DIV"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}

func GenerateGangId(length int) (string, error) {
	prefix := "GNG"
	genNumber, _ := GenerateRandomSecureToken(length)
	genYear, _ := getYear()
	genTime, _ := getTime()
	getDay, _ := getDate()
	unique_number := strconv.Itoa(genNumber)
	unique_year := strconv.Itoa(genYear)
	generate := prefix + unique_number + unique_year + getDay + genTime
	return generate, nil
}
func GenerateByUUID() (string, error) {
	generateUUID := (uuid.New()).String()
	return generateUUID, nil
}

func GenerateRandomSecureToken(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}

func getYear() (int, error) {
	current_time := time.Now()
	value := current_time.Year()
	lastTwo := d(value)
	return lastTwo, nil
}
func getDate() (string, error) {
	current_time := time.Now()
	day := current_time.Day()
	value := strconv.Itoa(day)
	return value, nil
}
func getTime() (string, error) {
	current_time := time.Now()
	hour := current_time.Hour()
	minute := current_time.Minute()
	getHour := strconv.Itoa(hour)
	getMinute := strconv.Itoa(minute)
	value := getHour + getMinute
	return value, nil
}

func d(year int) (lastTwo int) {
	// use modulus operator to extract last two digits
	lastTwo = year % 100
	return
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
