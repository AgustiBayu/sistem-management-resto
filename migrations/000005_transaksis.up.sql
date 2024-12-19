CREATE TABLE transaksis(
    id INT PRIMARY KEY AUTO_INCREMENT,
    pesanan_id INT NOT NULL,
    metode_pembayaran ENUM ('CASH', 'CARD', 'EWALLET') NOT NULL,
    jumlah_bayar INT NOT NULL,
    kembalian INT NOT NULL,
    tanggal_transaksi DATE NOT NULL,
    FOREIGN KEY(pesanan_id) REFERENCES pesanans(id)
);