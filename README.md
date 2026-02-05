# LINSITrack

## Correr el stack localmente

Teniendo Docker corriendo ( y abierto en windows), correr el siguiente comando en la raiz del proyecto: 

```bash
docker-compose up -d
```

## Documentacion

## Server de Documentacion / Requests de la API

Correr el siguiente comando en la raiz del proyecto:

```bash
go run utils/requests/linsi_track_documentation.go
```

A su vez, también está disponible la carpeta ./requests/LINSITrack_Requests/ para importar en Bruno, Postman, etc, junto con archivos de prueba en ./requests/Endpoints_Test_Data para simular uploads.

### Auth

- Admin

- Profesor

- Alumno

### Notificaciones

### Cursadas

### Comisiones

### Materias

### Evaluaciones

### TPs

### Competencias

### Archivos Anexos (por parte de profesores para los TPs)

### Entregas (de tps resueltos por los alumnos)

#### Crear Entrega con datos obligatorios
POST /entregas
body type: JSON
```json
{
  "fecha_hora": "2023-10-05T14:48:00Z",
  "alumno_id": 6,
  "tp_id": 11
}
```
(Almacenar el id de la entrega creada al recibir la respuesta del servidor para usarlo en las siguientes operaciones)

#### Adjuntar archivo a una entrega
POST /entregas/{id_de_entrega}/upload
body type: multipart/form-data
```json
{
  "file": "value: <base64-encoded-file-content>"
}
```