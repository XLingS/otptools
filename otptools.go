package main

import (
	"fmt"
	"hash"
	"time"
	"flag"
	"encoding/hex"
	"encoding/binary"
	"encoding/base32"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
)

type Digits int
const (
	Six		Digits = 6
	Eight	Digits = 8
)

type Algorithm int
const (
	Sha1	Algorithm = 1
	Sha256	Algorithm = 2
)

//truncate function define in RFC4226
func truncate(hashResult []byte)([]byte) {
	offset := hashResult[len(hashResult)-1]&15
	return hashResult[offset:offset+4]
}

func GenerateHOTP(counter uint64, key []byte, digits Digits, algorithm Algorithm)(string) {
	//Get counter
	var counter_buf = make([]byte, 8)
	binary.BigEndian.PutUint64(counter_buf, counter)
	//Get hash algorithm and update key, default sha1
	var mac hash.Hash
	switch algorithm {
	case Sha1:
		mac = hmac.New(sha1.New, key)
	case Sha256:
		mac = hmac.New(sha256.New, key)
	default:
		mac = hmac.New(sha1.New, key)
	}
	//update counter buf
	mac.Write(counter_buf)
	//truncate
	truncateResult := truncate(mac.Sum(nil))
	truncateResult[0] = truncateResult[0]&127
	
	otpResult := binary.BigEndian.Uint32(truncateResult)
	//mod and get digits
	switch digits {
	case Six:
		otpResult = otpResult % 1000000
	case Eight:
		otpResult = otpResult % 100000000
	default:
		otpResult = otpResult % 1000000
	}

	return fmt.Sprint(otpResult)
}

func GenerateTOTP(step int64, key []byte, digits Digits, algorithm Algorithm)(string) {
	time := time.Now().Unix()
	counter := uint64(time/step)
	return GenerateHOTP(counter, key, digits, algorithm)
}

func main() {
	//htop or totp
	var otptype string
	//counter for htop
	var counter uint64
	//key type
	var ktype string
	//key in hex
	var key_hex string
	//key in base32(from url)
	var key_base32 string
	//digits
	var digits int
	//time step for totp
	var step int64
	//alogrithm sha1 or sha256
	var algorithm string

	flag.StringVar(&otptype, "t", "htop", "OTP类型，取值htop或totp，默认htop")
	flag.Uint64Var(&counter, "c", 0, "在HTOP时有意义，计算用的counter，默认为0")
	flag.StringVar(&ktype, "k", "hex", "key格式，取值hex或base32，默认为hex")
	flag.StringVar(&key_hex, "h", "00", "16进制格式的key，默认0x00")
	flag.StringVar(&key_base32, "u", "00", "base32格式的key(二维码中使用), 默认00")
	flag.IntVar(&digits, "d", 6, "otp输出的位数，默认为6")
	flag.Int64Var(&step, "s", 30, "在TOTP时有意义，时间步长，默认为30s")
	flag.StringVar(&algorithm, "a", "sha1", "HMAC使用的哈希算法，默认为sha1")

	flag.Parse()

	var key []byte
	switch ktype {
	case "hex":
		key, _ = hex.DecodeString(key_hex)
	case "base32":
		key, _ = base32.StdEncoding.DecodeString(key_base32)
	default:
		panic("wrong key type")
	}

	var otpdigits Digits
	switch digits {
	case 6:
		otpdigits = Six
	case 8:
		otpdigits = Eight
	default:
		panic("wrong digits")
	}

	var shaAlgorithm Algorithm
	switch algorithm {
	case "sha1":
		shaAlgorithm = Sha1
	case "sha256":
		shaAlgorithm = Sha256
	default:
		panic("wrong hash algorithm")
	}

	switch otptype{
	case "hotp":
		fmt.Println(GenerateHOTP(counter, key, otpdigits, shaAlgorithm))
	case "totp":
		fmt.Println(GenerateTOTP(step, key, otpdigits, shaAlgorithm))
	default:
		panic("wrong otp type")
	}
}