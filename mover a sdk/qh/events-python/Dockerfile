# Imagen base oficial de Python
FROM python:3.9

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Ajusta la ruta para copiar el archivo requirements.txt
COPY requirements.txt .

# Instala las dependencias de la aplicación
RUN pip install --no-cache-dir -r requirements.txt

# Copia el resto del código de la aplicación
COPY . .

# Ajusta el CMD para reflejar la nueva estructura
CMD ["uvicorn", "cmd.main:app", "--host", "0.0.0.0", "--reload"]