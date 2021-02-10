## CalculateMac99 function
CalculateMac99 function represents ansi 9.9 algorithm.

### Example:

    data, _ := hex.DecodeString("4E6F77206973207468652074696D6520666F7220616C6C20ABCD")
    key1, _ := hex.DecodeString("0123456789ABCDEF")
    if mac99, er := ansimac.CalculateMac99(data, key1);err != nil {
      fmt.Printf("Mac 9.9:\t%v\n", strings.ToUpper(hex.EncodeToString(mac99)))
    }
    
## CalculateMac919 function
CalculateMac919 function represents ansi 9.19 algorithm.

### Example:

	data, _ := hex.DecodeString("4E6F77206973207468652074696D6520666F7220616C6C20ABCD")
	key1, _ := hex.DecodeString("0123456789ABCDEF")
	key2, _ := hex.DecodeString("FEDCBA9876543210")
	
	if mac919, er := ansimac.CalculateMac919(data, key1, key2); err != nil {
		fmt.Printf("Mac 9.19:\t%v\n", strings.ToUpper(hex.EncodeToString(mac919)))
	}
