# TaskMicroservices

## Configuración del Proyecto

### 1. Crear la Carpeta `config`

En la raíz del proyecto, crea una carpeta llamada `config`:

```bash
mkdir config

Configurar las variables de entorno (E caso de ya estar dockerizado esto no sera necesario):
 
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin


Referenciar correctamente la carpeta de config/key.json en el archivo firebase/firebase.go para poder inicializar firebase para el proyecto.# Task4Hub_microservices
