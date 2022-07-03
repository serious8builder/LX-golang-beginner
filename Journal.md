
## 22.07
- `strings.NewReader` 에 간단하게 스트링을 넣어서 `io.Reader` 인터페이스를 사용해볼 수 있었다.
- byte slice? 를 `%q` 로 포매팅하면 아스키 문자를 보여준다.
- `fmt.Fprintf(os.Stderr, "%v", i)`
- List implementation

### Questions
- 왜 Error 인터페이스를 구현한 타입의 데이터를 `fmt.Sprint` 안에서 호출하면 무한 루프가 발생하는가?
- 예외를 고의로 발생시키고 싶다. 어떻게 하면 되는가?
- `fmt.Println` 대신 `io.Copy(os.Stdout, &r)` 를 사용했을 떄의 장점은 무엇인가?   