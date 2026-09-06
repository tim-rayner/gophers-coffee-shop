```
╔══════════════════════════════════════════════════════════╗
║                                                          ║
║               ☕  GOPHERS COFFEE SHOP  ☕                  ║
║                ~ freshly brewed in Go ~                  ║
║                                                          ║
╚══════════════════════════════════════════════════════════╝
```

<p align="left">
  <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" width="180" alt="The Go gopher, pouring coffee in spirit if not in pixels" />
</p>

<p align="left">
  <em>A cheerful, over-caffeinated little gopher runs the till here.<br/>
  He's new at this. Be patient with him.</em>
</p>

---

## About the Shop

This is my **first ever Go project** — and my first attempt at building a
(somewhat) functional CLI app. It's a little interactive coffee shop that
takes your order at the terminal, tallies up your bill, and prints you a
proper little ASCII receipt on the way out.

I'm brand new to Go, so please don't expect production-grade code behind
the counter. What's here was built with:

- 📗 Two Codecademy lessons: _Learn Go_ and _Learn Go: Loops, Arrays,
  Maps, and Structs_
- 📘 The first few modules of the official [**Tour of Go**](https://go.dev/tour/)

If you spot rough edges (there will be rough edges), that's the learning
showing through, not a bug report waiting to happen. 😄

---

## The Menu

```
╔═════════════════════════════════════════════════════════════════════════════════════════╗
║                                   GOPHERS COFFEE SHOP                                   ║
║                                   ~ freshly brewed ~                                    ║
╠════════════════════════════════════════════╦════════════════════════════════════════════╣
║                 · DRINKS ·                 ║                 · SYRUPS ·                 ║
║                                            ║                                            ║
║ Decaf ······························ £2.00 ║ Caramel ···························· £1.00 ║
║ Espresso ··························· £3.00 ║ Chocolate ·························· £1.00 ║
║ Latte ······························ £4.00 ║ Hazelnut ··························· £1.00 ║
║ Cappuccino ························· £5.00 ║ Irish Cream ························ £1.00 ║
║ Americano ·························· £6.00 ║ Maple ······························ £1.00 ║
║ Mocha ······························ £7.00 ║ Peppermint ························· £1.00 ║
║ Macchiato ·························· £8.00 ║ Pumpkin Spice ······················ £1.00 ║
║ Iced Coffee ························ £9.00 ║ Vanilla ···························· £1.00 ║
║                                            ║ Cinnamon ··························· £1.05 ║
║                                            ║ Gophers Special Syrup ············ £200.00 ║
╠════════════════════════════════════════════╬════════════════════════════════════════════╣
║                 · SIZES ·                  ║              · TEMPERATURE ·               ║
║                                            ║                                            ║
║ Small ······························ £0.50 ║ Hot                                        ║
║ Medium ····························· £1.00 ║ Iced                                       ║
║ Grande ····························· £1.50 ║ Warm                                       ║
║ Venti ······························ £2.00 ║                                            ║
╚════════════════════════════════════════════╩════════════════════════════════════════════╝
```

> ⚠️ Order the **Gophers Special Syrup** at your own risk... lol.

The real menu is drawn straight from the code at runtime (sorted by price,
box-drawn and centered by hand), so what you see above is close, but the
app itself is the source of truth.

---

## How it works

Run it, and the gopher behind the counter will walk you through your order,
one question at a time:

1. What drink can I get you?
2. Which syrup would you like?
3. What size is that?
4. Hot, Warm, or Over Ice?

Say **"no"** when asked if there's anything else, and you'll get your total
plus a fully itemised receipt... box-drawn, of course.

```
$ go run .

...menu prints above...

Welcome to Gophers Coffee Shop!
Lets start with your drink, what can I get you?
> Latte
And which syrup would you like?
> Caramel
And what size is that in?
> Grande
And is that Hot, Warm, or Over Ice?
> Hot

{Latte 4} {Caramel 1} {Grande 1.5} {Hot 0} coming right up!
Anything else?
> no
Your order comes to £6.50

╔════════════════════════════════════════════╗
║            GOPHERS COFFEE SHOP             ║
║             ~ freshly brewed ~             ║
║                  RECEIPT                   ║
╠════════════════════════════════════════════╣
║ #1  Grande Hot Latte                       ║
║  Latte ···························· £4.00  ║
║  Caramel syrup ···················· £1.00  ║
║  Grande ··························· £1.50  ║
║  item ····························· £6.50  ║
╠════════════════════════════════════════════╣
║  TOTAL ···························· £6.50  ║
╠════════════════════════════════════════════╣
║         Thank you for your custom          ║
║             See you next brew!             ║
╚════════════════════════════════════════════╝

your order will be with you in about 5 seconds, speedy service I know
```

_(that `{Latte 4} {Caramel 1} ...}` line is a known rough edge. Printing a
struct with `%v` instead of a proper description. Charming, in its own way.)_

---

## Running the app

Requires Go 1.21+ (built and tested with go.mod targeting `go 1.27.1`).

```bash
git clone git@github.com:tim-rayner/gophers-coffee-shop.git
cd gophers-coffee-shop
go run .
```

---

## What I learned brewing this

- Structs and methods (`Drink`, `Syrup`, `Size`, `Temperature`, `Order`,
  `Menu`, `Receipt`) - modelling a real-ish thing as a handful of small types
- Maps for menu data, plus sorting a map's contents into a stable, priced
  order for display
- Pointers (`*OrderItem`, `*Order`) for mutating state across a loop
  without returning a new struct every time
- `bufio.Scanner` for reading line-by-line terminal input
- String formatting and padding by hand to draw the menu and receipt boxes
  with `strings.Repeat`, `utf8.RuneCountInString`, and a lot of trial and error
- Basic control flow patterns: nested loops as a makeshift state machine
  for walking through an order, step by step

---

## Known rough edges

- `Barista.MakeCoffee` / `Barista.ServeCoffee` are stubs... the gopher just
  _says_ your coffee is coming
- Input handling forgives case (`Latte` = `latte` = `LATTE`) and trims
  leading/trailing whitespace, but not typos or extra spaces mid-word
- No tests yet, next thing on the list once I'm past "Tour of Go"

---

<p align="center">
  <sub>Built by a very caffeinated Go beginner.</sub>
</p>
