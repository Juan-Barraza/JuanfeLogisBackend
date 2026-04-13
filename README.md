# JuanfeLogis Backend ⚙️

## 📋 Descripción General
El **Backend de JuanfeLogis** es el corazón de nuestro sistema de gestión logística e inventario. Está diseñado usando una arquitectura orientada a servicios, rápida y de muy baja latencia gracias a la concurrencia nativa de Go. Se encarga de exponer la API REST consumida por el Frontend móvil y web, manejando toda la lógica de negocio subyacente, reglas de transacciones de inventario, persistencia de datos (base de datos), generación de códigos QR y autenticación de usuarios.

## 🚀 La Solución y Lógica Integrada del Código QR
La implementación del sistema de seguimiento por código QR para cajas de inventarios es una de las soluciones clave de este ecosistema tecnológico. En lugares donde hay un volumen de productos donados masivo, llevar el registro unitario es caótico. 

**Flujo lógico implementado en el backend:**
1. **Generación Automática:** Cuando desde el cliente se da la instrucción de crear una nueva Caja física virtual, el Backend de Go le asigna un **UUID v4 único** que mitiga cualquier conflicto de enumeramiento.
2. **Creación visual del QR (`skip2/go-qrcode`):** Inmediatamente, un servicio construye una URL directa hacia el Frontend (ej. `https://dominio.com/boxes/{uuid}`) y la codifica físicamente dentro de una imagen de Código QR.
3. **Mapeo Real vs Virtual:** Esa imagen PNG generada puede ser obtenida por el frontal para su impresión en tickets/stickers.
4. **Inventario Seguro (BoxStock):** Esta caja pasa a ser un "microlmacén". Usando una tabla pivot `BoxStock`, el backend restringe y maneja de manera atómica todos los productos y su cantidad exacta asociada a esa caja. Así, las transferencias de productos (donaciones y entregas) se pueden trazar directamente a cajas individuales sin errores de stock global.

Este modelo logra centralizar información compleja en un simple cuadradito escaneable que cualquier responsable con smartphone puede decodificar y auditar en menos de 2 segundos.

## 🛠 Tecnologías Utilizadas
Este servicio está construido sobre tecnologías de alto nivel garantizando tolerancia a fallos y gran eficiencia de procesamiento:
- **Golang (Go 1.26):** Lenguaje principal ágil, de tipado fuerte, excelente rendimiento en redes y simplicidad mantenible.
- **Fiber (v3):** Framework web fuertemente inspirado en Express, pero excepcionalmente rápido para el enrutamiento HTTP en Go. Funciona de manera soberbia sobre fasthttp bajando los tiempos de respuesta de los endpoints.
- **PostgreSQL:** Sistema estructurado e íntegro de Base de Datos Relacional, ideal para sistemas contables o de inventario por su manejo de transacciones sólidas.
- **GORM (v1.31):** El ORM para Go más completo, facilitando las migraciones (`AutoMigrate`), las consultas a BD y las relaciones N:M, 1:N sin acoplar directamente instrucciones SQL complejas a la lógica.
- **Go-QRCode:** Librería nativa para generar las matrices visuales (códigos de imagen QR) dentro del servidor.
- **Golang JWT:** Gestión de tokens seguros de sesión. Las contraseñas están doblemente resguardadas tras *crypto/bcrypt*.
- **Docker & Docker Compose:** Contenerización del entorno para facilitar pruebas unitarias o despliegues instantáneos del backend + la base de datos sin fricciones de entorno local.

## 🗄 Arquitectura y Base de Datos
El proyecto implementa un modelo estructural en capas facilitando la separación de preocupaciones y tests aislables:
```plaintext
/config       -> Inicialización de variables de entorno y Base de Datos (GORM Connect).
/dtos         -> (Data Transfer Objects) Estructuras específicas para validar la entrada del cliente.
/handlers     -> Capa controladora HTTP (Reciben Request -> llaman a Services -> devuelven JSON/Response).
/middlewares  -> Capa de interceptores (como el Autenticador JWT, CORS, Logger, etc).
/models       -> Esquemas de GORM, correspondientes exactamente a las tablas PostgreSQL.
/repositories -> Capa central de persistencia (Abstrae las consultas de GORM CRUD).
/routes       -> Registro y declaración de endpoints expuestos en Fiber.
/services     -> Capa de Reglas de Negocio (El cerebro de las transacciones y lógica).
/utils        -> Funciones y utilidades misceláneas y Helpers.
```

**Principales Entidades de Negocio (Models):**
- `Box` & `Location`: Las cajas que tendrán el código QR adherido y su procedencia.
- `Product`, `ProductType` & `Donor`: Permiten asociar exactamente de qué donante vino y qué producto específico es (tipo, precio donación vs precio venta, estado material).
- `BoxStock`: Relación transaccional vital; indica la multiplicidad de cada producto almacenado específicamente en cada caja.
- `User`: Administradores/Operarios bajo control de autenticación.

## ⚙️ Inicialización 

1. **Variables de entorno:** Configura el archivo `.env` en la raíz (usando los datos de PostgreSQL local o en un contenedor docker).
   ```env
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=secret
   DB_NAME=juanfe_db
   DB_PORT=5432
   JWT_SECRET=tu_secreto_aqui
   ```
2. **Contenedorizado (Recomendado):**
   Usa archivo `docker-compose.yml` para levantar todo sin preocupaciones:
   ```bash
   docker-compose up -d --build
   ```
3. **Local/Manual:**
   Resuelve los módulos e inicia el servidor. Gorm hará la auto-migración de las tablas correspondientes si no existen:
   ```bash
   go mod tidy
   go run main.go
   ```
   *El panel de la API abrirá por defecto en `http://localhost:8080`. Se recomienda probarla con herramientas como Postman o cURL.*
