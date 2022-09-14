## Problem Statement
Ada sebuah bank memiliki nasabah 200 data. Buatlah fungsi end of day dimana data diambil dari Before Eod.csv (hasilnya di write ke After Eod.csv) kemudian proses dalam 1 waktu dengan step by step / proses sebagai berikut (gunakan multi-thread):
1. Hitung average balance setiap nasabah, dengan meng-average field ‘Balanced’ dengan
‘Previous Balanced’ kemudian update data ke field ‘Average Balanced’. (Jumlah Thread
bebas dan pastikan No Thread yang di gunakan tertulis di ‘No 1 Thread-No’)
2. Nasabah bisa mendapatkan benefit:
a. Jika balanced di antara 100-150 akan mengupdate free transfer menjadi 5
(Jumlah Thread bebas dan pastikan No Thread yang di gunakan tertulis di ‘No 2a
Thread-No’)
b. Jika balanced di atas 150 akan mendapatkan tambahan balanced sebesar 25
(Jumlah Thread bebas dan pastikan No Thread yang di gunakan tertulis di ‘No 2b
Thread-No’)

3. Bank memiliki budget 1.000 yang akan di bagikan ke 100 orang pertama (data urutan no 1-100) akan mendapatkan tambahan balance sebesar 10, untuk case ini buatlah 8 thread yang
akan berjalan secara bersamaan (pastikan No Thread yang di gunakan tertulis di ‘No 3
Thread-No’).

## Project Overview
I was going to use pipeline pattern but the problem was to maintain the order of data. When using pipeline pattern with fan in & fan out, we can't maintain the order of data.

Because of that, I decided to use simple go routines and wait group to solve this problem.

There are 6 main functions for this problem:
1. **readCsv** is used to read data from csv files
2. **prepareData** is used to convert string to eod struct according the After Eod.csv, this eod struct is a domain for this solution
3. **stage1** is used to calculate the average, averaged Balance will be updated
4. **stage2** is used to set the free transfer and balanced per user
   -  if the user meets the requirement of the stage 2a: Free Transfer and No2aThread will be updated
   -  if the user meets the requirement of the stage 2b: Balanced and No2bThread will be updated
   -  if the user doesn't meet the requirement stage 2a or 2b: No2aThread and No2bThread will be empty
5. **stage3** is used to give free credits for first 100 users: Balanced will be updated for 100 users only but No3ThreadNo will be updated to all users
6. **writeCsv** is used to write csv file for After Eod.csv

## Requirement
Install [Go](https://golang.org/doc/install) on your system

## Run
go run main.go

## Test
go test -p 1 -race ./... -coverprofile coverage.out  
