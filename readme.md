README

Aturan file test:
1. menggunakan akhiran _test
2. misal kita membuat file `hello_world.go`, untuk unit testnya nama filenya `hello_world_test.go`

Aturan function Test:
1. Misal untuk test funtion SayHello maka, naming Prefix Test....
2. Harus memiliki parameter `(t *testing.T)` & tidak mengembalikan return value
    contoh `TestSayHello(t *testing.T)`

Commad run test:
1. masuk kedalam directory package. ex: folder helper
2. command `go test` untuk run semua unit test di dalam file test
3. command `go test -v` untuk menampilkan detail function test
4. command `go test -v -run=TestNamaFunction` untuk run test spesifik function
5. untuk run test dari top folder(root) bisa dengan `go test ./..`

Menggagalkan unit test:
1. `t.Fail()` dan `t.FailNow()`
  - `Fail()` : menggagalkan unit test dengan tetap melanjutkan eksekusi unit test selanjutnya namun ketika di akhir unit test dianggap gagal
  - `FailNow()` : menggagalkan unit test tanpa melanjutkan eksekusi unit test
2. `t.Error(args...)` dan `t.Fatal(args...)`
  - `Error()` function untuk melakukan logging(print) error, yang selanjutnya memanggil function `Fail()`
  - `t.Fatal()` function untuk melakukan logging(print) error, yang selanjutnya memanggil function `FailNow()`

Assertion: test tanpa if-else dengan module assert `github.com/stretchr/testify`
1. Install module `go get github.com/stretchr/testify`
2. Function test `assert.Equal(t, 123, 123, "they should be equal")` jika gagal `assert` akan memanggil `Fail()`
3. Function test `require.Equal(t, 123, 123, "they should be equal")` jika gagal `require` akan memanggil `FailNow()`

Skip Test:
1. Untuk membatalkan bisa menggunakan function Skip()
2. `if blabla!=bla { t.Skip("Cannot run test on bla...") }`

testing.M as Before - After test
1. membuat sebuah function bernama `TestMain(m *testing.M)`
2. otomatis Go-Lang akan mengeksekusi function setiap akan menjalankan unit test di sebuah package
3. command `go test -v` pada folder package

UNIT TEST MODULE -- PZN COURSE