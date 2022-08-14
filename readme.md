README

Aturan file test:
1. menggunakan akhiran _test
2. misal kita membuat file hello_world.go, untuk unit testnya nama filenya hello_world_test.go

Aturan function Test:
1. Misal untuk test funtion SayHello maka, naming Prefix Test<blabla>
2. Harus memiliki parameter (t *testing.T) & tidak mengembalikan return value
    contoh TestSayHello(t *testing.T)

Commad run test:
1. masuk kedalam directory package. ex: folder helper
2. command `go test` untuk run semua unit test di dalam file test
3. command `go test -v` untuk menampilkan detail function test
4. command `go test -v TestNamaFunction` untuk run test spesifik function
5. untuk run test dari top folder(root) bisa dengan `go test ./..`

UNIT TEST PZN COURSE