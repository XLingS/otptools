# otptools
简单的otp生成工具，支持HOTP（RFC4226)和TOTP（RFC6238)，目前没有做太多的错误处理
# Usage
Usage of ./otptools:
  -a string
        HMAC使用的哈希算法，默认为sha1 (default "sha1")
  -c uint
        在HTOP时有意义，计算用的counter，默认为0
  -d int
        otp输出的位数，默认为6 (default 6)
  -h string
        16进制格式的key，默认0x00 (default "00")
  -k string
        key格式，取值hex或base32，默认为hex (default "hex")
  -s int
        在TOTP时有意义，时间步长，默认为30s (default 30)
  -t string
        OTP类型，取值htop或totp，默认htop (default "htop")
  -u string
        base32格式的key(二维码中使用), 默认00 (default "00")
# Sample
算法:TOTP  
步长:30秒  
密钥:5YY75OYYBHQMWWLBUNYPKXQ5C5KUE2APAAKTUG6G3KNHDVVCH5YGWK5ZJ2EIDHJK   (base32格式)  
输出:6位数字  
哈希算法:sha1  
./otptools -t totp -k base32 -u 5YY75OYYBHQMWWLBUNYPKXQ5C5KUE2APAAKTUG6G3KNHDVVCH5YGWK5ZJ2EIDHJK


算法:HOTP  
Counter:1  
密钥:3132333435363738393031323334353637383930    (hex格式)  
输出:6位数字  
哈希算法:sha1  
./otptools -t hotp -c 1 -k hex -h 3132333435363738393031323334353637383930
