# My Golang Gallery

## Database Design

Using [GORM](https://gorm.io/) it will be easier to manage and handle db actions (CRUD). GORM provide an OOut Of the Box DB Design:

```text
// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

I extended this model for `Image` db to include some requirde features:

```text
// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
// gorm.Model definition
// ImageDB
type Image struct {
  ID            uint           `gorm:"primaryKey"`
  Name          string
  UserID        uint
  UploadedBy    User
  CreatedAt     time.Time
  UpdatedAt     time.Time
  DeletedAt     gorm.DeletedAt `gorm:"index"`
}
```

For UserDB I used almost the same model:

```text
// GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
// gorm.Model definition
// UserDB
type User struct {
  ID            uint           `gorm:"primaryKey"`
  Name          string
  Password      string
  NickName      string
  Avatar        string
  CreatedAt     time.Time
  UpdatedAt     time.Time
  DeletedAt     gorm.DeletedAt `gorm:"index"`
}
```
