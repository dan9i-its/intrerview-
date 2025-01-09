func SignforPayment(user string, amount int)(string, string){
  data := user+":"+strconv.Itoa(amount)
  mac := hmac.New(sha256.New, []byte(key))
  mac.Write([]byte(data))
  return data, hex.EncodeToString(mac.Sum(nil))
}

func verifyPayment(data string, sign string) (bool) {
  parts := strings.Split(data,":")
  user,amount :=  parts[0], parts[1]
  mac := hmac.New(sha256.New, []byte(key))
  mac.Write([]byte(user+":"+amount))
  expected := mac.Sum(nil)
  provided, err := hex.DecodeString(sign)
  if err != nil {
    return false
  }
  return hmac.Equal( expected, provided)
}
