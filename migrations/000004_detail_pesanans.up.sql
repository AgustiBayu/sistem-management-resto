CREATE TABLE detail_pesanans(
    id INT PRIMARY KEY AUTO_INCREMENT,
    pesanan_id INT NOT NULL,
    menu_item_id INT NOT NULL,
    jumlah_item_pesanan INT NOT NULL,
    harga_item INT NOT NULL,
    FOREIGN KEY(pesanan_id) REFERENCES pesanans(id),
    FOREIGN KEY(menu_item_id) REFERENCES menu_items(id)
);