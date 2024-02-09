# godocs-playground

godocs Example and quick documentation, also serves as a playground to test all features for the godoc.

https://pkg.go.dev/golang.org/x/tools/cmd/godoc

## Requirements

- Install go on your computer
- Install godocs

```
go install golang.org/x/tools/cmd/godoc@latest
```

## Before to start

The godocs only will display the public functions or vars/constants that we created.

## DOCS

https://pkg.go.dev/golang.org/x/tools/cmd/godoc

## Descriptions

A comented line before a constant or function is going to create a decription regarding that.

If you write before the package declaration is goign to have 2 description, the synopsis (the ones appears on the home of the documentation page) and the overview (the one appears on the same package documentation page).

## Examples

On a test file you can create a `Example()` to address a global example from the file.

If you want to include a func you need to write as follows: `Example{MethodName}()`

but if you need a other example from the same func you can do it like this: `Example{MethodName}_{NameOfThisCase}`

## Command Line

To access on your local you need to run on your terminal:

```
godocs -http :8080
```

### Repo created because this videos

https://www.youtube.com/watch?v=80VT3xexcWs