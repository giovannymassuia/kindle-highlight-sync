# Kindle Highlights Extractor

This is a simple Go program that extracts kindle highlights from the read.amazon.com website.

## Usage

1. Go to [read.amazon.com](https://read.amazon.com/notebook)
2. Extract the following cookies from the website:
   - `ubid-main`
   - `at-main`
   - `x-main`
3. Add the cookies to the environment variables:
   ```bash
   export UBID_MAIN="cookie_value"
   export AT_MAIN="cookie_value"
   export X_MAIN="cookie_value"
   ```
4. Run the program:
   ```bash
   go run main.go
   ```

## ‼️ Attention

This program is not intended to be used for any malicious purposes. It is only for educational purposes.
The program does not store any data and only extracts the highlights from the website.

Also, the cookies used here are very sensitive and should not be shared with anyone.
They are only used to authenticate the user and extract the highlights from the website.

So, be careful with the cookies and do not share them with anyone.
