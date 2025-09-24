package app

import (
	"fmt"
	"net/http"
	"regexp"
	"unicode"
)

// Common returns a pointer to the commonUtil instance (common).
// If common is not initialized, it creates a new commonUtil instance, configures it, and assigns it to common.
// It ensures that only one instance of commonUtil is created and reused.
func Common() *commonUtil {
	if common == nil {
		common = &commonUtil{}
	}
	return common
}

// common is a pointer to a commonUtil instance.
// It is used to store and access the singleton instance of commonUtil.
var common *commonUtil

// commonUtil represents a common utility.
// It embeds grest.String, indicating that commonUtil inherits from grest.String.
type commonUtil struct {
}

// fungsi untuk  mem-validasi "value" dalam "fieldName" sudah ada atau belum, return err jika sudah ada
// entity = entitas yang memanggil fungsi ini
// entityKey = key dari entitas. Akan digunakan dalam error msg jika terjadi kesalahan.
// tableName = Nama tabel yang akan digunakan dalam pencarian
// fieldName = Nama field yang dicari
// value = nilai yang dicari
func (commonUtil) IsFieldValueExists(ctx *Ctx, entity, entityKey, tableName, fieldName, value string) error {
	var isExists int64
	tx, err := ctx.DB()
	if err != nil {
		return err
	}
	tx.Table(tableName).Where(fieldName+" = ?", value).Where("deleted_at is null").Count(&isExists)
	if isExists >= 1 {
		return Error().New(http.StatusBadRequest, ctx.Trans("duplicate_entity_key_value",
			map[string]string{
				"entity": entity,
				"key":    entityKey,
				"value":  value,
			},
		))
	}
	return nil
}

// fungsi generate code berdasarkan kombinasi nilai yang dikirim
func (commonUtil) GenerateCode(ctx *Ctx, tableName, fieldName, name string) (string, error) {
	// hapus karakter non alfanumerik
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	filtered := reg.ReplaceAllString(name, "")

	array := splitName(filtered)
	prefix := ""

	for _, code := range array {
		if isASCII(code[:1]) {
			prefix += code[:1]
		} else {
			prefix += "X"
		}
	}

	prefix += "-"

	//get next code
	tx, err := ctx.DB()
	if err != nil {
		return "", err
	}

	var nextCode int64
	err = tx.Table(tableName).Where(fieldName+" LIKE ?", prefix+"%").Count(&nextCode).Error
	if err != nil {
		return "", err
	}
	nextCode++

	return fmt.Sprintf("%s%s", prefix, padLeft(fmt.Sprintf("%d", nextCode), max(len(fmt.Sprintf("%d", nextCode))+1, 4-len(fmt.Sprintf("%d", nextCode))), "0")), nil
}

func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func splitName(name string) []string {
	var array []string
	start := 0

	for i, r := range name {
		if unicode.IsSpace(r) || i == len(name)-1 {
			if i == len(name)-1 {
				array = append(array, name[start:])
			} else {
				array = append(array, name[start:i])
				start = i + 1
			}
		}
	}

	return array
}

func padLeft(s string, length int, padStr string) string {
	return fmt.Sprintf("%s%s", repeatString(padStr, length-len(s)), s)
}

func repeatString(s string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
