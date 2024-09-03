El comportamiento que observas, donde los logs se detienen después de hacer `Ctrl + S` en la terminal, no está relacionado directamente con `Air` o `Delve`, sino más bien con cómo la terminal maneja los comandos de flujo de control.

### Explicación

- **`Ctrl + S`** es un atajo de teclado que tradicionalmente se utiliza para **pausar el flujo de salida en una terminal**. Esto es parte de los controles de flujo XON/XOFF, que se remontan a los terminales seriales más antiguos.
- Cuando presionas `Ctrl + S`, la terminal entra en un estado "pausado" y deja de mostrar la salida. Puedes reanudar la salida presionando **`Ctrl + Q`**, que envía el comando para "continuar".

### Cómo Solucionarlo

1. **No presiones `Ctrl + S` en la terminal**: Evita usar este atajo en la terminal si no es necesario.

2. **Reanudar la salida con `Ctrl + Q`**:
   - Si accidentalmente presionas `Ctrl + S` y la salida se detiene, simplemente presiona `Ctrl + Q` para reanudar la salida.

3. **Deshabilitar los controles de flujo XON/XOFF en la terminal**:
   - En algunas terminales, puedes desactivar este comportamiento modificando la configuración. Por ejemplo, si estás utilizando `bash`, puedes agregar lo siguiente a tu archivo `.bashrc` o `.bash_profile` para desactivar el control de flujo:

   ```bash
   stty -ixon
   ```

   Después de agregar esto, guarda el archivo y ejecuta `source ~/.bashrc` o `source ~/.bash_profile` para aplicar los cambios.

### Verificación

Después de hacer cualquiera de estos cambios o ajustes, verifica que los logs continúen mostrando la salida esperada incluso después de presionar `Ctrl + S`.

### Conclusión

El problema no está en `Air` o `Delve`, sino en cómo la terminal maneja el control de flujo. Usando los pasos anteriores, puedes evitar o solucionar este comportamiento para que los logs se muestren continuamente sin interrupciones.

### Resumen Simplificado y Mejorado para Configurar Docker y `xbindkeys`

#### Problema Principal
Cuando la aplicación en Docker finaliza, presionar `F5` en Visual Studio Code no reinicia el depurador porque la aplicación ya no está corriendo. Para solucionar esto, necesitamos:
- Reiniciar automáticamente la aplicación cuando termina (esto ya se maneja en el script actual al guardar cambios).
- Manejar el bloqueo de la terminal y reiniciar la sesión si es necesario.
- Capturar combinaciones de teclas como "Alt + F5" para ejecutar comandos específicos (por ejemplo, `rerunServer`).

#### Solución con `xbindkeys` en Docker

Para capturar combinaciones de teclas dentro de un contenedor Docker, debemos instalar y configurar `xbindkeys` dentro del contenedor y permitir la interacción con el servidor X del host.

### Modificación del Dockerfile

Modifica tu Dockerfile para instalar `xbindkeys` y las dependencias necesarias para capturar combinaciones de teclas.

**Dockerfile Modificado:**

```Dockerfile
FROM golang:1.22.3-alpine3.20

# Instala dependencias y herramientas de depuración
RUN apk update && \
    apk add --no-cache bash git libc-dev g++ inotify-tools curl xbindkeys xdotool xauth

# Configura el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia todos los archivos del proyecto al contenedor
COPY . .

# Copia el archivo de configuración de xbindkeys
COPY .xbindkeysrc /root/.xbindkeysrc

# Da permisos de ejecución al script de entrada
RUN chmod +x /app/scripts/entrypoint.sh

# Descarga las dependencias del módulo Go
RUN go mod download && go mod verify

# Crea el directorio bin si no existe
RUN mkdir -p /app/bin

# Expone los puertos necesarios
EXPOSE 8080 2345

# Inicia xbindkeys y ejecuta el script de entrada
CMD ["sh", "-c", "xbindkeys -f /root/.xbindkeysrc & /app/scripts/entrypoint.sh"]
```

### Cambios Realizados

1. **Instalación de Dependencias:**
   - **`xbindkeys`**: Captura combinaciones de teclas.
   - **`xdotool` y `xauth`**: Necesarios para interactuar con el servidor X del host.

2. **Archivo de Configuración `.xbindkeysrc`:**
   - Configura `xbindkeys` para capturar "Alt + F5" y ejecutar `rerunServer`.

3. **Comando de Inicio (`CMD`):**
   - Ejecuta `xbindkeys` en segundo plano y luego el script de entrada, permitiendo escuchar combinaciones de teclas mientras la aplicación se ejecuta.

### Configuración del Archivo `.xbindkeysrc`

Crea un archivo `.xbindkeysrc` junto a tu Dockerfile con este contenido:

```plaintext
# Ejecuta un comando cuando se presiona Alt + F5
"sh /app/scripts/entrypoint.sh rerunServer"
    Alt+F5
```

### Ejecutar el Contenedor Docker con Acceso al Servidor X

Para que Docker capture eventos de teclado desde el host:

1. **Permitir Acceso al Servidor X:**

   Ejecuta en el host para dar acceso al servidor X:

   ```bash
   xhost +local:docker
   ```

2. **Ejecutar el Contenedor Docker:**

   Usa este comando para iniciar el contenedor con acceso al servidor X y a los dispositivos de entrada:

   ```bash
   docker run -it \
     --env="DISPLAY" \
     --volume="/tmp/.X11-unix:/tmp/.X11-unix:rw" \
     --device /dev/input \
     --privileged \
     --name my-container \
     my-image
   ```

### Consideraciones de Seguridad

Permitir que un contenedor acceda al servidor X y a los dispositivos de entrada puede suponer riesgos de seguridad. Asegúrate de realizar estas configuraciones en un entorno controlado y comprende los riesgos asociados.

### Conclusión

Con esta configuración:
- El contenedor Docker puede detectar combinaciones de teclas como "Alt + F5".
- Puedes ejecutar comandos dentro del contenedor en respuesta a eventos de teclado, facilitando la gestión de la aplicación sin salir del entorno Docker.

Asegúrate de revisar siempre los permisos y configuraciones de seguridad cuando trabajes con Docker y X11 para evitar posibles vulnerabilidades.