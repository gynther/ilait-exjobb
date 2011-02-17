package main

import (
	"fmt"
	"os"
	"scanner"
	"readHandler"
	// "github.com/mikejs/gomongo/mongo"
	)

var chunkSize = 5
var testOutput = "testOutput%d.txt"

// Struct for databaseuseage
 type metadata struct {
	Originalnamn string
	Chunknummer int
}



func main() {

	var scanInput scanner.Scanner 
	scanInput.Init( os.Stdin ) 
	
	// main loop, reads scanner input and parses, EOF is strictly not used at this time.
	for scanInput.Scan() != scanner.EOF {
		switch scanInput.TokenText() {
		case "load","l":
			scanInput.Scan()
			fmt.Printf("Loading file \"%s\" from database\n", scanInput.TokenText() )
			go readHandler.LoadFile( scanInput.TokenText() )
		case "save","store","s":
			scanInput.Scan()
			fmt.Printf("Storing file \"%s\" to database\n", scanInput.TokenText() )
			go storeFile( scanInput.TokenText() )
		case "hora":
			fmt.Printf( "kingkong IS ONE!" )
			os.Exit(0)
		}
	}
}



func storeFile ( fileName string ) {
	file, err := os.Open( fileName, os.O_RDONLY, 0 )
	buf := make( []byte, chunkSize )
	defer file.Close()
	if err != nil {
		fmt.Printf( "ERROR (storeFile) (%s)\n", err )
		return
	}
	
	for  i := 0; ; i++ {
		bytesRead, err := file.ReadAt( buf, int64( i*chunkSize ) )
		if err != nil {
			writeChunk( buf[0:bytesRead], i )
			break
		} else {
			writeChunk( buf, i )
		}
	}
}

func writeChunk ( buf []byte, chunkNumber int ) {
	// here we need to know the mongoDB-ID for the chunk we are to save.
	// but in the meantime we just use a temporary name.

	file, err := os.Open( fmt.Sprintf( testOutput, chunkNumber ), os.O_RDWR | os.O_CREATE, 0666 )
	defer file.Close()
	if err != nil {
		fmt.Printf("ERROR (storeChunk: creating file) (error:%s)", err )
	}
	
	bytesWritten, err := file.Write( buf )
	if err != nil {
		fmt.Printf( "ERROR (storeChunk: writing file) (error:%s)", err )
	}
	fmt.Printf( "Written bytes:%d file:%s) output:%s\n" ,bytesWritten, fmt.Sprintf( testOutput, chunkNumber ) , string( buf ) )
}

/*func mongoWriter() {
	conn, _ := mongo.Connect("127.0.0.1")
	coll := conn.GetDB("gomongo").GetCollection("kuk")

	testDoc := &metadata {
		Originalnamn : fileToRead,
		Chunknummer : chunkFileNr,
	}

 	bsonDocIn, _ := mongo.Marshal(testDoc)
 	coll.Insert(bsonDocIn)

	fmt.Println("lol det fungerar\n")
	coll.Drop()	 
}*/
