# 🗑️🪰 goka-scraper

![GitHub release (latest by tag)](https://img.shields.io/github/v/release/miha-staric/goka-scraper)
![Language](https://img.shields.io/badge/language-Go-Green)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-blue)
![License](https://img.shields.io/github/license/miha-staric/goka-scraper)

[VoKa scraper](https://github.com/miha-staric/voka-scraper) but this time written in Go. It's a scraper for VoKa underground garbage cans.

![GoKa Scraper in action!](https://github.com/user-attachments/assets/0a0c8da7-2fca-4b0c-a1a8-077f07b24775)

## 📘 Background

Ljubljana, the capital of Slovenia, has recently built a network of underground garbage bins, managed by the city's waste disposal company VoKa Snaga. These bins are aimed at enhancing urban cleanliness and sustainability and are installed primarily in the city center. The units replace traditional street-level bins, reducing visual clutter and improving public space aesthetics. The system comprises numerous underground collection units, each serving multiple waste categories, including glass, paper, and packaging, which are free to use, while biological and residual waste need to be charged for.

Residents and businesses access the units for mixed and biological trash using a special RFID card. Each of these cards has its own `chip card number` and a `password` which one can use to check the number of disposals that have been made on their account. The official website for checking the disposal data leaves much to be desired. This is why the author of `voka-scraper` and `goka-scraper` decided to take the matter into his own hands and build this user-friendly app to skip the not-so-pleasant web interface.

![VoKa Underground Bins With Trash Around](https://github.com/user-attachments/assets/500d6110-6a88-46d7-af7b-dd0d9505556a)

## 🗑️ Disposals

To quote the VoKa's webpage at <https://www.mojiodpadki.si/odpadki/podzemne-zbiralnice>, the fixed disposals included in the price are:

BIO: 8x
MKO: 6x

The price for extra BIO disposal is 0.1350 EUR and for MKO disposal 2.5626 EUR.

Note that BIO is Biological trash and MKO is Residual trash (mešani komunalni odpadki).

The following was taken from the website on 22. 3. 2026, so the current prices may be higher!

```txt
Cena enega vnosa preostanka odpadkov je 2,5626 €, cena enega vnosa BIO odpadkov pa 0,1350 €.
Mesečni strošek za ravnanje z odpadki za štirinajst minimalnih vnosov (šestkrat preostanek
odpadkov in osemkrat BIO odpadki) skupaj z DDV znaša 16,46 €.
```

## 🖨️ Output options

You have three different modes of output to choose from, `default`, `months`, and `years`.

### 🧻 Default Mode

The default mode displays all the data of each of the dumpings between two dates, together with the location and time.

### 🗓️ Months Mode

This one is the most useful, as it groups the data to particular months for the period between two exact dates and also adds the columns of `real_cost` and `total_cost`. The `real_cost` is the sum of all the particular dumpings (either BIO or MKO), multiplied by the price for each type of dumping. The `total_cost` is the actual cost that you are supposed to pay for the dumpings, because the system has a minimum set of each of the dumpings already included in the monthly bill, so if you don't use all the included dumpings, you still pay the minimum cost, whereas if you make extra dumpings per month, those will be paid on top.

### ⌛ Years Mode

This mode groups all the BIO and MKO dumpings made in a particular date for the period between the two exact dates. Nothing special, but shows you how many trash you have produced in that timeframe.

## 🔧 Configuration

Copy the example config and fill in your credentials, your desired mode, and date ranges:

```bash
cp .env.example .env
```

Or you can also use environment variables to set the desired configuration.

```env
# Mode (default, months, years)
GOKA_MODE=default

# Log level
GOKA_LOGLEVEL=info

# Card data
GOKA_CARD={card number}
GOKA_PASS={card password}

# URLs
GOKA_LOGINURL=https://potko.ekoplus.si/vs_uporabniki/login
GOKA_DASHBOARDURL=https://potko.ekoplus.si/vs_uporabniki/dashboard
GOKA_BASEURL=https://potko.ekoplus.si

# Date ranges
GOKA_DATEFROM=2026-03-01
GOKA_DATETO=2026-03-31

# Costs of dumpings
GOKA_COSTBIO=0.1350
GOKA_COSTMKO=2.5626
GOKA_MINBIO=1.0800
GOKA_MINMKO=15.3756
```

## 🐳 Running the app using Docker

Docker support is right around the corner.
