# Cointracking Export File Converter

A converter for [CoinTracking](https://cointracking.info/) export .csv files. 

## About

CoinTracking is a widely used crypto portfolio tracker. Though it has many integrations for all kinds of crypto exchanges it lacks in my opinion of a good user experience.
As I tried out [Blockpit](https://www.blockpit.io/), i found that this is much better but it lacks of an import from CoinTracking exports in order to have an easy migration path.
This was the reason why I created this tool which takes CoinTracking export files in `.csv` format and converts it into Blockpit's `.xslx` files for import.

Currently there is only support for a conversion to Blockpit but it is possible to extend it in the future to other portfolio tracker.

## Features

- Convert trade transaction fees from CoinTracking to Blockpit as they are handled differently (see: [Blockpit](https://help.blockpit.io/hc/en-us/articles/360011882020-Transaction-Label-Trade) and [CoinTracking](https://cointracking.freshdesk.com/en/support/solutions/articles/29000039588-how-fees-are-handled-within-cointracking) documentation)
- Handle CoinTracking 'swap' transactions either as 'trade' transaction or split it into two non-taxable transactions (out and in)
- User configuration for custom mapping between CoinTracking and Blockpit transaction types
- User configuration for timezone of the timestamps in the CoinTracking export file 
- Skipping of empty transactions that don't have an outgoing, incoming or fee amount
- Generated a transaction ID as CoinTracking does not include the ID in the export file

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

### Development dependencies

- Taskfile

To build a redistributable, production mode package, use `wails build`.

## Known Issues

CoinTracking does not export the transaction IDs in their `.csv` files. Transaction IDs are used by portfolio trackers to distinguish if a transaction is a duplicate of another one. 
Therefore, the tool creates a unique transaction ID from the transaction data. This means that Blockpit can see if the transaction was already imported but if you have created another integration in Blockpit that has imported this transaction automatically, the IDs are different and Blockpit won't notice that these transactions are the same.
