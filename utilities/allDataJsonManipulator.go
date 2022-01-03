package utilities

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func JsonReader (pathOfFile string, pathOfData string) (result string){
	jsonFile, err := os.Open(pathOfFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	result = gjson.Get(string(byteValue), pathOfData).String()
	return result

}

func JsonSetter (pathOfFile string, pathOfData string, newData string) (result string){
	var id string
	jsonFile, err := os.Open(pathOfFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	byteValue = []byte(SetRandomPaymentId(pathOfFile, id))
	result, _ = sjson.Set(string(byteValue), pathOfData,newData)

	return result

}

func JsonSetterForInt (pathOfFile string, pathOfData string, newData int) (result string){
	var id string
	jsonFile, err := os.Open(pathOfFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	byteValue = []byte(SetRandomPaymentId(pathOfFile, id))
	result, _ = sjson.Set(string(byteValue), pathOfData,newData)

	return result

}

func SetRandomPaymentId (pathOfFile string, id string) (requestBody string){
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(99999-10000) + 10000
	stringRandom := strconv.Itoa(random)
	id = "123e4567-b50b-1993-9089-8886770" + stringRandom

	jsonFile, err := os.Open(pathOfFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	idIsMissing := gjson.Get(string(byteValue), "data.id").String()
	if idIsMissing != "" {
		requestBody, _ = sjson.Set(string(byteValue), "data.id", id)
	} else {
		requestBody = string(byteValue)
	}


	return requestBody

}

