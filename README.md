# SNIIM Scraper

Scraper para obtener los datos de precios sobre algunos productos como frutos y carnes en M茅xico.


## 馃搫 Variables de entorno

| Nombre | Descripci贸n | Default |
|--|--|--|
| SNIIM_ADDR | Direcci贸n del sitio fuente de la informaci贸n | http://www.economia-sniim.gob.mx |
| CATALOGUE_SRC | Nombre de la base de datos o directorio del  filesystem para guardar los cat谩logos | SNIIM_DATA |
| DEBUG | Bandera para habilitar el modo debug | false |
| MONGO_URI* | Direcci贸n de la base de datos mongo, solo se intentar谩 conectar si est谩 presente. | '' |


\* [TODO](./docs/README.md)
## 馃摝 Instalaci贸n

Para realizar una instalaci贸n de la herramienta puede visitar la [p谩gina de releases](https://github.com/everitosan/sniim-scrapper/releases) o bien descargar el c贸digo fuente y seguir los pasos de [compilaci贸n](./docs/Compilation.md).

Si decide usar una versi贸n [precompilada]((https://github.com/everitosan/sniim-scrapper/releases)), una vez que descargue el archivo debe descomprimir el archivo .tar.gz

<br>

*馃惂 Ejemplo en Linux*
```bash
# Descomprimir con tar
$ tar -xf sniim-scraper_0.0.1_Linux_arm64.tar.gz

# Ejecuci贸n de prueba
$ ./sniim-cli -h                        
Usage:
  sniim-cli [flags]
  sniim-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  consult     Manage consults
  help        Help about any command
  init        Create catalogues
  request     Request information

Flags:
  -h, --help      help for sniim-cli
  -v, --version   Show version of the cli

Use "sniim-cli [command] --help" for more information about a command.

```




## 馃崌 Uso

Inicializaci贸n
---

El CLI necesita tener los cat谩logos disponibles para poder operar.

Para esto existe el comando init. 

馃暪锔? *Carga de cat谩logos.*

```bash
sniim-cli init
```
![](./docs/img/init.png)

Registro de consultas
---

Para poder saber el precio de alg煤n art铆culo, podemos crear una consulta. 

El comando **consult** con la bandera **-c** 贸 **--create** invocar谩 una serie de preguntas basadas en la categor铆a y art铆culo seleccionados.

Si deseamos guardar esa consulta para volver a ejecutarla posteriormente, podemos agregar la bandera **-s** 贸 **--save**.

馃暪锔? *Crea un registro de consulta.*

```bash
sniim-cli consult -c
```
![](./docs/img/consult--create.png)

> 馃摑 Nota:  
*Al usar la palabra reservada **now**, estaremos indicando que la fecha de inter茅s ser谩 la fecha en la que se ejecute la petici贸n con el comando **request**.*

Tambi茅n podemos mostar una lista de las consultas que tenemos guardadas en registros.

馃暪锔? *Listar las consultas registradas.*

```bash
sniim-cli consult -l
```

![](./docs/img/consult--list.png)

Ejecuci贸n de Consultas
---

Si deseamos repetir alguna de las consultas guardadas, podemos usar el comando request.

馃暪锔? *Realiza una consulta.*

```bash
sniim-cli request -i [铆ndice]
```

![](./docs/img/request--index.png)

> 馃摑 Nota:  
***-s** Es una bandera opcional que se puede agregar al comando de request para guardar en resultado obtenido en lugar de mostrarlo en la consola*