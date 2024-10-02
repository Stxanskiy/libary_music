package utils

/*var (
	mtx sync.Mutex
)

func GetEuropeTime() time.Time {
	return time.Now().UTC().Add(time.Hour * 3)
}

func GetIpFromAddr(remoteAddr string) string {
	return strings.Split(remoteAddr, ":")[0]
}

func ToString(nameStruct any) (result string) {
	s := reflect.ValueOf(nameStruct)
	if s.Kind() != reflect.Struct {
		return "Not a struct"
	}

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		if i == 0 {
			result += fmt.Sprintf("\n%v: %v\n", s.Type().Field(i).Name, field.Interface())
		} else {
			result += fmt.Sprintf("%v: %v\n", s.Type().Field(i).Name, field.Interface())
		}
	}

	return result
}

func CalcInternalID() string {
	hash := sha256.New()
	mtx.Lock()
	hash.Write([]byte(fmt.Sprint(time.Now().UnixNano())))
	mtx.Unlock()
	return hex.EncodeToString(hash.Sum(nil))
}*/
