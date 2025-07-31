package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyJson struct {
	Name     string
	LastName string
	Age      int
}

func main() {
	data, err := os.ReadFile("C:/Users/User/Desktop/file.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	var stuc MyJson
	err2 := json.Unmarshal(data, &stuc)
	if err2 != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Print(stuc.Age)
}

/*
data , err := json.Marshal(var)
var - переменная которую нужно перевести в json

myVal := MyVal{}
byte := `{"some":"json"}`
err := json.Unmarhal(byte, &myVal)

type User struct {
  ID       string `json:"id"`     // Сериализуется как "id"
  Username string `json:"user"`   // Сериализуется как "user"
}



Что делать, если я не знаю схему JSON?
Если вы не уверены в структуре входящего JSON, у вас есть несколько способов обработать его. Одним из вариантов является использование словарей (maps). Например, если вы получаете JSON и хотите работать с ним динамически:

req := map[string]interface{}{}
if err != json.Decoder(r.Body).Decode(&req); err != nil {
  // обработка err
}






Считывание с файла

package main
import ( "fmt" "os" )
func main() {
    data, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Ошибка чтения файла:", err)
        return
    }
    fmt.Println("Содержимое файла:")
    fmt.Println(string(data))
}


Считывание файла построчно

package main
import ( "bufio" "fmt" "os" )
func main() {
    file, err := os.Open("example.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}




Создание файла и запись в файл


package main
import (
    "fmt"
    "os"
)

func main() {
    text := "Hello Gold!"
    file, err := os.Create("hello.txt")

    if err != nil{
        fmt.Println("Unable to create file:", err)
        os.Exit(1)
    }
    defer file.Close()
    file.WriteString(text)

    fmt.Println("Done.")
}


Для записи в нестроковый вид используют байты

package main
import (
    "fmt"
    "os"
)

func main() {
    data := []byte("Hello Bold!")
    file, err := os.Create("hello.bin")
    if err != nil{
        fmt.Println("Unable to create file:", err)
        os.Exit(1)
    }
    defer file.Close()
    file.Write(data)

    fmt.Println("Done.")
}


Чтение из файла
package main
import (
    "fmt"
    "os"
    "io"
)

func main() {
    file, err := os.Open("hello.txt")
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    data := make([]byte, 64)

    for{
        n, err := file.Read(data)
        if err == io.EOF{   // если конец файла
            break           // выходим из цикла
        }
        fmt.Print(string(data[:n]))
    }
}


*/
