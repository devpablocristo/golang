1. Leer Repo

2. Leers monitor.yml

2. Analizar Inversion de dependencias

2.1 Analizar solo paquetes que esten en los layes Domain o Application
2.2 Verificar el paquete actual es domain o application

2.3 Si es domain

2.4 Si es application
obtener lista de variales y tipos de la variable implementadas

tengo la lista de los directorios y los layers
tengo los archivos filtrados por domian y application
tengo a que paqute pertence cada uno 


 
3. Imprimir violanciones de DIP

Resumen de cómo se calculan los scores:
Score 1: Si el dominio depende de algo que no pertenece al dominio.
Score 2: Si la aplicación depende directamente de la infraestructura.
Score 3: Si no se detectan violaciones en ninguna de las capas.