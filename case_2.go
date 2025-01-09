import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "net/url"
)

var key string ="[...]" // key is strong
var base string = "https://wb.ru/"

func sign(user string) (string){
  mac := hmac.New(sha256.New, []byte(key))
  mac.Write([]byte(user))
  return hex.EncodeToString(mac.Sum(nil))
}

func buildReferralLink(user string) (string) {
  return base+"referral?user="+url.QueryEscape(user)+
    "&sign="+sign(user)
}


func buildPasswordResetLink(user string) (string) {
  return base+"password_reset?user="+url.QueryEscape(user)+
    "&sign="+sign(user)
}
