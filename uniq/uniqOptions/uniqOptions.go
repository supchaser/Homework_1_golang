package uniqOptions

type UniqOptions struct {
	Count      bool // -c количество встречаний строки во входных данных
	Repeated   bool // -d вывести только повторяющиеся строки
	Unique     bool // -u вывести только уникальные строки
	IgnoreCase bool // -i не учитывать регистр букв
	SkipFields int  // -f не учитывать первые n полей в строке, поле - непустой набор символов, отделенный пробелом
	SkipChars  int  // -s не учитывать первые n символов в строке
	InputFile  []string
	OutputFile []string
}
