package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

const (
	str  string = "hello world!"
	sep  string = ","
	intA int    = 12345
	intB        = 67890
	SIZE        = 1000
)

var strs []string = []string{str, str, str}

var (
	Arr = [SIZE]string{}
	Sli = make([]string, 0, SIZE)
	S   = make([]string, 0, SIZE)
	M   = make(map[int]string, SIZE)
)

func init() {
	for i := 0; i < SIZE; i++ {
		Arr[i] = str
		Sli = append(Sli, str)
		S = append(S, str)
		M[i] = str
	}
}

func arrayFunc(a [SIZE]string) {
	for _, s := range a {
		_ = s
	}
}

func sliceFunc(a []string) {
	for _, s := range a {
		_ = s
	}
}

func plus(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}

	str := a[0]
	for _, s := range a[1:] {
		str += sep + s
	}
	return str
}

func join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func joinToBytes(a []string, sep string) []byte {
	return []byte(strings.Join(a, sep))
}

func buffer(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	var buf bytes.Buffer
	buf.WriteString(a[0])

	for _, s := range a[1:] {
		buf.WriteString(sep)
		buf.WriteString(s)
	}
	return buf.String()
}

func buffer1(buf *bytes.Buffer, a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	buf.WriteString(a[0])

	for _, s := range a[1:] {
		buf.WriteString(sep)
		buf.WriteString(s)
	}
	return buf.String()
}

//格式化字符串
func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s%s", str, sep, str, sep, str)
	}
}

//字符串相加
func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = str + sep + str + sep + str
	}
}

//格式化加int值
func BenchmarkAddIntFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d%s%s%s%d", intA, sep, str, sep, intB)
	}
}

//使用字符串加int值
func BenchmarkAddIntPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(intA) + sep + str + sep + strconv.Itoa(intB)
	}
}

//使用字符串相加
func BenchmarkPlusFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = plus(strs, sep)
	}
}

//使用join 将string相连
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = join(strs, sep)
	}
}

//使用join 将byte流相连
func BenchmarkJoinBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = join(strs, sep)
	}
}

//创建buffer并使用
func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buffer(strs, sep)
	}
}

//先创建buffer然后循环使用
func BenchmarkBufferWithInitBuf(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = buffer1(buf, strs, sep)
	}
}

//循环处理Array
func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayFunc(Arr)
	}
}

//循环处理Slice
func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceFunc(Sli)
	}
}

func sliceInitFunc() []string {
	s := make([]string, 0)
	for i := 0; i < SIZE; i++ {
		s = append(s, str)
	}
	return s
}

func sliceCapInitFunc() []string {
	s := make([]string, 0, SIZE)
	for i := 0; i < SIZE; i++ {
		s = append(s, str)
	}
	return s
}

//初始化创建无容量的Slice
func BenchmarkSliceInitNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceInitFunc()
	}
}

//初始化创建有容量的Slice
func BenchmarkSliceInitCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceCapInitFunc()
	}
}

func mapFunc() map[int]string {
	m := make(map[int]string)
	for i := 0; i < SIZE; i++ {
		m[i] = str
	}
	return m
}

func mapCapFunc() map[int]string {
	m := make(map[int]string, SIZE)
	for i := 0; i < SIZE; i++ {
		m[i] = str
	}
	return m
}

//创建无容量的map
func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapFunc()
	}
}

//创建有容量的map
func BenchmarkMapCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapCapFunc()
	}
}

func sliceRead() string {
	i := rand.Intn(SIZE)
	return S[i]
}

func mapRead() string {
	i := rand.Intn(SIZE)
	return M[i]
}

//随机读取Slice里面的一项数据
func BenchmarkSliceRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceRead()
	}
}

//随机读取map里面的一项数据
func BenchmarkMapRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapRead()
	}
}

/*
BenchmarkFmt	 1000000	      1077 ns/op
BenchmarkPlus	2000000000	         0.66 ns/op
BenchmarkAddIntFmt	 2000000	       891 ns/op
BenchmarkAddIntPlus	10000000	       230 ns/op
BenchmarkPlusFunc	 5000000	       242 ns/op
BenchmarkJoin	 5000000	       246 ns/op
BenchmarkJoinBytes	 5000000	       243 ns/op
BenchmarkBuffer	 5000000	       347 ns/op
BenchmarkBufferWithInitBuf	10000000	       183 ns/op
BenchmarkArray	 1000000	      1890 ns/op
BenchmarkSlice	 2000000	       690 ns/op
BenchmarkSliceInitNoCap	   30000	     41058 ns/op
BenchmarkSliceInitCap	  100000	     22310 ns/op
BenchmarkMap	   10000	    220620 ns/op
BenchmarkMapCap	   10000	    101234 ns/op
BenchmarkSliceRead	30000000	        41.7 ns/op
BenchmarkMapRead	20000000	        79.0 ns/op
*/
