CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(255),
    pengguna ENUM('ADMIN', 'PENGGUNA') NOT NULL,
    tanggal_buat_akun DATE NOT NULL
);