package data

var datas []string

func Add(s string) string {
	data := []byte(s)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
