Estudiar:
Libreria estandar

    1.json
        json.NewEncoder(w).Encode(sl)
       	json.NewDecoder(r.Body).Decode(&l)
    2.log
        log.Println()
    3.fmt
        fmt.Fprintf()


Librerias no estandar
Las libs no estandar hay que instlarlas a mano.
Si estan ya descargadas, code las agrega automaticamente.

go get github.com/gorilla/mux
go get github.com/subosito/gotenv
go get github.com/eaigner/jet
go get github.com/lib/pq

    1. gorilla/mux
        mux.Vars()
        mux.NewRouter()
        HandleFunc()





PostgreSQL

https://api.elephantsql.com/console/1d3605ef-151b-4528-b959-85d9afbae85a/browser?

cuenta creada con tucbox@gmail.com

create table libros (id serial, titulo varchar, autor varchar, año varchar);

insert into libros (titulo, año, autor) values('Dune','1965','Frank Herbert');

insert into libros (titulo, año, autor) values('Cita con Rama', '1974', 'Arthur C. Clarke');

insert into libros (titulo, año, autor) values(	
'Un guijarro en el cielo', '1950', 'Isaac Asimov');

insert into libros (titulo, año, autor) values('Redshirts', '2013', 'John Scalzi');

insert into libros (titulo, año, autor) values(
'Hyperion', '1990', 'Dan Simmons');

select * from libros;







