CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(255),
    pengguna TEXT CHECK (pengguna IN ('ADMIN', 'PENGGUNA')),
    tanggal_buat_akun DATE NOT NULL
);