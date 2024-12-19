CREATE TABLE menu_items(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name INT NOT NULL,
    deskripsi VARCHAR(255) NOT NULL,
    harga INT NOT NULL,
    tanggal_penambahan_item_menu DATE NOT NULL
);