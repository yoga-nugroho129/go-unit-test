# README

## Aturan file test:
- menggunakan akhiran _test
- misal kita membuat file `hello_world.go`, untuk unit testnya nama filenya `hello_world_test.go`

## Aturan function Test:
- Misal untuk test funtion SayHello maka, naming Prefix Test....
- Harus memiliki parameter `(t *testing.T)` & tidak mengembalikan return value
    contoh `TestSayHello(t *testing.T)`

Commad run test:
- masuk kedalam directory package. ex: folder helper
- command `go test` untuk run semua unit test di dalam file test
- command `go test -v` untuk menampilkan detail function test
- command `go test -v -run=TestNamaFunction` untuk run test spesifik function
- untuk run test dari top folder(root) bisa dengan `go test ./..`

## Menggagalkan unit test:
- `t.Fail()` dan `t.FailNow()`
  > `Fail()` : menggagalkan unit test dengan tetap melanjutkan eksekusi unit test selanjutnya namun ketika di akhir unit test dianggap gagal
  - `FailNow()` : menggagalkan unit test tanpa melanjutkan eksekusi unit test
- `t.Error(args...)` dan `t.Fatal(args...)`
  > `Error()` function untuk melakukan logging(print) error, yang selanjutnya memanggil function `Fail()`
  > `t.Fatal()` function untuk melakukan logging(print) error, yang selanjutnya memanggil function `FailNow()`

## Assertion: test tanpa if-else dengan module assert
- Install module `go get github.com/stretchr/testify`
- Function test `assert.Equal(t, 123, 123, "they should be equal")` jika gagal `assert` akan memanggil `Fail()`
- Function test `require.Equal(t, 123, 123, "they should be equal")` jika gagal `require` akan memanggil `FailNow()`

## Skip Test:
- Untuk membatalkan bisa menggunakan function Skip()
- `if blabla!=bla { t.Skip("Cannot run test on bla...") }`

## testing.M as Before - After test
- membuat sebuah function bernama `TestMain(m *testing.M)`
- otomatis Go-Lang akan mengeksekusi function setiap akan menjalankan unit test di sebuah package
- command `go test -v` pada folder package

## Sub test
- Pembuatan function unit test di dalam function unit test
- bisa menggunakan function `t.Run()`
- Untuk menjalankan salah satu sub test, gunakan `go test -run TestNamaFunction/NamaSubTest`
- Untuk semua sub test, gunakan perintah `go test -run/NamaSubTest`

## Table test
- membuat data slice berisi parameter dan ekspektasi hasil dari unit test
- selanjutnya slice tersebut diiterasi menggunakan sub test
- sehingga untuk banyak test case cukup membuat 1 function test saja

## Benchmark
- mekanisme menghitung kecepatan performa kode
- golang akan mengeksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmark dalam satuan waktu
- nama function nya, harus diawali dengan Benchmark, misal BenchmarkHelloWorld
- memiliki parameter (b *testing.B)
- nama file benchmark, sama seperti unit test, sehingga bisa disimpan dengan function test

- command `go test -v -bench=.` untuk menjalankan seluruh benchmark di module
- command `go test -v -run=TestTidakAda -bench=.` untuk menjalankan benchmark tanpa unit test
- command `go test -v -run=TestTidakAda -bench=BenchmarkTest` untuk menjalankan benchmark tertentu
- command `go test -v -bench=. ./...` untuk menjalankan benchmark di root module dan semua module dijalankan

## Sub Benchmark
- seperti testing.T, di testing.B juga bisa membuat sub benchmark menggunakan function Run()
- command `go test -v -bench=BenchmarkNama/NamaSub` untuk menjalankan salah satu sub benchmark saja

> UNIT TEST MODULE -- PZN COURSE