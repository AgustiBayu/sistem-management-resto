CREATE TABLE pesanans (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    total_harga INT NOT NULL,    
    status TEXT CHECK (status IN ('PENDING', 'PAID', 'CANCELLED')),
    tanggal_pembuatan_pesanan DATE NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);