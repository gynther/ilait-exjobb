package readHandler

import 

// fileName is not used at this time. it will be used later with database access.
// database fetch requests will be places in here.
func readChunk ( fileName string, chunkNumber int ) (buf []byte) {
	buf = make( []byte, chunkSize )
	file, err := os.Open( fmt.Sprintf(testOutput, chunkNumber ), os.O_RDONLY, 0 )
	if err != nil {
		return nil
	}
	// bytesRead, err := file.ReadAt(buf, int64(chunkNumber*chunkSize) )
	bytesRead, err := file.ReadAt(buf, int64(chunkNumber*chunkSize) )
	if err != nil && bytesRead >= 0 {
		fmt.Printf( "ERROR:readChunk():%s:%s\n", err, string(buf) )
	}
	
	// temporal for testing purposes
	os.Remove( fmt.Sprintf( testOutput, chunkNumber ) )
	
	return buf
}

func LoadFile ( fileName string ) {
	// output file
	file, _ := os.Open( fmt.Sprintf( fileName ), os.O_RDWR | os.O_CREATE, 0666)
	defer file.Close()
	// buf := make([]byte, chunkSize)
	
	for i := 0 ; ; i++ {
		buf := readChunk ( fileName, i ) 
		if buf == nil {
			fmt.Printf( "blargh EOF" )
			break
		}
		_, err := file.WriteAt( buf, int64( i*chunkSize ) )
		
		if err != nil {
			fmt.Printf( "blargh write errors" )
			break
		}
	}
}
