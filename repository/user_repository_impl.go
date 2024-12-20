package repository

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(name, email,password,pengguna,tanggal_buat_akun) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Pengguna, user.TanggalBuatAkun)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}
func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id,name, email,password,pengguna,tanggal_buat_akun FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Pengguna, &user.TanggalBuatAkun)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id,name, email,password,pengguna,tanggal_buat_akun FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Pengguna, &user.TanggalBuatAkun)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("email not found")
	}
}
func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id,name, email,password,pengguna,tanggal_buat_akun FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Pengguna, &user.TanggalBuatAkun)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("id user not found")
	}
}
func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET name = ?, email = ?,password = ?,pengguna = ?,tanggal_buat_akun = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Pengguna, user.TanggalBuatAkun, user.Id)
	helper.PanicIfError(err)
	return user
}
func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}