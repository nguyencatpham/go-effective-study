package user_test

import (
	"testing"

	"github.com/nguyencatpham/go-effective-study/pkg/api/user"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/mock"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/mock/mockdb"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type args struct {
		c   echo.Context
		req model.User
	}
	cases := []struct {
		name     string
		args     args
		wantErr  bool
		wantData *model.User
		udb      *mockdb.User
		rbac     *mock.RBAC
		sec      *mock.Secure
	}{{
		name: "Fail on is lower role",
		rbac: &mock.RBAC{
			AccountCreateFn: func(echo.Context, model.AccessRole, int, int) error {
				return model.ErrGeneric
			}},
		wantErr: true,
		args: args{req: model.User{
			FirstName: "John",
			LastName:  "Doe",
			Username:  "JohnDoe",
			RoleID:    1,
			Password:  "Thranduil8822",
		}},
	},
		{
			name: "Success",
			args: args{req: model.User{
				FirstName: "John",
				LastName:  "Doe",
				Username:  "JohnDoe",
				RoleID:    1,
				Password:  "Thranduil8822",
			}},
			udb: &mockdb.User{
				CreateFn: func(db orm.DB, u model.User) (*model.User, error) {
					u.CreatedAt = mock.TestTime(2000)
					u.UpdatedAt = mock.TestTime(2000)
					u.Base.ID = 1
					return &u, nil
				},
			},
			rbac: &mock.RBAC{
				AccountCreateFn: func(echo.Context, model.AccessRole, int, int) error {
					return nil
				}},
			sec: &mock.Secure{
				HashFn: func(string) string {
					return "h4$h3d"
				},
			},
			wantData: &model.User{
				Base: model.Base{
					ID:        1,
					CreatedAt: mock.TestTime(2000),
					UpdatedAt: mock.TestTime(2000),
				},
				FirstName: "John",
				LastName:  "Doe",
				Username:  "JohnDoe",
				RoleID:    1,
				Password:  "h4$h3d",
			}}}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, tt.sec)
			usr, err := s.Create(tt.args.c, tt.args.req)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantData, usr)
		})
	}
}

func TestView(t *testing.T) {
	type args struct {
		c  echo.Context
		id int
	}
	cases := []struct {
		name     string
		args     args
		wantData *model.User
		wantErr  error
		udb      *mockdb.User
		rbac     *mock.RBAC
	}{
		{
			name: "Fail on RBAC",
			args: args{id: 5},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return model.ErrGeneric
				}},
			wantErr: model.ErrGeneric,
		},
		{
			name: "Success",
			args: args{id: 1},
			wantData: &model.User{
				Base: model.Base{
					ID:        1,
					CreatedAt: mock.TestTime(2000),
					UpdatedAt: mock.TestTime(2000),
				},
				FirstName: "John",
				LastName:  "Doe",
				Username:  "JohnDoe",
			},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return nil
				}},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (*model.User, error) {
					if id == 1 {
						return &model.User{
							Base: model.Base{
								ID:        1,
								CreatedAt: mock.TestTime(2000),
								UpdatedAt: mock.TestTime(2000),
							},
							FirstName: "John",
							LastName:  "Doe",
							Username:  "JohnDoe",
						}, nil
					}
					return nil, nil
				}},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			usr, err := s.View(tt.args.c, tt.args.id)
			assert.Equal(t, tt.wantData, usr)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		c   echo.Context
		pgn *model.Pagination
	}
	cases := []struct {
		name     string
		args     args
		wantData []model.User
		wantErr  bool
		udb      *mockdb.User
		rbac     *mock.RBAC
	}{
		{
			name: "Fail on query List",
			args: args{c: nil, pgn: &model.Pagination{
				Limit:  100,
				Offset: 200,
			}},
			wantErr: true,
			rbac: &mock.RBAC{
				UserFn: func(c echo.Context) *model.AuthUser {
					return &model.AuthUser{
						ID:         1,
						CompanyID:  2,
						LocationID: 3,
						Role:       model.UserRole,
					}
				}}},
		{
			name: "Success",
			args: args{c: nil, pgn: &model.Pagination{
				Limit:  100,
				Offset: 200,
			}},
			rbac: &mock.RBAC{
				UserFn: func(c echo.Context) *model.AuthUser {
					return &model.AuthUser{
						ID:         1,
						CompanyID:  2,
						LocationID: 3,
						Role:       model.AdminRole,
					}
				}},
			udb: &mockdb.User{
				ListFn: func(orm.DB, *model.ListQuery, *model.Pagination) ([]model.User, error) {
					return []model.User{
						{
							Base: model.Base{
								ID:        1,
								CreatedAt: mock.TestTime(1999),
								UpdatedAt: mock.TestTime(2000),
							},
							FirstName: "John",
							LastName:  "Doe",
							Email:     "johndoe@gmail.com",
							Username:  "johndoe",
						},
						{
							Base: model.Base{
								ID:        2,
								CreatedAt: mock.TestTime(2001),
								UpdatedAt: mock.TestTime(2002),
							},
							FirstName: "Hunter",
							LastName:  "Logan",
							Email:     "logan@aol.com",
							Username:  "hunterlogan",
						},
					}, nil
				}},
			wantData: []model.User{
				{
					Base: model.Base{
						ID:        1,
						CreatedAt: mock.TestTime(1999),
						UpdatedAt: mock.TestTime(2000),
					},
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@gmail.com",
					Username:  "johndoe",
				},
				{
					Base: model.Base{
						ID:        2,
						CreatedAt: mock.TestTime(2001),
						UpdatedAt: mock.TestTime(2002),
					},
					FirstName: "Hunter",
					LastName:  "Logan",
					Email:     "logan@aol.com",
					Username:  "hunterlogan",
				}},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			usrs, err := s.List(tt.args.c, tt.args.pgn)
			assert.Equal(t, tt.wantData, usrs)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}

}

func TestDelete(t *testing.T) {
	type args struct {
		c  echo.Context
		id int
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		udb     *mockdb.User
		rbac    *mock.RBAC
	}{
		{
			name:    "Fail on ViewUser",
			args:    args{id: 1},
			wantErr: model.ErrGeneric,
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (*model.User, error) {
					if id != 1 {
						return nil, nil
					}
					return nil, model.ErrGeneric
				},
			},
		},
		{
			name: "Fail on RBAC",
			args: args{id: 1},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (*model.User, error) {
					return &model.User{
						Base: model.Base{
							ID:        id,
							CreatedAt: mock.TestTime(1999),
							UpdatedAt: mock.TestTime(2000),
						},
						FirstName: "John",
						LastName:  "Doe",
						Role: &model.Role{
							AccessLevel: model.UserRole,
						},
					}, nil
				},
			},
			rbac: &mock.RBAC{
				IsLowerRoleFn: func(echo.Context, model.AccessRole) error {
					return model.ErrGeneric
				}},
			wantErr: model.ErrGeneric,
		},
		{
			name: "Success",
			args: args{id: 1},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (*model.User, error) {
					return &model.User{
						Base: model.Base{
							ID:        id,
							CreatedAt: mock.TestTime(1999),
							UpdatedAt: mock.TestTime(2000),
						},
						FirstName: "John",
						LastName:  "Doe",
						Role: &model.Role{
							AccessLevel: model.AdminRole,
							ID:          2,
							Name:        "Admin",
						},
					}, nil
				},
				DeleteFn: func(db orm.DB, usr *model.User) error {
					return nil
				},
			},
			rbac: &mock.RBAC{
				IsLowerRoleFn: func(echo.Context, model.AccessRole) error {
					return nil
				}},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			err := s.Delete(tt.args.c, tt.args.id)
			if err != tt.wantErr {
				t.Errorf("Expected error %v, received %v", tt.wantErr, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		c   echo.Context
		upd *user.Update
	}
	cases := []struct {
		name     string
		args     args
		wantData *model.User
		wantErr  error
		udb      *mockdb.User
		rbac     *mock.RBAC
	}{
		{
			name: "Fail on RBAC",
			args: args{upd: &user.Update{
				ID: 1,
			}},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return model.ErrGeneric
				}},
			wantErr: model.ErrGeneric,
		},
		{
			name: "Fail on ViewUser",
			args: args{upd: &user.Update{
				ID: 1,
			}},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return nil
				}},
			wantErr: model.ErrGeneric,
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (*model.User, error) {
					if id != 1 {
						return nil, nil
					}
					return nil, model.ErrGeneric
				},
			},
		},
		{
			name: "Fail on Update",
			args: args{upd: &user.Update{
				ID: 1,
			}},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return nil
				}},
			wantErr: model.ErrGeneric,
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (*model.User, error) {
					return &model.User{
						Base: model.Base{
							ID:        1,
							CreatedAt: mock.TestTime(1990),
							UpdatedAt: mock.TestTime(1991),
						},
						CompanyID:  1,
						LocationID: 2,
						RoleID:     3,
						FirstName:  "Joanna",
						LastName:   "Doep",
						Mobile:     "334455",
						Phone:      "444555",
						Address:    "Work Address",
						Email:      "golang@go.org",
					}, nil
				},
				UpdateFn: func(db orm.DB, usr *model.User) error {
					return model.ErrGeneric
				},
			},
		},
		{
			name: "Success",
			args: args{upd: &user.Update{
				ID:        1,
				FirstName: mock.Str2Ptr("John"),
				LastName:  mock.Str2Ptr("Doe"),
				Mobile:    mock.Str2Ptr("123456"),
				Phone:     mock.Str2Ptr("234567"),
			}},
			rbac: &mock.RBAC{
				EnforceUserFn: func(c echo.Context, id int) error {
					return nil
				}},
			wantData: &model.User{
				Base: model.Base{
					ID:        1,
					CreatedAt: mock.TestTime(1990),
					UpdatedAt: mock.TestTime(2000),
				},
				CompanyID:  1,
				LocationID: 2,
				RoleID:     3,
				FirstName:  "John",
				LastName:   "Doe",
				Mobile:     "123456",
				Phone:      "234567",
				Address:    "Work Address",
				Email:      "golang@go.org",
			},
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, id int) (*model.User, error) {
					return &model.User{
						Base: model.Base{
							ID:        1,
							CreatedAt: mock.TestTime(1990),
							UpdatedAt: mock.TestTime(1991),
						},
						CompanyID:  1,
						LocationID: 2,
						RoleID:     3,
						FirstName:  "Joanna",
						LastName:   "Doep",
						Mobile:     "334455",
						Phone:      "444555",
						Address:    "Work Address",
						Email:      "golang@go.org",
					}, nil
				},
				UpdateFn: func(db orm.DB, usr *model.User) error {
					usr.UpdatedAt = mock.TestTime(2000)
					return nil
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := user.New(nil, tt.udb, tt.rbac, nil)
			usr, err := s.Update(tt.args.c, tt.args.upd)
			assert.Equal(t, tt.wantData, usr)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestInitialize(t *testing.T) {
	u := user.Initialize(nil, nil, nil)
	if u == nil {
		t.Error("User service not initialized")
	}
}
