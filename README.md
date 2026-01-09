# Task4Hub Microservices

Aplicación de microservicios desarrollada en **Go** para gestionar tareas, empresas, órdenes y usuarios. La solución utiliza **Firebase** como base de datos y está diseñada para ser escalable y modular.

## 📋 Tabla de Contenidos

- [Descripción General](#descripción-general)
- [Microservicios](#microservicios)
- [Requisitos](#requisitos)
- [Instalación](#instalación)
- [Configuración](#configuración)
- [Ejecución](#ejecución)
- [Estructura del Proyecto](#estructura-del-proyecto)

## 🎯 Descripción General

Task4Hub es una plataforma de gestión de tareas y órdenes con integración empresarial. Los microservicios están separados por dominio de negocio:

- **User Management**: Gestión de usuarios e invitaciones
- **Company Management**: Gestión de empresas e inventario
- **Order Management**: Gestión de órdenes y actividades

## 🔧 Microservicios

### 1. **Users Management** (`usersManagement/`)
Servicio encargado de la gestión de usuarios y control de acceso.

**Funcionalidades:**
- Crear usuarios
- Editar perfil de usuario
- Obtener información de usuarios
- Invitar usuarios
- Verificar invitaciones

**Puerto:** 8001 (predeterminado)

### 2. **Company Management** (`companyManagement/`)
Servicio para la gestión de empresas e inventario.

**Funcionalidades:**
- Crear y gestionar empresas
- Gestionar stock de productos
- Agregar y editar stock
- Reportes de inventario
- Insertar datos base

**Puerto:** 8002 (predeterminado)

### 3. **Order Management** (`orderManagement/`)
Servicio para la gestión de órdenes y actividades.

**Funcionalidades:**
- Crear órdenes
- Asignar empleados a órdenes
- Gestionar actividades
- Órdenes programadas
- Integración con Airbnb
- Reportes de órdenes diarias

**Puerto:** 8003 (predeterminado)

## 📦 Requisitos

- **Go** >= 1.18
- **Firebase** (cuenta con credenciales)
- **Docker** (opcional, para contenedorización)
- **Git**

## 🚀 Instalación

### 1. Clonar el Repositorio

```bash
git clone https://github.com/JostinAlvaradoS/Task4Hub_microservices.git
cd Task4Hub_microservices
```

### 2. Crear Carpeta de Configuración

```bash
mkdir -p config
```

### 3. Configurar Variables de Entorno

```bash
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

### 4. Descargar Dependencias

Para cada microservicio:

```bash
cd usersManagement && go mod download
cd ../companyManagement && go mod download
cd ../orderManagement && go mod download
```

## ⚙️ Configuración

### Firebase Configuration

1. Descarga el archivo `key.json` de tu proyecto Firebase desde la consola de Firebase
2. Colócalo en la carpeta `config/` creada anteriormente:

```
config/
└── key.json
```

3. Asegúrate de que cada servicio referencia correctamente la ubicación:

```go
// En firebase/firebase.go
opt := option.WithCredentialsFile("../../config/key.json")
```

### Variables de Entorno (Opcional)

Puedes crear un archivo `.env` en cada carpeta del microservicio:

```bash
# .env
PORT=8001
FIREBASE_PROJECT_ID=your-project-id
```

## 🏃 Ejecución

### Opción 1: Ejecutar Localmente

**Users Management:**
```bash
cd usersManagement
go run main.go
# Disponible en: http://localhost:8001
```

**Company Management:**
```bash
cd companyManagement
go run main.go
# Disponible en: http://localhost:8002
```

**Order Management:**
```bash
cd orderManagement
go run main.go
# Disponible en: http://localhost:8003
```

### Opción 2: Ejecutar Todos los Servicios

```bash
# En terminales separadas
cd usersManagement && go run main.go &
cd companyManagement && go run main.go &
cd orderManagement && go run main.go &
```

### Opción 3: Con Docker

```bash
docker-compose up
```

## 📁 Estructura del Proyecto

```
Task4Hub_microservices/
├── usersManagement/
│   ├── main.go
│   ├── go.mod
│   ├── firebase/
│   ├── handlers/
│   ├── models/
│   └── router/
├── companyManagement/
│   ├── main.go
│   ├── go.mod
│   ├── firebase/
│   ├── handlers/
│   ├── models/
│   └── router/
├── orderManagement/
│   ├── main.go
│   ├── go.mod
│   ├── firebase/
│   ├── handlers/
│   ├── models/
│   └── router/
└── config/
    └── key.json (crear manualmente)
```

## 🔗 Conexión entre Microservicios

Los servicios se comunican mediante HTTP requests. Ejemplos:

- Order Management llama a Company Management para restar stock
- Order Management llama a Users Management para validar usuarios
- Company Management valida empresas desde Users Management

## 📝 Modelos Principales

### User
- ID, nombre, email, rol, empresa asignada

### Company
- ID, nombre, información, stock de productos

### Order
- ID, cliente, empleados asignados, actividades, estado

### Stock
- ID, producto, cantidad, empresa

### Activity
- ID, descripción, estado, orden asociada, empleado

## 🛠️ Desarrollo

Para agregar nuevos endpoints:

1. Crear handler en `handlers/`
2. Definir modelo en `models/`
3. Registrar ruta en `router/router.go`
4. Implementar lógica de Firebase en `firebase/firebase.go`

## 📞 Contacto

**Autor:** Jostin Alvarado S.  
**GitHub:** [@JostinAlvaradoS](https://github.com/JostinAlvaradoS)

