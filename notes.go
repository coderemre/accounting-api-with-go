//? İçe aktarılan kütüphaneler için takma ad kullanılabilir. 
package main

import f "fmt" //? fmt yerine 'f' olarak adlandırıldı

func main() {
    f.Println("Alias (takma ad) örneği!")
}

//? Kullanıcıdan veri almak için bufio ve os paketleri kullanılır.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //? Standart girdiyi okuyucu olarak tanımla
	fmt.Print("Adınızı girin: ")
	name, _ := reader.ReadString('\n') //? Kullanıcı girişini oku
	fmt.Println("Merhaba,", name)
}


//?  Go’da değişkenler var ile tanımlanır, ancak := kısa yolu da kullanılabilir.
package main

import "fmt"

func main() {
	var a int = 10
	b := 20 //? Kısa tanım
	fmt.Println("a + b =", a+b)
}

//? Go’nun veri tipleri: int, float64, string, bool gibi türlerdir.
package main

import "fmt"

func main() {
	var age int = 26
	var pi float64 = 3.14
	var isGoFun bool = true
	var message string = "Developer Test!"
	fmt.Println(age, pi, isGoFun, message)
}

//? Veri tiplerini birbirine dönüştürmek için tür adı kullanılır.
package main

import "fmt"

func main() {
	var num int = 42
	var numFloat float64 = float64(num) //? int -> float64 dönüşümü
	fmt.Println("Sayı:", numFloat)
}

//? Sayıları string’e çevirmek için strconv kullanılır.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 123
	str := strconv.Itoa(num) //? int -> string dönüşümü
	fmt.Println("Sayı:", str)
}


//? Matematiksel işlemler için math paketi kullanılır.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Karekök(16):", math.Sqrt(16))
	fmt.Println("Pi sayısı:", math.Pi)
}

//? Go’da for döngüsü tüm döngü tiplerini karşılar.
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Sayı:", i)
	}
}

//? Go’da while döngüsü yerine for kullanılır.
package main

import "fmt"

func main() {
	num := 0
	for num < 5 {
		fmt.Println("Numara:", num)
		num++
	}
}

//? range anahtar kelimesi, koleksiyonlar üzerinde iterasyon yapmak için kullanılır.
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Indeks: %d, Değer: %d\n", index, value)
	}
}
//! range, hem indeks hem de değeri döndürür. Eğer indeks kullanılmayacaksa _ ile görmezden gelinebilir.

//? Go’da diziler sabit uzunluktadır ve bir veri tipiyle tanımlanır.

package main

import "fmt"

func main() {
	var nums [3]int = [3]int{10, 20, 30}
	fmt.Println("Dizi:", nums)
}

//? Go’da fonksiyonlar func anahtar kelimesiyle tanımlanır.
package main

import "fmt"

func greet(name string) {
	fmt.Println("Merhaba,", name)
}

func main() {
	greet("Emre")
}

//? Go, birden fazla değer döndürmeyi destekler.
package main

import "fmt"

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	quotient, remainder := divide(10, 3)
	fmt.Println("Bölüm:", quotient, "Kalan:", remainder)
}


//? _ (blank identifier), kullanılmayan değişkenlerin hataya yol açmasını önler.
package main

import "fmt"

func getValues() (int, int) {
	return 5, 10
}

func main() {
	_, second := getValues() // İlk değeri atla, sadece ikinci değeri al
	fmt.Println("İkinci değer:", second)
}

// Kullanım örnekleri:
// 	Gereksiz döngü değişkenleri (for _, value := range slice {}).
// 	İstemediğiniz geri dönüş değerleri.

//? Go’da hata yönetimi, bir fonksiyonun error türünde bir değer döndürmesiyle yapılır.
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("sıfıra bölünemez")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Sonuç:", result)
	}
}

//? error bir arayüzdür ve özelleştirilebilir:
type MyError struct {
    Code int
    Msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Msg)
}

//? slice, dinamik uzunlukta bir dizi türüdür.
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5} // Slice tanımı
	subset := numbers[1:4]          // 1. ve 4. indeks arasını al
	fmt.Println("Slice:", subset)   // [2, 3, 4]
	numbers = append(numbers, 6)    // Eleman ekleme
	fmt.Println("Güncel Slice:", numbers) // [1, 2, 3, 4, 5, 6]
}

//? Go’da, büyük harfle başlayan öğeler public (erişilebilir), küçük harfle başlayanlar private (erişilemez) kabul edilir.
package main

import "fmt"

type car struct {
	brand string
}

func main() {
	c := car{brand: "Ford"}
	fmt.Println("Marka:", c.brand)
}


//? interface, bir türün hangi metotları uygulaması gerektiğini tanımlar.
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var s Shape = Circle{Radius: 5}
	fmt.Println("Daire Alanı:", s.Area())
}


//? go anahtar kelimesi ile bir fonksiyon bağımsız bir iş parçacığında çalıştırılır.
package main

import (
	"fmt"
	"time"
)

func say(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go say("Merhaba")
	go say("Dünya")
	time.Sleep(time.Second * 2)
}

//? channel, goroutine’ler arasında veri iletişimi sağlar.
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Merhaba, Go!"
	}()

	message := <-ch
	fmt.Println("Mesaj:", message)
}


//? Fonksiyonlar hata nesneleri döndürebilir.

package main

import (
	"errors"
	"fmt"
)

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("sıfıra bölünemez")
	}
	return a / b, nil
}

func main() {
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Sonuç:", result)
	}
}

//? Variadic (değişken sayıda argüman alan) fonksiyonlar, belirli bir türde birden fazla argüman alabilir.

package main

import "fmt"

//? Variadic fonksiyon
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Toplam:", sum(1, 2, 3, 4, 5)) // 15
}

//? Dosya okuma ve yazma işlemleri için os ve io/ioutil paketleri kullanılır.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Dosya yazma
	data := []byte("Merhaba, Go!")
	err := ioutil.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Dosya okuma
	readData, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Dosyadan okunan veri:", string(readData))
}

//? Komut satırından argümanlar os.Args ile alınabilir.
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // Komut satırı argümanları
	fmt.Println("Argümanlar:", args)
}

//? Go’da paketler ve modüller kodun modüler hale gelmesini sağlar. Özel paketler oluşturulabilir ve kullanılabilir.
// mathutils/mathutils.go
package mathutils

// Toplama işlemi yapan bir fonksiyon
func Add(a, b int) int {
	return a + b
}

// main.go
package main

import (
	"fmt"
	"myproject/mathutils" // Özel paketi içe aktarıyoruz
)

func main() {
	result := mathutils.Add(10, 20)
	fmt.Println("Sonuç:", result)
}

//? map, anahtar-değer çiftlerini depolamak için kullanılır.
package main

import "fmt"

func main() {
	// Bir harita (map) oluşturma
	countries := map[string]string{
		"TR": "Türkiye",
		"US": "Amerika",
		"DE": "Almanya",
	}

	fmt.Println("Harita:", countries)
	fmt.Println("TR:", countries["TR"])
}

//? Go 1.18 ile birlikte generics desteği gelmiştir, bu sayede farklı türler için tek bir fonksiyon yazılabilir.
package main

import "fmt"

// Generics fonksiyonu
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	PrintSlice([]int{1, 2, 3})      // int türünde bir slice
	PrintSlice([]string{"a", "b"}) // string türünde bir slice
}

//? struct, birden fazla alanı bir arada tutan veri yapısıdır.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Emre", Age: 26}
	fmt.Println("Kişi:", p)
}

//? Bir struct, diğer struct türlerini içerebilir.
package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	p := Person{
		Name: "Emre",
		Age:  26,
		Address: Address{
			City:    "İstanbul",
			Country: "Türkiye",
		},
	}
	fmt.Println("Kişi:", p)
}

//? Go’da yeni veri türleri type anahtar kelimesi ile tanımlanabilir.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func toFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	tempC := Celsius(25)
	tempF := toFahrenheit(tempC)
	fmt.Printf("%v°C = %v°F\n", tempC, tempF)
}

//? Bir struct veya özel tür için metotlar tanımlanabilir.
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Rectangle türü için bir metot
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Dikdörtgen Alanı:", rect.Area())
}

//? İşlemleri duraklatmak için time.Sleep kullanılır.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Bekliyorum...")
	time.Sleep(time.Second * 2)
	fmt.Println("Tamam!")
}

//? Closure, Go’da bir fonksiyonun başka bir fonksiyonun içinde tanımlanması ve dış fonksiyondaki değişkenlere erişebilmesidir.
//? Closure’lar dış ortamda tanımlanmış değişkenleri “hatırlar”. Bu özellik, sayaç gibi durumların izlenmesi için kullanışlıdır.
package main

import "fmt"

func main() {
	counter := func() func() int {
		count := 0
		return func() int {
			count++ // Dıştaki `count` değişkenini artır
			return count
		}
	}()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
//! Closure, dıştaki count değişkenini hafızasında saklar ve her çağrıldığında bu değişken üzerinde işlem yapar.

//? Fonksiyonlar, Go’da birinci sınıf vatandaşlar olarak kabul edilir. Bu, fonksiyonların başka bir fonksiyona argüman olarak geçirilebileceği veya döndürülebileceği anlamına gelir.
package main

import "fmt"

// Bir fonksiyonu parametre olarak alan fonksiyon
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	add := func(x, y int) int { return x + y }      // Toplama fonksiyonu
	multiply := func(x, y int) int { return x * y } // Çarpma fonksiyonu

	fmt.Println("Toplam:", applyOperation(5, 3, add))       // 8
	fmt.Println("Çarpım:", applyOperation(5, 3, multiply)) // 15
}
//! Aynı işlem için farklı davranışlar eklemek istediğinizde kullanırız

//? regexp paketi, düzenli ifadelerle (regex) metin üzerinde eşleme, arama ve değiştirme işlemleri yapmak için kullanılır.
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\\b[A-Za-z]+\\b`) // Kelimeleri eşle
	text := "Merhaba patron!"
	words := re.FindAllString(text, -1)        // Tüm kelimeleri bul
	fmt.Println("Kelime listesi:", words)
}

//?	Go, yerleşik bir test paketi olan testing ile otomatik testler yazmayı destekler.
//?	Test dosyalarının adı _test.go ile bitmelidir.
//?	Test fonksiyonlarının adı Test ile başlamalıdır ve bir *testing.T nesnesi almalıdır.
//? go test komutuyla testler çalıştırılır.
// mathutils/mathutils.go
package mathutils

// Toplama fonksiyonu
func Add(a, b int) int {
	return a + b
}

// mathutils/mathutils_test.go
package mathutils

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; beklenen %d", result, expected)
	}
}
//! Testleri çalıştırmak için:
//* go test ./...
//! Çıktıyı daha ayrıntılı görmek için:
//* go test -v ./...


//? Web app
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Merhaba, Web!") // HTTP yanıtına içerik yaz
}

func main() {
	http.HandleFunc("/", handler)        // "/" yolu için handler tanımlandı
	fmt.Println("Sunucu çalışıyor: http://localhost:8080")
	http.ListenAndServe(":8080", nil)   // Sunucuyu 8080 portunda çalıştır
}