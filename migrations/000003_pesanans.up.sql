CREATE TABLE pesanans (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    total_harga INT NOT NULL,
    status ENUM ('PENDING', 'PAID', 'CANCELLED') NOT NULL,
    tanggal_pembuatan_pesanan DATE NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);