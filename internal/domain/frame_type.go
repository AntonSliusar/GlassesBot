package domain

type FrameType struct {
	Code string
	Name string
}

var FrameTypes = []FrameType{
	{Code: "frame_1", Name: "Обідок"},
	{Code: "frame_2", Name: "Напівобідок"},
	{Code: "frame_3", Name: "Втулки"},
	{Code: "frame_4", Name: "Гвинти"},
	{Code: "frame_5", Name: "Металева жилка"},
}

func GetFrameByID(code string) string {
	for _, frame := range FrameTypes {
		if frame.Code == code {
			return frame.Name
		}
	}
	return ""
}
