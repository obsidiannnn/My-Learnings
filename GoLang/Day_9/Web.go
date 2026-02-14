package main

// const url = "https://jsonplaceholder.typicode.com/posts"

func main() {

	urll()

	// fmt.Println("Web Request")

	// response, err := http.Get(url)

	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Printf("Response is of type: %T \n", response)

	// // ans is response is of type  *http.Response
	// // * means reponse is in pntr format that is we are not getting copy of
	// // reponse but actual data.

	// defer response.Body.Close()

	// // using get or post or reading or writing never closes the request

	// // it is import to close it at the end

	// // reading the response

	// databyte, err := io.ReadAll(response.Body)

	// // databyte again will give data in bytes to we have to convert it into string.

	// if err != nil {
	// 	panic(err)

	// }

	// content := string(databyte)
	// fmt.Println("Data inside response if : \n", content)

	// // good practise if convert data into string store in varibale
	// // then print , do not directly conver into print statement.

}
