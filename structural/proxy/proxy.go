package proxy

import (
	"fmt"
)

// UserFinder should be implemented both by the database and proxy
type UserFinder interface {
	FindUser(id int32) (User, error)
}

// User ...
type User struct {
	ID int32
}

// UserList ...
type UserList []User

// FindUser ...
func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found", id)
}

// UserListProxy ...
type UserListProxy struct {
	SomeDatabase           *UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

// FindUser ...
func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)
	fmt.Println("Returning user from database")
	u.DidLastSearchUsedCache = false
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}
