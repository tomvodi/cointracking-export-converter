# Cointracking Export File Converter

A converter for [CoinTracking](https://cointracking.info/) export .csv files. 

## About

CoinTracking is a widely used crypto portfolio tracker. Though it has many integrations for all kinds of crypto exchanges it lacks in my opinion of a good user experience.
As I tried out [Blockpit](https://www.blockpit.io/), i found that this is much better but it lacks of an import from CoinTracking exports in order to have an easy migration path.
This was the reason why I created this tool which takes CoinTracking export files in `.csv` format and converts it into Blockpit's `.xslx` files for import.

Currently there is only support for a conversion to Blockpit but it is possible to extend it in the future to other portfolio tracker.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

### Known Issues

CoinTracking does not export the transaction IDs in their `.csv` files. Transaction IDs are used by portfolio trackers to distinguish if a transaction is a duplicate of another one. 
Therefore, the tool creates a unique transaction ID from the transaction data. This means that Blockpit can see if the transaction was already imported but if you have created another integration in Blockpit that has imported this transaction automatically, the IDs are different and Blockpit won't notice that these transactions are the same.
