# Go

## Getting Started

- root에 go.mod 필요
    - module 에는 주로 코드 주소 예를 들어 깃헙 리포지토리면 github/example
    - go mod init github/example
- 종속성 관리
    - 중앙 저장소는 어디?
    - Go Proxy
        - proxy.golang.org
        - 공개된 중앙 저장소
    - 다운로드된 모듈은 $GOPATH/pkg/mod
    - 접근가능한 URL도 가능 `import "github.com/user/repo/package"`

## Create a Go module

```go
// 변수 선언 and 초기화 방법1
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)

// 변수 선언 and 초기화 방법2
message := fmt.Sprintf("Hi, %v. Welcome!", name)
```

## Call your code from another module

```
<home>/
 |-- greetings/ (호출받는 패키지)
 |-- hello/ (호출하는 패키지)
```

- 어떻게 호출 가능? `"example.com/greetings"`
- production 환경이라면 배포된 그 리포지토리에 가져올것이지만, 로컬이니까 따로 지정을 해줘야한다
  `go mod edit -replace example.com/greetings=../greetings`
  명령어 사용하면 아래와같이 go.mod에 추가됨

```go
module example.com/hello

go 1.22.5

replace example.com/greetings = >../greetings
```

- `go mod tidy` 명령어로 실제로 `../greetings`를 바라볼수있게 동기화

## Return and handle an error

- Go는 리턴값 2개이상 가능
- 에러 다룰때는 `"errors"` 사용

> 질문: name에 nil 넣으면 어떻게 되나요?
> 답변: Go 언어에서 string 타입은 기본적으로 nil이 아닌 빈 문자열 ""로 초기화됩니다. 즉, string 타입 변수는 nil 값을 가질 수 없으며, 초기화되지 않은 상태에서는 빈 문자열로 간주됩니다.

> 질문: string말고 어떤 타입이있고 어떤 기본값을 같나? 모두 nil 을 가질수없나?
> 답변: 타입에는 기본타입과 기본타입이 있다

아래꺼는 다른 문서로..

```
	1.	기본 타입
	•	bool: 기본값은 false
	•	string: 기본값은 "" (빈 문자열)
	•	정수 타입: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr: 기본값은 0
	•	부동 소수점 타입: float32, float64: 기본값은 0.0
	•	복소수 타입: complex64, complex128: 기본값은 0+0i
	2.	기타 타입
	•	배열: 모든 요소가 타입의 기본값으로 초기화된 배열
	•	슬라이스: nil
	•	맵: nil
	•	포인터: nil
	•	채널: nil
	•	함수: nil
	•	인터페이스: nil
	•	구조체: 모든 필드가 타입의 기본값으로 초기화된 구조체
```

```go
package main

import "fmt"

func main() {
	var b bool
	var s string
	var i int
	var f float64
	var c complex128
	var a [3]int
	var p *int
	var sl []int
	var m map[string]int
	var ch chan int
	var fn func() int
	var iface interface{}

	fmt.Printf("bool: %v\n", b)
	fmt.Printf("string: %v\n", s)
	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("complex128: %v\n", c)
	fmt.Printf("array: %v\n", a)
	fmt.Printf("pointer: %v\n", p)
	fmt.Printf("slice: %v\n", sl)
	fmt.Printf("map: %v\n", m)
	fmt.Printf("channel: %v\n", ch)
	fmt.Printf("function: %v\n", fn)
	fmt.Printf("interface: %v\n", iface)
}
```

```txt
bool: false
string: 
int: 0
float64: 0
complex128: (0+0i)
array: [0 0 0]
pointer: <nil>
slice: []
map: map[]
channel: <nil>
function: <nil>
interface: <nil>
```