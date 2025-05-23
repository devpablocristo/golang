### **Resumen Breve: Manejo de Ambientes (Dev, Stg, Prod)**

#### **1. Desarrollo (Dev):**
- **Variables de entorno**: Se cargan desde un archivo local `.env`.
- **Gestión de secretos**: Opcional, pero puedes usar un respaldo automatizado cifrado GPG.
- **Herramientas**: No necesitas Consul o Vault; carga directa desde el `.env` usando **godotenv**.
- **Propósito**: Flexibilidad y facilidad para iterar rápidamente.

#### **2. Staging (Stg):**
- **Variables de entorno**: 
  - Datos críticos se inyectan directamente en el contenedor (via Docker Compose).
  - Configuración dinámica y secretos gestionados con **Consul** + **Vault**.
- **Gestión de secretos**: Vault asegura secretos sensibles.
- **Propósito**: Ambiente que simula producción, usado para pruebas finales.

#### **3. Producción (Prod):**
- **Variables de entorno**: 
  - Todas gestionadas de forma centralizada mediante **Vault** y/o servicios cloud (AWS Secrets Manager).
- **Configuración dinámica**: Usar **Consul** para configuración y servicio discovery.
- **Infraestructura**: Automatización con **Terraform** para despliegues consistentes y versionados.
- **Propósito**: Robustez, seguridad y escalabilidad.

#### **Conclusión:**
Cada ambiente tiene herramientas y configuraciones adaptadas a sus necesidades:
- **Dev**: Simple y local.
- **Stg**: Simula prod con herramientas de gestión de secretos y configuración dinámica.
- **Prod**: Máxima seguridad, automatización y uso de servicios cloud.

## **Guía para usar GPG (GNU Privacy Guard)**
### **1. Instalación de GPG**

- **Linux**: 
  ```bash
  sudo apt update && sudo apt install gnupg
  ```
- **Mac**:
  ```bash
  brew install gnupg
  ```
- **Windows**:
  - Descarga e instala desde [GPG for Windows](https://gnupg.org/download/index.html).

---

### **2. Crear una Clave GPG**

1. Ejecuta el siguiente comando para generar una clave GPG:
   ```bash
   gpg --full-generate-key
   ```
2. Selecciona los parámetros:
   - **Tipo de clave**: RSA y RSA.
   - **Tamaño**: 4096 bits.
   - **Validez**: Depende de tu preferencia (e.g., 0 para que no expire).
3. Introduce tu nombre y correo electrónico.
4. Crea una contraseña segura para proteger tu clave privada.

---

### **3. Exportar tu Clave Pública y Privada**

- **Clave pública** (compártela con tu equipo si trabajan juntos):
  ```bash
  gpg --armor --export tu_email@example.com > public_key.gpg
  ```
- **Clave privada** (guárdala en un lugar seguro, nunca la compartas):
  ```bash
  gpg --armor --export-secret-key tu_email@example.com > private_key.gpg
  ```

---

### **4. Cifrar el Archivo `.env`**

1. Cifra el archivo `.env`:
   ```bash
   gpg --encrypt --recipient "tu_email@example.com" .env
   ```
2. Esto creará un archivo llamado `.env.gpg`.

---

### **5. Descifrar el Archivo `.env`**

1. Para descifrar el archivo `.env.gpg`:
   ```bash
   gpg --decrypt .env.gpg > .env
   ```
2. Te pedirá la contraseña que configuraste al crear tu clave GPG.

---

### **6. Almacenar y Compartir**

- **En tu repositorio**:
  - No subas el archivo `.env`, pero puedes subir el archivo `.env.gpg`.
  - Asegúrate de que el equipo tenga la clave pública para descifrar.

- **En producción**:
  - Descifra el archivo al momento de desplegar:
    ```bash
    gpg --decrypt .env.gpg > .env
    ```

---

### **7. Automatización con Git Hooks (Opcional)**

Puedes agregar un **pre-commit hook** para cifrar automáticamente el `.env` antes de cada commit. Por ejemplo:

```bash
#!/bin/bash
gpg --encrypt --recipient "tu_email@example.com" .env
```

Guarda este script en `.git/hooks/pre-commit` y dale permisos de ejecución:

```bash
chmod +x .git/hooks/pre-commit
```

---

### **8. Consejos de Seguridad**
- Asegúrate de almacenar la clave privada en un lugar seguro.
- Usa una contraseña fuerte para tu clave GPG.
- Mantén el archivo `.env` fuera del repositorio agregándolo a `.gitignore`.- No es recomendable compartir la clave privada con tu equipo.

### ¿Qué hacer si trabajas en equipo?

1. **Usar claves públicas individuales:**
   - Cada miembro del equipo genera su propia clave pública y privada.
   - Tú o el responsable del archivo `.env` cifran el archivo con las claves públicas de todos los miembros del equipo. Así, cada uno puede descifrarlo con su clave privada.

   **Ejemplo de cifrado para múltiples destinatarios:**
   ```bash
   gpg --encrypt --recipient "miembro1@example.com" --recipient "miembro2@example.com" .env
   ```

2. **Rotación periódica de claves:**
   - Si decides compartir una clave privada (lo cual no es ideal), asegúrate de rotarla periódicamente y cambiarla si algún miembro del equipo abandona el proyecto.

La práctica más profesional es **nunca compartir claves privadas directamente**.

El flujo de control para la carga de variables y secretos entre Terraform, Vault y Consul, todos ejecutándose como contenedores Docker, funciona de la siguiente manera:

---

### **1. Terraform: Orquestador de infraestructura y configuración**
- **Rol**: Terraform actúa como la herramienta para configurar y administrar Vault y Consul.
- **Acción**: 
  - Terraform utiliza los proveedores (`providers`) de Vault y Consul para conectarse a sus APIs.
  - Lee los archivos `.tf` que contienen la definición de las variables, secretos y configuraciones.
  - Aplica los cambios necesarios en Vault y Consul.
- **Cómo opera**:
  - Terraform interactúa con los contenedores de Vault y Consul a través de sus direcciones (`http://vault:8200` y `http://consul:8500`).
  - Al ejecutar `terraform apply`, Terraform envía las configuraciones definidas en `main.tf` a Vault y Consul.

---

### **2. Vault: Almacén de secretos**
- **Rol**: Vault es el almacén seguro de secretos. Guarda información sensible como claves API, contraseñas de bases de datos, tokens, etc.
- **Acción**:
  - Terraform crea o actualiza los secretos en Vault usando recursos como `vault_generic_secret`.
  - Los secretos se almacenan en rutas jerárquicas, por ejemplo, `secret/staging/app`.
  - Las aplicaciones pueden obtener los secretos en tiempo de ejecución autenticándose con Vault.
- **Cómo opera**:
  - Vault recibe los secretos desde Terraform.
  - Los secretos están protegidos y requieren autenticación (token o mecanismo más avanzado) para acceder.

---

### **3. Consul: Configuración dinámica y descubrimiento de servicios**
- **Rol**: Consul almacena configuraciones generales de la aplicación (no necesariamente sensibles) y actúa como una base de datos distribuida para la configuración dinámica.
- **Acción**:
  - Terraform crea o actualiza claves en Consul usando recursos como `consul_key_prefix`.
  - Estas claves se estructuran en jerarquías, por ejemplo, `staging/app/name`.
  - Las aplicaciones leen estas configuraciones para obtener valores como el entorno, versión, URLs, etc.
- **Cómo opera**:
  - Consul actúa como un sistema de configuración accesible mediante su API o clientes específicos en las aplicaciones.
  - Los servicios pueden suscribirse a cambios en las configuraciones almacenadas para actualizarse dinámicamente.

---

### **Flujo Completo del Control de Variables**

1. **Inicialización**:
   - Los contenedores de Vault, Consul y Terraform se levantan con Docker Compose.
   - Terraform se inicializa (`terraform init`) y verifica que los proveedores de Vault y Consul estén disponibles.

2. **Ejecución de Terraform**:
   - Ejecutas `terraform apply` dentro del contenedor de Terraform.
   - Terraform lee los archivos `.tf` y envía las configuraciones a Vault y Consul.

3. **Carga de Variables en Vault**:
   - Las claves sensibles definidas en el archivo `.tf` (como claves API o contraseñas) se cargan en Vault.
   - Vault las almacena de forma segura, y las aplicaciones pueden autenticarse para obtenerlas.

4. **Carga de Configuraciones en Consul**:
   - Las configuraciones no sensibles (como nombres de aplicaciones, versiones, URLs) se cargan en Consul.
   - Consul permite a las aplicaciones leer estas configuraciones dinámicamente.

5. **Acceso desde las Aplicaciones**:
   - Las aplicaciones como `qh` se configuran para leer secretos de Vault y configuraciones de Consul en tiempo de ejecución.
   - El acceso a Vault y Consul se realiza a través de bibliotecas o APIs que autentican las aplicaciones.

6. **Actualizaciones Dinámicas**:
   - Si Terraform cambia las configuraciones, los contenedores de Vault y Consul se actualizan automáticamente.
   - Las aplicaciones que dependen de estas configuraciones pueden adaptarse en tiempo real (Consul) o mediante relectura programada (Vault).

---

### **Beneficios del Flujo**
- **Seguridad**: Los secretos sensibles están en Vault, protegidos por autenticación.
- **Centralización**: Consul centraliza las configuraciones dinámicas para todos los servicios.
- **Automatización**: Terraform gestiona tanto Vault como Consul, asegurando consistencia en los entornos.
- **Escalabilidad**: Este flujo es ideal para entornos complejos donde múltiples servicios necesitan secretos y configuraciones.

Esta configuración es modular y puede escalarse para incluir más entornos como `dev`, `qa`, y `prod`.