# Tienda_iglu Admin Panel

Este es un panel de administración para gestionar productos y órdenes en una tienda en línea. Permite a los administradores iniciar sesión, registrar nuevos administradores, y gestionar productos y órdenes.

## Requisitos

- [Node.js](https://nodejs.org/) (v14 o superior)
- [MongoDB](https://www.mongodb.com/) (local o en la nube)

## Instalación

1. **Clona el repositorio:**

   ```
   bash
   git clone 
   cd Backend_panel_admin
   ```

2. **Instala las dependencias:**

   ```
   bash
   npm install
   ```

3. **Configura las variables de entorno:**

   Crea un archivo `.env` en la raíz del proyecto y agrega las siguientes variables:

   ```plaintext
   JWT_SECRET=tu_secreto_jwt
   MONGODB_URI=tu_uri_de_mongodb
   ```

   Asegúrate de reemplazar `tu_secreto_jwt` y `tu_uri_de_mongodb` con tus propios valores.

## Ejecución

Para iniciar el servidor en modo desarrollo, utiliza el siguiente comando:

bash
npm run dev

Esto iniciará el servidor y podrás acceder al panel de administración en `http://localhost:3000/admin`.

## Rutas

- **/admin/login**: Página de inicio de sesión para administradores.
- **/admin/register**: Página para registrar nuevos administradores.
- **/admin/products**: Página para gestionar productos.
- **/admin/orders**: Página para gestionar órdenes.

## Contribuciones

Si deseas contribuir a este proyecto, por favor abre un issue o envía un pull request.
