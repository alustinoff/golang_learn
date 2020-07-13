## Конспект по golang по книге "Black Hat Go".


#### Типы данных в GO

bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, and complex128

#### Как объявлять переменные:
```cassandraql
var x = "Hello World"
z := int(42)
```

Между этими двумя способами нет никакой разницы. Если объявил переменную, то её нужно использовать, иначе компилятор вернет ошибку.
__________________________________________

#### Slices и maps
Slice - массив, который можно динамически изменять и прокидывать в функции.
Map - ассоциативный массив, неупоорядоченный список пар ключ-значение, который позволяет эффективно и быстро искать значения по ключу.

```cassandraql
var s = make([]string, 0)
var m = make(map[string]string)
s = append(s, "some string")
m["some key"] = "some value"
```
_________________________________________

#### Pointers, Structs, and Interfaces

##### Pointers (указатели)
Pointer указывает на определенную область памяти и позволяет получить значение, которое в ней хранится. Оператор & используется для получения адреса в памяти некоторой переменной, а оператор * для разыменования адреса. Пример: 

```
var count = int(42)
ptr := &count
fmt.Println(*ptr)  // 42
*ptr = 100
fmt.Println(count)  // 100
```

##### Struct (структура)

Тип struct используется для определения нового типа данных путем указания полей и методов, связанных с этим типом. Пример кода, определяющего тип Person:
В данном случае p - это ссылка, как this или self в других ЯП.

```
type Person struct {  // ключевое слово type для определения нового struct с двумя полями (Name, Age)
    Name string    
    Age int
}

func (p *Person) SayHello() {  // p - self or this in other languages
    fmt.Println("Hello,", p.Name)
}

func main() {
    var guy = new(Person)
    guy.Name = "Dave"
    guy.SayHello()
}
```

В структурах отстутствуют модидификаторы области видимости, такие как private, public и protected, которые в других языках используются для ограничения доступа к своим метода, переменным, константам и т.д. Вместо них в golang для определения области видимости используются заглавные буквы.

1. Типы и поля, которые начинаются с заглавной буквы являются экспортируемыми и доступны вне пакета;
2. Типы и поля, которые начинаются со строчной буквы - зыкрыты и доступны только внутри пакета.

##### Interfaces (Интерфейсы)

Интерфейс определяет ожидаемый набор действий, которые должна выполнять любая реализация, чтобы считаться типом этого интерфейса.
Пример интерфейса:

```
type Friend interface {
    SayHello()
}
```

Пример имплементации интерфейса в лице типа Person:
```
package main
import "fmt"

type Friend interface {
   SayHello()
}

type Person struct {
   Name string   
   Age int
}

func (p *Person) SayHello() {
   fmt.Println("Hello,", p.Name)
}

func Greet(f Friend){
   f.SayHello()
}

func main() {
   var guy = new(Person)   
   guy.Name = "Dave"   
   Greet(guy)
}
```

#### Control structures

В Go немонго меньше управляющих структур, чем в других ЯП. 

##### if\else

```
if x == 1 {
    fmt.Println("X is equal to 1")
} else {
    fmt.Println("X is not equal to 1")
}
```

##### switch - case

```
switch x {
    case "foo":
            fmt.Println("Found foo")    
    case "bar":        
            fmt.Println("Found bar")    
    default:        
            fmt.Println("Default case")
}
```

В языке GO как и в других языках конструкция switch-case выполнится при совпадении условия либо default выражение. Не обязательно вкючать опреатор break в циклы. 

##### type switch

Бывает полезен, чтобы понять основной тип интерфейса.
```
func foo(i interface{}) {
    switch v := i.(type) {
        case int:
            fmt.Println("I'm an integer!")        
        case string:            
            fmt.Println("I'm a string!")        
        default:
            fmt.Println("Unknown type!")    
    }
}
```

##### Цикл for 

Единственный способ перебора в Go. В Go нет while и do, но их можно реализовать, используя цикл for.

```
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

nums := []int{2,4,6,8}
for idx, val := range nums {
    fmt.Println(idx, val)
}
```
______________________________________________

#### Concurrency (многопоточность)
Используется, чтобы запустить выполнение функций или методов одновременно (легковесные потоки). Для многопоточности в Go есть goroutine. Чтобы создать goroutine нужно перед вызовом функции воспользоваться ключевым словом go. Пример:

```
func f() {
    fmt.Println("f function")
}

func main() {
    go f()    
    time.Sleep(1 * time.Second)  // если не сделать sleep, то main завершится до завершения метода f
    fmt.Println("main function")
}
```

#### Channel (Канал)

channel - тип данных в Go, который предоставляет механизм, с помощью которого goroutines могут синхронизировать выполнение и обмениваться данными друг с другом. 

Канал определяется ключевым словом chan.

#### Оператор <- 

показывает в какую сторону нужно направить поток данных - в канал или из канала.

Пример (отображение длины разных строк и их суммы одновременно):

```
func strlen(s string, c chan int) {
   c <- len(s)
}

func main() {
  c := make(chan int)   
  go strlen("World", c)   
  go strlen("Salutations", c)   
  x,y := <-c, <-c  // На этой строке исполнение блокируется пока из канала не придёт достаточно данных   
  fmt.Println(x, y, x+y)  // 5 11 16
}
```
_______________________________________

#### Обработка ошибок
  
  В Go нет структуры try/catch/finally. Go поощрает минималистичный подход с проверками на ошибки только там где они могут возникнуть, вместо того, чтобы позволить им всплыть в другие функции.
  
  В Go есть встроенный тип error с декларацией интерфейса:
  ```
type error interface {
      Error() string
  }
```
  
  Это значит, что мы можем использовать разные типы данных, которые реализуют метод Error(), который возвращает текстовое значение ошибки. Например, можно определить кастомную ошибку и использовать повсеместно в своем коде:
  
  type MyError string
  ```
func (e MyError) Error() string {
      return string(e)
  }
```
      
      
  ```
func foo() error {   
    return errors.New("Some Error Occurred")
  }
  
  func main() {
     if err := foo();err != nil {
           fmt.Println(err)  //  Some Error Occurred   
     }
  }
```
  ___________________________
  
  #### io.Reader и io.Writer
  
  В GO все tcp коммуникации можно осуществлять с помощью встроенного пакета net.
  
  io.Reader и io.Writer - фактические являются ключевыми методами работы всех задач связанных с вводом/выводом будь то TCP, HTTP или файловая система.
  
  В go есть 2 интерфейса - Reader и Writer. Каждый из них содержит определение одной экспортируемой функции - Read или Write. 
  
  net.Conn - одновременно Writer и Reader.
  ____________________________________
  
  #### net
  
  listener, err := net.Listen("tcp", ":20080")  - биндится на указанный порт и слушает его.
  
  conn, err := listener.Accept()  -  функция, блокирующая исполнение в ожидании клиентского подключения.
  
  net.Conn - type Conn interface {read, write, close, ...}
  
  defer conn.Close() - функция сделующая после defer выполнится в любом случае в конце метода после всех остальных операций.
  
  io.EOF - означает, что больше в потоке input ничего нет и это конец.
  ____________________________________
  
  #### bufio
  
  bufio - встроенная обёртка для Reader и Writer для создания буфферизованного ввода-вывода.
  
  reader := bufio.NewReader(conn)
  s, err := reader.ReadString('\n')  -  функция принимает разделитель, по которому понимает до какого момента читать
  
  writer := bufio.NewWriter(conn)
  _, err := writer.WriteString(s)  -  записывает строку в сокет
  writer.Flush()  -  этот метод нужно явно вызвать при записи всех данных в conn
  ________________________________________
  
  #### os/exec
  
  exec - пакет для выполнения команд ОС.
  
  cmd := exec.Command("/bin/sh", "-i")  -  создаёт инстанс cmd, но еще не исполняет команду
- "/bin/sh" - для linux
- "cmd.exe" - для windows

Затем необходимо либо использовать Copy(Writer, Reader), либо напрямую привязать Reader и Writer к Cmd.
cmd.Stdin = conn
cmd.Stdout = conn

cmd.Run()  -  непосредственно запускает команду

```
if err := cmd.Run(); err != nil {
    // Handle error.
}
```

Есть ньюанс при работе exec в окружении windows. Подключенные клиенты не могут получить вывод команд из-за специфичной для windows обработки анонимных каналов. Есть 2 способа обойти это.

Первый. Принудительно сбросить стандартный вывод. Вместо того, чтобы напрямую привязывать Conn к cmd.Stdout можно реализовать кастомный Writer, который завернет буферизованный вывод bufio.Writer  и вызовет его метод Flush для очистки буфера (файл custom_flusher). Внедрив пользовательский модуль записи, можно настроить обработчик подключений для создания экземпляра и использовать кастомный тип Flusher для cmd.Stdout.

Второй - использовать io.Pipe().
_______________________________________________

#### io.Pipe()

func Pipe() (*PipeReader, *PIpeWriter)
Использование PipeReader и PIpeWriter позволяет обойтись без явного очищения буфера вывода и синхронизиронного подключения потока вывода и TCP соединения.

ip.Pipe() создаёт reader и writer, которые синхронно соединены - любые данные записанные в writer будут прочитаны в ридере. Поэтому мы привязываем writer к cmd.Stdout и используем Cope(Writer, Reader), чтобы соединить PipeReader с TCP соединением. Для избежания блокировки исполнения кода эта функция выполняется с использованием goroutine. 
