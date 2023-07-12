import { GoogleSpreadsheet } from 'google-spreadsheet';
import { JWT } from 'google-auth-library'
import { GenerateCSV } from './csv'

const dotenv = require('dotenv');
dotenv.config()

const SHEET_ID = process.env.GOOGLE_SHEET_ID // spreadsheet key is the long id in the sheets URL

// Initialize auth - see https://theoephraim.github.io/node-google-spreadsheet/#/guides/authentication
const serviceAccountAuth = new JWT({
  // env var values here are copied from service account credentials generated by google
  // see "Authentication" section in docs for more info
  email: process.env.GOOGLE_SERVICE_ACCOUNT_EMAIL,
  key: process.env.GOOGLE_PRIVATE_KEY,
  scopes: [
    'https://www.googleapis.com/auth/spreadsheets',
  ],
});

//
const main = async () => {
  const doc = new GoogleSpreadsheet(SHEET_ID, serviceAccountAuth);

  await doc.loadInfo(); // loads document properties and worksheets
  console.log(doc.title);
  
  const sheet = doc.sheetsByIndex[2]; // select "Weekly data consumed" tab
  const MAX_ROW = 500; // max row count to 500 : assume we don't have more than 500 pools
  const row_range = 'A1:D' + MAX_ROW

  // loads range of cells into local cache - DOES NOT RETURN THE CELLS
  await sheet.loadCells(row_range);
  
  // total cells, loaded, how many non-empty
  console.log(sheet.cellStats); 

  // get next epoch elys allocation
  let next_elys_allocation = sheet.getCell(2, 3).value
  if (!next_elys_allocation) {
    return
  }

  // Round decimal 0
  next_elys_allocation = (next_elys_allocation as number).toFixed(0)

  // csv data input
  let data = []

  // empty string
  const emptyString = new String()
  
  // start row = 2
  let row = 2
  do {
    // Get pool info
    const pool = sheet.getCell(row,1).value
    if (!pool) {
      break
    }

    // get multiplier, fix the decimal number to 3 as there is conversion between javascript and google spread sheet
    const multiplier = (sheet.getCell(row,2).value as number).toFixed(3)
    // insert data to csv data input
    data.push({'poolId':row-2, 'poolName':pool, 'multiplier': multiplier,'eden_next_allocation': emptyString })
    row ++
  }
  while (row < MAX_ROW); 

  // insert eden next epoch amount
  data.push({'poolId': emptyString, 'poolName': emptyString, 'multiplier': emptyString, 'eden_next_allocation': next_elys_allocation})

  // Generate csv file and output
  GenerateCSV(data)
}

// call main function
main()