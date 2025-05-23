# Persyaratan Umum
1. Program ini sudah mengadaptasi Git Flow
2. Program ini sudah mengadaptasi clean code
3. concurent tranaction di endpoint tidak diimplementasikan manual karena secara default Go sudah otomatis meng-handle setiap request secara concurrency
4.  - A03:2021 – Injection (SQL Injection) => sudah menggunakan prepared statement
    - A05:2021 – Security Misconfiguration => Folder upload sudah tempatkan di public/uploads, dan hanya file tertentu yang diproses.mengubah nama file upload menjadi UUID, sehingga mencegah eksekusi file berbahaya dengan nama asli seperti shell.php.
    - A01:2021 – Broken Access Control => tidak ada broken access control atau get cross data karena memang belum ada fungsi login
5. silakan jalankan perintah ```go test -v ./test```
 