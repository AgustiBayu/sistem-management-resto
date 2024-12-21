CREATE TABLE transaksis(
    id SERIAL PRIMARY KEY,
    pesanan_id INT NOT NULL,
    metode_pembayaran TEXT CHECK (metode_pembayaran IN ('CASH', 'CARD', 'EWALLET')),    
    jumlah_bayar INT NOT NULL,
    kembalian INT NOT NULL,
    tanggal_transaksi DATE NOT NULL,
    FOREIGN KEY(pesanan_id) REFERENCES pesanans(id)
);