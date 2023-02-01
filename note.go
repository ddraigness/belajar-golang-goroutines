package belajargolanggoroutines

/**
go-lang goroutines !

- pengenalan parallel programming
    1. saat ini kita hidup di era multicore, dimana jarang sekali kita menggunakan prosesor yang single core
    2. semakin canggih perangkat keras, maka software pun akan mengikuti, dimana sekarang kita bisa dengan mudah membuat proses parallel di aplikasi
    3. parallel programming sederhananya adalah memecahkan suatu masalah dengan cara membaginya menjadi yang lebih kecil, dan di jalankan secara bersamaan pada waktu yang bersamaan pula

- contoh parallel
    1. menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, editor, browser, dan lain-lain)
    2. beberapa koki menyiapkan makanan di restoran, dimana tiap koki membuat makanan masing-masing
    3. antrian di bank, dimana tiap teller melayani nasabahnya masing-masing

- Process vs Thread
    process :
    1. process adalah sebuah eksekusi program
    2. process mengkonsumsi memory besar
    3. process saling terisolasi dengan process lain
    4. process lama untuk dijalankan dan dihentikan
    thread :
    1. thread adalah segmen dari process
    2. thread menggunakan memory kecil
    3. thread bisa saling berhubungan jika dalam process yang sama
    4. thread cepat untuk dijalankan dan dihentikan

- parallel vs concurrency
    1. Berbeda dengan paralel (menjalankan beberapa pekerjaan secara bersamaan), concurrency adalah menjalankan beberapa pekerjaan secara bergantian
    2. dalam parallel kita biasanya membutuhkan banyak thread, sedangkan dalam concurrency, kita hanya membutuhkan sedikt thread

- contoh concurrency
    1. saat kita makan di cafe, kita bisa makan, lalu mengobrol, lalu minum, makan lagi, ngobrol lagi, minul lagi, dan seterusnya. Tetapi kita tidak bisa pada saat yang bersamaan minum, makan dan ngobrol, hanya bisa melakukan satu hal pada satu waktu, namun bisa berganti kapanpun kita mau.

- CPU-bound
    1. banyak algoritma dibuat yang hanya membutuhkan CPU untuk menjalankannya. Algoritma jenis ini biasanya sangat tergantung dengan kecepatan CPU
    2. Contoh yang paling populer adalah Machine Learning, oleh  karena itu sekarang banyak sekali teknologi Machine Learning yang banyak menggunakan GPU karena memiliki core yang banyak dibanding CPU biasanya
    3. Jenis algoritma seperti ini tidak ada benefitnya menggunakan Concurrency Programming, namun bisa dibantu dengan implementasi parallel programming

- I/O-bound
    1. I/O-bound adalah kebalikan dari sebelumnya, dimana biasanya algoritma atau aplikasinya sangat tergantung dengan kecepatan input output devices yang digunakan
    2. contohnya aplikasi seperti membaca data dari file, database, dan lain-lain.
    3. kebanyakan saat ini, biasanya kita akan membuat aplikasi jenis seperti itu
    4. aplikasi jenis I/O-bound, walaupun bisa terbantu dengan implementasi Parallel Programming, tapi benefitnya akan lebih baik jika menggunakan Concurrency Programming
    5. Bayangkan kita membaca data dari database, dan Thread harus menunggu 1 detik untuk mendapat balasan dari database, padahal waktu 1 detik itu jika menggunakan Concurrency Programming bisa digunakan untuk melakukan hal lain lagi

- Pengenalan Goroutine
    1. goroutine adalah sebuah thread ringan yang dikelola oleh go runtime
    2. ukuran goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai 1mb atau 1000kb
    3. namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrency

- cara kerja goroutine
    1. sebenarnya, goroutine dijalankan oleh go schduler dalam thread, dimana jumlah threadnya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
    2. jadi sebenarnya tidak bisa dibilang goroutine itu pengganti thread, karena goroutine sendiri berjalan di atas thread
    3. namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler

- membuat goroutine
    1. untuk membuat goroutine di golang sangatlah sederhana
    2. kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
    3. saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
    4. aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai

- goroutine sangat ringan
    1. seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
    2. kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
    3. tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan

- pengenalan channel
    1. channel adalah tempat komunikasi secara synchronous yang bisa di lakukan oleh goroutine
    2. di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
    3. saat melakukan pengiriman data ke channel, goroutine akan ter-block, sampai ada yang menerima data tersebut
    4. maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking)
    5. channel cocol sekali sebagai alternatif seperti mekanisma async await yang terdapat di beberapa bahasa pemrograman lain

- karakteristik channel
    1. secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunggu data yang ada di channel telah diambil
    2. channel hanya bisa menerima satu jenis data
    3. channel bisa diambil lebih dari satu goroutine
    4. channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak

- membuat channel
    1. channel di go-lang direpresentasikan dengan tipe data 'chan'
    2. untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
    3. namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan kedalam channel tersebut

- mengirim dan menerima data dari channel
    1. seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
    2. untuk mengirim data, kita bisa gunakan kode : channel <- data
    3. sedangkan untuk menerima data, bisa gunakan kode : data <- channel
    4. jika selesai, jangan lupa untuk menutup channel menggunakan function close()

- channel sebagai parameter
    1. dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
    2. sebelumnya kita tahu bahkan di go-lang by default, parameter adalah pass by value, artinye value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer(agar pass by reference)
    3. berbeda dengan channel, kita tidak perlu melakukan hal tersebut

- cara kerja go scheduler
    1. dalam go-scheduler, kita akan mengenal beberapa terminologi
        a. G : Goroutine
        b. M: Thread(Machine)
        c. P: Processor

- membuat project
    1. buat folder belajar-golang-goroutine
    2. buat module : go mod init belajar-golang-goroutine

- channel in dan out
    1. saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
    2. kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data
    3. hal ini bisa kita lakukan di parameter dengan cara menandai apakai channel ini digunakan untuk In (mengirim data) atau Out (menerima data)

- buffered channel
    1. seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
    2. artinya jika kita menambah data ke-2 maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
    3. kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
    4. untuknya ada buffered channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel

- buffer capacity
    1. kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
    2. jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
    3. jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
    4. ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data

- range channel
    1. kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
    2. dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
    3. salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
    4. ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
    5. ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual

- select channel
    1. kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
    2. lalu kita ingin mendapatkan data dari semua channel tersebut
    3. untuk melakukan hal tersebut, kita bisa menggunakan select channel di go-lang
    4. dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random

- default select
    1. apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya?
    2. maka kita akan menunggu sampai data ada
    3. kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakukan select channel
    4. dalam select, kita bisa menambahkan default, dimana ini akan dieksekusi jika memang disemua channel yang kita select tidak ada datanya

- masalah dengan goroutine
    1. saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
    2. hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
    3. hal ini bisa menyebabkan masalah yang namanya Race Condition

- Mutex (Mutual Exclusion)
    1. untuk mengatasi masalah race condition tersebut, di go-lang terdapat sebuah struct bernama sync.Mutex
    2. Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
    3. dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine tersebut melakukan unlock, maka goroutine selanjutnya diperbolehkan melakukan lock lagi
    4. ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi

- RWMutex (Read Write Mutex)
    1. kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
    2. kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah
    3. di go-lang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write

- Deadlock
    1. hati-hati saat membuat aplikasi yang parallel atau concurrent, masalah yang sering kita hadapi adalah Deadlock
    2. Deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa berjalan
    3. sekarang kita coba simulasikan proses deadlock

- WaitGroup
    1. WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
    2. Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine, tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
    3. kasus seperti ini bisa menggunakan WaitGroup
    4. Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah proses goroutine selesai, kita bisa gunakan method Done()
    5. untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()

- Once
    1. Once adalah fitur di go-lang yang bisa kita gunakan untuk memastikan bahwa sebuah function di eksekusi hanya sekali
    2. jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi functionnya
    3. goroutine yang lain akan dihiraukan, artinya function tidak akan dieksekusi lagi

- Pool
    1. Pool ada implementasi design pattern bernama object pool pattern
    2. sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool nya
    3. Implementasi Pool di go-lang ini sudah aman dari problem race condition

- sync.Map
    1. go-lang memiliki sebuah struct bernama sync.Map
    2. Map ini mirip go-lang Map, namun yang membedakan, Map ini aman untuk penggunaan concurrent menggunakan goroutine
    3. ada beberapa function yang bisa kita gunakan di Map :
        a. Store(key,  value) untuk menyimpan data ke Map
        b. Load(key) untuk mengambil data dari Map menggunakan key
        c. Delete(key) untuk menghapus data di Map menggunakan key
        d. Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map

- Cond (Condition)
    1. Cond adalah implementasi locking berbasis kondisi
    2. Cond membutuhkan Locker(bisa menggunakan Mutex atau RWMutex) untuk implementasi locking nya, namun berbeda dengan Locker biasanya, di Cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
    3. function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
    4. untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)

- Atomic
    1. go-lang memiliki package yang bernama sync/atomic
    2. atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
    3. contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter. Hal ini sebenarnya bisa digunakan menggunakan Atomic package
    4. ada banyak function di atomic package, kita bisa eksplore sendiri di halaman dokumentasinya
    5. https://golang.org/pkg/sync/atomic/

- Timer
    1. Timer adalah representasi satu kejadian
    2. ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
	3. untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)

- time.After()
    1. kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
    2. untuk melakukan hal itu kita bisa menggunakan function time.After(duration)

- time.AfterFunc()
    1. kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
    2. kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
    3. kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil ketika Timer mengirim kejadiannya

- time.Ticker
    1. Ticker adalah representasi kejadian yang berulang
    2. ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
    3. untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
    4. untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()

- time.Tick()
    1. kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja
    2. jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja

- GOMAXPROCS
    1. sebelumnya diawal kita sudah bahas bahwa goroutine itu sebenarnya dijalankan di dalam Thread
    2. pertanyaannya, seberapa banyak Thread yang ada di go-lang ketika aplikasi kita berjalan?
    3. untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread
    4. secara default, jumlah thread di go-lang itu sebanyak jumlah CPU di komputer kita.
	5. kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()
*/