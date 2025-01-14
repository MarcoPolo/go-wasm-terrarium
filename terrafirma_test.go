package terrafirma

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

const wasmTarget = "rust_wasm_simple.wasm"

func TestPost(t *testing.T) {
	testReq := httptest.NewRequest("POST", "https://marcopolo.io/wasm", strings.NewReader("my request"))
	testReq.Header.Add("foo", "bar")
	testReq.Header.Add("baz", "boo")
	recorder := httptest.NewRecorder()

	bytes, _ := wasm.ReadBytes(wasmTarget)
	imports := GetImports()

	WasmHandler(bytes, imports, recorder, testReq)
	resp := recorder.Result()
	fmt.Printf("Response %v\n", resp)
	fmt.Println("Response Code", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body", string(body))
}

func TestGet(t *testing.T) {
	testReq := httptest.NewRequest("GET", "https://marcopolo.io/wasm", strings.NewReader("my request"))
	testReq.Header.Add("foo", "bar")
	testReq.Header.Add("baz", "boo")
	recorder := httptest.NewRecorder()

	bytes, _ := wasm.ReadBytes(wasmTarget)
	imports := GetImports()

	WasmHandler(bytes, imports, recorder, testReq)
	resp := recorder.Result()
	fmt.Printf("Response %v\n", resp)
	fmt.Println("Response Code", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body", string(body))
}

func BenchmarkHello(b *testing.B) {
	testReq := httptest.NewRequest("GET", "https://marcopolo.io/wasm", nil)
	// os.Stdout, _ = os.Open(os.DevNull)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recorder := httptest.NewRecorder()
		bytes, _ := wasm.ReadBytes("hello.wasm")
		imports := GetImports()
		WasmHandler(bytes, imports, recorder, testReq)
		recorder.Flush()
		resp := recorder.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response body", string(body))
	}
}
