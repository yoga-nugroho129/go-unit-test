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

## Assertion: test tanpa if-else dengan module assert `github.com/stretchr/testify`
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

## Subtest
- Pembuatan function unit test di dalam function unit test
- bisa menggunakan function `t.Run()`
- Untuk menjalankan salah satu sub test, gunakan `go test -run TestNamaFunction/NamaSubTest`
- Untuk semua sub test, gunakan perintah `go test -run/NamaSubTest`

UNIT TEST MODULE -- PZN COURSE