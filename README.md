# Gotvl
Gotvl is a simple abstraction for message translation and internationalization `(i18n)` for Spanish and English languages (en, es), aimed at abstracting the implementation logic. It includes the necessary middleware for routing and default validation for Gin Validator. Note that you need to use `Accept-Language` in routes where this middleware is used.


## Installation

```go
go get -u github.com/juanmachuca95/gotvl
```

You can see a practical example in the [example](https://github.com/juanmachuca95/gotvl/tree/main/example) directory. 


## Requirements
This package only supports the Gin framework, so it only supports this router.

## Usage

This package uses [Gin](https://github.com/gin-gonic/gin) to obtain TVL, i.e. the translation for the specified language, default validation by language, and message locator for i18n in two languages.

It is necessary to have the [goi18n](https://github.com/nicksnyder/go-i18n#command-goi18n) tool installed for file generation.

Generate the translation files. You will need a translations folder where all `active.*.toml` files will be hosted. If you don't have it, you can use the `Makefile` provided by this repository below.

Finally, you can use this Makefile for the generation of the translation files that the middleware will consume.


```makefile
# Generate translations (en, es)
# Create by definitions
.PHONY: init
init:
	mkdir translations && cd translations; touch active.en.toml active.es.toml

.PHONY: gen
gen:
	cd translations && goi18n merge active.en.toml active.es.toml 

# Use the Finish command only when all translations have been completed.
.PHONY: finish
finish:
	cd translations; echo "\n" >> active.es.toml; cat translate.es.toml >> active.es.toml;

.PHONY: reset
reset: 
	cd translations; rm -rf active.es.toml translate.es.toml; touch active.en.toml

```


### Setting the middleware

```go
// Accept-Language (en or es) required
r.Use(gotvl.SetInstancesTranslate)
```

### Obtaining the instance
```go
tvl, err := gotvl.GetTVLContext(ctx)
```

## Contributing

We welcome contributions to Gotvl. To contribute, please follow these steps:

1. Fork the repository
2. Create a new branch for your feature or bugfix
3. Write tests for your changes
4. Implement your feature or bugfix
5. Commit your changes and push your branch to your forked repository
6. Open a pull request to this repository with a detailed description of your changes

Please ensure your code follows the Go coding style guidelines and that all tests pass before submitting a pull request. Thank you for contributing to Gotvl!

## License

This project is licensed under the terms of the MIT license. See [LICENSE](LICENSE) for more information.
