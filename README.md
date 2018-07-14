# Zero to GraphQL Using Buffalo

The purpose of this example is to provide details as to how one would go about using GraphQL with the Buffalo Web Framework. Thus, I have created two major sections which should be self explanatory: Quick Installation and Tutorial Installation.

## Getting Started

## Software requirements

- [Go 1.10.3 or higher](https://golang.org/dl)

- [Buffalo 0.12.2 or higher](https://gobuffalo.io/en/docs/installation)

- PostgreSQL 10.4.0 or higher

## Communication

- If you **need help**, use [Stack Overflow](http://stackoverflow.com/questions/tagged/graphql). (Tag 'graphql')
- If you'd like to **ask a general question**, use [Stack Overflow](http://stackoverflow.com/questions/tagged/graphql).
- If you **found a bug**, open an issue.
- If you **have a feature request**, open an issue.
- If you **want to contribute**, submit a pull request.

## Quick Installation

1.  clone this repository

    ```
    $ git clone git@github.com:conradwt/zero-to-graphql-using-buffalo.git
    ```

2.  change directory location

    ```
    $ cd /path/to/zero-to-graphql-using-buffalo
    ```

3.  install dependencies

    ```
    WIP
    ```

4.  create, migrate, and seed the database

    ```
    $ buffalo db create
    $ buffalo db migrate
    $ buffalo task db:seed
    ```

5.  start the server

    ```
    $ buffalo dev
    ```

6.  navigate to our application within the browser

    ```
    WIP
    ```

7.  enter and run GraphQL query

    ```
    {
     person(id: "1") {
       firstName
       lastName
       username
       email
       friends {
         firstName
         lastName
         username
         email
       }
     }
    }
    ```

## Tutorial Installation

WIP

## Production Setup

Ready to run in production? Please [check our deployment guides](https://gobuffalo.io/en/docs/building).

## Buffalo References

- Official website: https://gobuffalo.io
- Guides: https://gobuffalo.io/en/docs/goth
- Docs: https://godoc.org/github.com/gobuffalo/buffalo
- Mailing list: WIP
- Source: https://github.com/gobuffalo

## GraphQL References

- Official Website: https://gobuffalo.io
- GraphQL Go: https://github.com/graphql-go/graphql

## Support

Bug reports and feature requests can be filed with the rest for the Zero to GraphQL using Buffalo project here:

- [File Bug Reports and Features](https://github.com/conradwt/zero-to-graphql-using-buffalo/issues)

## License

Zero to GraphQL using Buffalo is released under the [MIT license](https://mit-license.org).

## Copyright

copyright:: (c) Copyright 2018 Conrad Taylor. All Rights Reserved.
