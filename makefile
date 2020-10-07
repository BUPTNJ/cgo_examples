my_buffer_capi.o: *.cpp *.h
	g++ -c -std=c++11 my_buffer_capi.cpp
libbuffer.a: my_buffer_capi.o
	ar -crs libbuffer.a my_buffer_capi.o
run: *.go libbuffer.a
	go run *.go
clean: 
	rm *.a *.o 
