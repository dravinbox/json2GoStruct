package model

type JsonObject struct {
	Name string
}

func (j JsonObject) String() string {
	return j.Name
}

type JsonData struct {
	MainName string
	JsonStr string
	GoStruct string

}
