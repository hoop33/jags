# Jags

> Scrapes the Jaguars' roster from the official website and creates a CSV

## Installation

```sh
$ go get github.com/hoop33/jags
```

## Usage

```sh
$ jags
```

## Credits

* [goquery](https://github.com/PuerkitoBio/goquery)

## Disclaimers

* It works as of May 16, 2018
* It scrapes <https://www.jaguars.com/team/players-roster/>
* It's vulnerable to a MITM attack (TLS verification is turned off) because reasons
* Fields aren't surrounded by quotes, so a comma in a value could throw things off (I could surround with quotes, and then escape embedded quotes, and then . . .)

## License

Copyright &copy; 2018 Rob Warner

Licensed under the [MIT License](https://hoop33.mit-license.org/)

