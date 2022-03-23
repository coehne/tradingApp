import { format, parseISO } from "date-fns"


export const numberToUSD = (value: number) =>
  new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD'
  }).format(value);


export const stringToDate =(value: string)=> {
  return format(parseISO(value), "dd.MM.YYY HH:mm")
}