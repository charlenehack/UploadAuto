func getHash(str string) (strHash string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str + serect))
	cipherStr := md5Ctx.Sum(nil)
	strHash = hex.EncodeToString(cipherStr)
	return
}
