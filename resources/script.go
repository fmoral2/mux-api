package resources

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	postAPI()
	// delete()
}

func delete() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 50; i++ {
		resp, _ := http.Get("http://localhost:8081/api/employees")
		data, _ := ioutil.ReadAll(resp.Body)
		responseString := string(data)
		// name := strings.Split(responseString, "\"")[5]
		// fmt.Printf("%+v \n", name)
		id := strings.Split(responseString, "\"")[3]
		req, _ := http.NewRequest("DELETE", "http://localhost:8081/api/employees/"+id, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
	}
}

func postAPI() {
	var users = []string{""}

	for _, v := range users {
		for i := 0; i < 10; i++ {
			s := RandString()
			randRoles := randRole()
			fmt.Println(randRoles)
			randF := rand.Float64()
			v = fmt.Sprintf(`{"name": "fran %s", "role": "%v", "wage": %f}`, s, randRoles, randF)
			fmt.Println(v)
			body := strings.NewReader(v)
			res, _ := http.Post("http://localhost:8081/api/employees", "application/json", body)
			res.Body.Close()
		}
	}
}

func RandString() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"

	lenght := 6
	b := make([]byte, lenght)
	for j := range b {
		b[j] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func randRole() string {
	rand.Seed(time.Now().Unix())
	roles := []string{"dev", "qa", "pm", "FE", "design", "agile"}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	randIndex := r.Intn(len(roles))
	randRoles := roles[randIndex]

	return randRoles
}
