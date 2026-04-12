# Lambda Go

API sencilla escrita en Go para AWS Lambda usando **Lambda Function URL**.

## Funcionalidad

Este proyecto expone dos rutas:

- `GET /api/hello` → devuelve `Hello, World!`
- `GET /api/key` → devuelve el valor del header `Authorization`
- Cualquier otra ruta → `404 Not Found`

## Requisitos

- Go 1.25+
- AWS CLI configurado
- Cuenta de AWS con permisos para Lambda
- Runtime de AWS Lambda compatible con Go

## Estructura del proyecto

- `main.go` — lógica principal de la Lambda
- `main_test.go` — tests unitarios
- `template.yml` — plantilla de despliegue
- `buildspec.yml` — configuración de build
- `README.md` — documentación del proyecto

## Ejecución local

Instala dependencias y ejecuta los tests:
```bash
go test ./..
```

Para compilar el binario:
bash GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go

## Despliegue en AWS Lambda

1. Compila el binario para Linux ARM64:
   ```bash
   GOOS=linux CGO_ENABLED=0 GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
   ```

2. Sube el binario a AWS Lambda.

3. Configura la función con:
    - **Handler:** `main.handler`
    - **Runtime:** el compatible con Go en Lambda
    - **Arquitectura:** `arm64`

4. Si usas una **Function URL**, asegúrate de habilitarla para invocar la Lambda.

## Despliegue con SAM

El proyecto incluye un archivo `template.yml` para desplegar con AWS SAM (Serverless Application Model):

1. Construye el proyecto:
   ```bash
   sam build
   ```

2. Despliega en AWS:
   ```bash
   sam deploy --guided
   ```

3. En el primer despliegue, SAM te pedirá configurar:
    - Nombre del stack
    - Región de AWS
    - Confirmación de cambios

4. La URL de la API se mostrará en los outputs del despliegue.

## Integración con CodePipeline

El archivo `buildspec.yml` permite integrar el proyecto en AWS CodePipeline para CI/CD automatizado:

1. Crea un pipeline en AWS CodePipeline.

2. Configura el repositorio como fuente (CodeCommit, GitHub, etc.).

3. Añade una fase de build usando AWS CodeBuild con el `buildspec.yml` incluido.

4. El pipeline ejecutará automáticamente:
    - Tests unitarios
    - Compilación con SAM
    - Despliegue a AWS Lambda

5. Asegúrate de que el rol de CodeBuild tenga permisos para:
    - CloudFormation
    - Lambda
    - S3
    - IAM (crear roles de ejecución)
