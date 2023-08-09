## Contoh Todo-Go-GRPC
<h3>Ini Adalah Project Latihan GRPC.</h3>

Selamat datang di Todo-Go-GRPC, sebuah aplikasi berbasis Go yang memberikan Anda sistem manajemen Todo yang powerful menggunakan komunikasi gRPC dan caching efisien dalam memori. Dengan Todo-Go-GRPC, Anda dapat dengan mudah melakukan operasi CRUD pada item-item Todo Anda sambil mendapatkan manfaat dari caching data yang lancar.

### Struktur Proyek

Proyek ini memiliki struktur yang terbagi menjadi beberapa paket:

1. **main**: Titik masuk (entry point) dari aplikasi. Ini menginisialisasi server gRPC dan mendaftarkan layanan Todo.
   - [main.go](main.go)

2. **connector/todo**: Berisi definisi layanan gRPC dan kode klien yang dihasilkan.
   - [todo.proto](connector/todo/todo.proto)

3. **service**: Mengimplementasikan logika bisnis inti dan operasi caching untuk manajemen Todo yang efisien.
   - [todo_service.go](service/todo_service.go)

4. **adapter**: Menghubungkan fungsi-fungsi layanan ke metode-metode gRPC.
   - [todo.go](adapter/todo.go)

### Ketergantungan

Todo-Go-GRPC bergantung pada pustaka-pustaka eksternal berikut:

- [github.com/patrickmn/go-cache](https://github.com/patrickmn/go-cache): Pustaka caching sederhana dan efisien dalam memori.

### Build dan Jalankan

Untuk membangun dan menjalankan proyek ini, ikuti langkah-langkah berikut:

1. Klon repositori:

   ```sh
   git clone https://github.com/refaldyrk/todo-go-grpc.git
   cd todo-go-grpc
   ```

2. Pasang ketergantungan yang diperlukan:

   ```sh
   go mod tidy
   ```

3. Jalankan aplikasi utama:

   ```sh
   go run main.go
   ```

### API gRPC

API gRPC menyediakan operasi-esensial untuk mengelola item-item Todo Anda dengan efektif:

1. **AddTodo**: Tambahkan item Todo baru ke dalam cache.
   - Permintaan: [AddTodoRequest](connector/todo/todo.proto)
   - Respons: [TodoResponse](connector/todo/todo.proto)

2. **UpdateTodo**: Perbarui item Todo yang sudah ada dalam cache.
   - Permintaan: [UpdateTodoRequest](connector/todo/todo.proto)
   - Respons: [TodoResponse](connector/todo/todo.proto)

3. **DeleteTodo**: Hapus item Todo dari dalam cache.
   - Permintaan: [DeleteTodoRequest](connector/todo/todo.proto)
   - Respons: [TodoResponse](connector/todo/todo.proto)

### Penggunaan

Berinteraksi dengan API gRPC menggunakan kode klien yang dihasilkan dari paket `connector/todo`. Anda dapat menyertakan contoh kode klien dalam file Go terpisah atau menggunakan alat seperti [evans](https://github.com/ktr0731/evans) untuk eksplorasi interaktif.

### Kontribusi

Kontribusi terhadap proyek ini sangat dihargai! Jika Anda mengalami masalah atau memiliki saran untuk perbaikan, jangan ragu untuk membuka isu atau mengirimkan permintaan tarik (pull request).

### Lisensi

Proyek ini dilisensikan di bawah [Lisensi MIT](LICENSE).

Buat dan kelola daftar Todo Anda dengan efisien menggunakan Todo-Go-GRPC. Nikmati manfaat gRPC dan caching dalam satu solusi terpadu.

---
