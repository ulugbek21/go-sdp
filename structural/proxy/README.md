# Proxy Pattern

The Proxy pattern usually wraps an object to hide some of its characteristics. These
characteristics could be the fact that it is a remote object (remote proxy), a very heavy object such as a very big image or the dump of a terabyte database (virtual proxy), or a restricted access object (protection proxy). (Contreras, 2017)

## Implementation

For our example, in the codebase we will create a remote proxy, which works a cache of
objects before accessing a database. Let's imagine that we have a database with many users, but instead of accessing the database each time we want information about a user, we will have a First In First Out (FIFO) stack of users in a Proxy pattern (FIFO is a way of saying that when the cache needs to be emptied, it will delete the first object that entered first).

```go
type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found", id)
}

type UserListProxy struct {
	SomeDatabase           *UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

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
```

## Summary

Wrap proxies around types that need some intermediate action, like giving authorization to the user or providing access to a database, like in our example. 

Our example is a good way to separate application needs from database needs. If our application accesses the database too much, a solution for this is not in your database. Remember that the Proxy uses the same interface as the type it wraps, and, for the user, there shouldn't be any difference between the two.(Contreras, 2017)
