package hconnc

import (
        "fmt"
        "net"
)

func cvictim(IP string, Port string)(connc net.Conn, err error){
	Addy := IP + ":" + Port

        lis, err := net.Listen("tcp", Addy)
        if err != nil {
                return connc, err
	}
	connc, err = lis.Accept()
	if err != nil{ 
        	return connc, err)
        } else {
            return 

	}

}
