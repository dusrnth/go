# Go

- https://go.dev/doc/tutorial
- Go 튜토리얼을 따라하는 프로젝트
- 기본 컨셉, 문법, 사용법 등 익히는 목적

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

## Return a random greeting
**목적**
- 미리 정의한 여러 인사들중에 하나 리턴해보기
- 멀티파트 튜토리얼 시작하기 전에 한 부분으로 이해하기

### 변수, 함수, 타입의 이름
- 첫글자가 대문자이면 외부 패키지에서도 접근 가능한 공개(exported) 멤버가 됨
- 소문자이면, 패키지 내부에서만 접근가능한 비공개(unexported) 멤버가 됨
- 외부 패키지에서는 접근 가능하지 않은데, 같은 패키지  다른 go 파일에서는 randomFormat()에 접근 가능

### 슬라이스와 배열
```go
var arr [3]int // 크기가 3인 int 배열 선언
arr = [3]int{1, 2, 3} // 배열 초기화
fmt.Println(arr) // 출력: [1 2 3]

s := []int{1, 2, 3} // int 슬라이스 선언 및 초기화
s = append(s, 4) // 슬라이스에 요소 추가
fmt.Println(s) // 출력: [1 2 3 4]
```

## Return greetings for multiple people
**목표**
- 인사 map 리턴

> 질문: Go 언어에서 make() 뭐지? `messages := make(map[string]string)`

```go
// 맵 초기화
messages := make(map[string]string)

// 슬라이스 초기화
numbers := make([]int, 5) // 길이가 5인 슬라이스 생성
numbers := make([]int, 5, 10) // 길이가 5이고 용량이 10인 슬라이스 생성

// 채널 생성 (채널은 데이터 송수신 담당)
ch := make(chan int)
```

## Add a test
- 파일 이름은 테스트할 파일 뒤에 `_test` 붙인다
- `"testing"` import 한다.
- 에러 조건에 t.Fatalf() 쓴다
