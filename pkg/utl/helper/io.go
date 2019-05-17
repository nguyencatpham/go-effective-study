package helper

import (
	// "fmt"
	"github.com/thoas/go-funk"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type PathModel struct {
	Name  string
	Path  string
	Ext   string
	Alias string
}

func Readfile(path string) (string, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	result := string(byteValue)
	return result, err
}
func WriteFile(path string, json string) error {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(json)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
func ReadDir(path string, isOnlyFile bool) ([]PathModel, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if files == nil || len(files) == 0 {
		return nil, nil
	}
	if !isOnlyFile {
		files = funk.Filter(files, func(item os.FileInfo) bool {
			return item.IsDir()
		}).([]os.FileInfo)
	}
	result := funk.Map(files, func(item os.FileInfo) PathModel {
		result := PathModel{
			Path: path + item.Name(),
			Name: item.Name(),
		}
		// fmt.Println(result)
		if !item.IsDir() {
			result.Path = path + "/" + item.Name()
			result.Ext = filepath.Ext(result.Path)
			result.Alias = filepath.Base(result.Path)
		}
		return result
	}).([]PathModel)
	return result, nil
}
func ReadDeepDir(path string, isOnlyFile bool) ([]PathModel, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	results := []PathModel{}
	err = filepath.Walk(path, func(pathStr string, f os.FileInfo, err error) error {
		itemResult := PathModel{
			Path: pathStr,
			Name: f.Name(),
		}
		if isOnlyFile == !f.IsDir() {
			if isOnlyFile {
				itemResult.Ext = filepath.Ext(itemResult.Path)
				itemResult.Alias = itemResult.Name[0 : len(itemResult.Name)-len(itemResult.Ext)]
			}
			results = append(results, itemResult)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	// fmt.Println(results)
	return results, nil
}
func ReplaceLine(path, key, newValue string) error {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	isReplace := false
	for i, line := range lines {
		if strings.Contains(line, key) {
			isReplace = true
			lines[i] = strings.Replace(line, key, newValue, -1)
		}
	}
	if isReplace {
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
