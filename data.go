package go_pprof_example

var datas []string

func Add(s string) string {
	data := []byte(s)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
