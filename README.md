# Overview

Simple dependency injection container, inspired by some popular framework like pimple/laravel.

# Usage

## Bind

Use `Bind` to create new instance when calling object.

```go
ioc.Bind("tester", func() interface{} {
	return &objTest{
		"test",
		"default",
	}
})

obj := ioc.Make("tester")

//different instance will be returned for each call
obj2 := ioc.Make("tester")

obj.name = "test-update"
obj.name == obj2.name //false
```

## Singleton

Use `Singleton` if you want to call same instance anytime object is called.

```go
ioc.Singleton("tester", func() interface{} {
	return &objTest{
		"test",
		"default",
	}
})

obj := ioc.Make("tester")

//same instance will be returned for each call
obj2 := ioc.Make("tester")

obj.name = "test-update"
obj.name == obj2.name //true
```

# Contributing

Please feel free to submit issue and pull request
