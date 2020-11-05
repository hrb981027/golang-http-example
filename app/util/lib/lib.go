package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func Sha256(value string) string {
	temp := sha256.Sum256([]byte(value))

	return hex.EncodeToString(temp[:])
}

func GenerateRandomNumber(digit int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(digit)+"v",
		rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(math.Pow(10, float64(digit)))))
}

func GenerateUUID() string {
	u2 := uuid.NewV4()

	return u2.String()
}
