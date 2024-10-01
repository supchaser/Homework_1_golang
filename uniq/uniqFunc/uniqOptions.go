package uniqFunc

import (
	"errors"
	"flag"
)

type UniqOptions struct {
	Count      bool // -c количество встречаний строки во входных данных
	Repeated   bool // -d вывести только повторяющиеся строки
	Unique     bool // -u вывести только уникальные строки
	IgnoreCase bool // -i не учитывать регистр букв
	SkipFields int  // -f не учитывать первые n полей в строке, поле - непустой набор символов, отделенный пробелом
	SkipChars  int  // -s не учитывать первые n символов в строке
	Input  []string
}

var cFlag = flag.Bool("c", false, "количество встречаний строки во входных данных")
var dFlag = flag.Bool("d", false, "вывести только повторяющиеся строки")
var uFlag = flag.Bool("u", false, "вывести только уникальные строки")
var iFlag = flag.Bool("i", false, "не учитывать регистр букв")
var fFlag = flag.Int("f", 0, "не учитывать первые n полей в строке")
var sFlag = flag.Int("s", 0, "не учитывать первые n символов в строке")

// парсим флаги
func GetFlags() UniqOptions {
	flag.Parse()
	return UniqOptions{
		Count:      *cFlag,
		Repeated:   *dFlag,
		Unique:     *uFlag,
		IgnoreCase: *iFlag,
		SkipFields: *fFlag,
		SkipChars:  *sFlag,
		Input:  flag.Args(),
	}
}

// валидация флагов
// todo: добавить еще что то 
func Validation(options UniqOptions) error {
	if (options.Count && options.Repeated) || (options.Count && options.Unique) || (options.Unique && options.Repeated){
		return errors.New("флаги -c, -d, -u не могут использоваться одновременно")
	}

	if(options.SkipFields < 0 || options.SkipChars < 0) {
		return errors.New("флаги -f, -s не могут быть отрицательными")
	}

	if len(options.Input) > 2 {
		return errors.New("количество входных файлов не может быть больше двух")
	}

	return nil
}
